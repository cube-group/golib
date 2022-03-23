//Author: linyang
//Time: 2021-02
// 统一的http请求类

//基于《微软云设计模式》
//基于微软环形熔断器论文：https://docs.microsoft.com/en-us/previous-versions/msp-n-p/dn589784%28v=pandp.10%29?redirectedfrom=MSDN
//断路器实现实现主要分为三部分：状态统计、状态转移、请求执行
//
//状态统计：统计已经执行的请求的成功失败的数量，以确定是否需要进行状态转移
//状态转移：根据当前统计信息和当前状态来进行目标状态的确定及转移操作
//请求执行：代理前端任务的执行，如果当前状态不需要进行尝试执行，就直接返回错误，避免资源浪费
//
//熔断器有三种状态，四种状态转移的情况：
//
//三种状态：
//熔断器关闭状态, 服务正常访问
//熔断器开启状态，服务异常
//熔断器半开状态，部分请求限流访问
//
//四种状态转移：
//在熔断器关闭状态下，当失败后并满足一定条件后，将直接转移为熔断器开启状态。(open)
//在熔断器开启状态下，如果过了规定的时间，将进入半开启状态，验证目前服务是否可用。
//在熔断器半开启状态下，如果出现失败，则再次进入关闭状态。(half-open)
//在熔断器半开启后，所有请求（有限额）都是成功的，则熔断器关闭。所有请求将正常访问。(closed)
//
// readyToTrip: 自定义断路器open状态的条件函数，return true则会断路器开启
package breaker

import (
	"github.com/sony/gobreaker"
	"github.com/cube-group/golib/log"
	"time"
)

type BreakerDoFunc func() (interface{}, error)

type BreakerConfig struct {
	//熔断单元内正式进入失败率判断的请求次数
	//默认：5，最小：5，最大：100
	BreakerStartRequests uint32
	//熔断单元内正式熔断的失败率
	//默认：0.5，最小：0.5，最大：0.9
	BreakerStartFailureRatio float64
	//maxRequests: 限制half-open状态下最大的请求数，避免海量请求将在恢复过程中的服务再次失败
	//熔断单元进入半开状态后最大允许请求次数，若为0则不作限制
	//默认：0
	BreakerMaxHalfOpenRequests uint32
	//interval: 用于在closed状态下，断路器多久清除一次Counts信息，如果设置为0则在closed状态下不会清除Counts
	//熔断单元生命周期（单位：秒）
	//默认：30，最小：30，最大：300
	BreakerInterval int64
	//timeout: 进入open状态下，多长时间切换到half-open状态，默认60s
	//熔断单元从half-open到open的等待时间（单位：秒）
	//默认：5，最小：5，最大：60
	BreakerTimeout int64
}

//熔断器暂存器
var breakers = make(map[string]*gobreaker.CircuitBreaker)

//设置熔断器
func Set(name string, cfg BreakerConfig) error {
	if cfg.BreakerStartRequests < 5 {
		cfg.BreakerStartRequests = 5
	} else if cfg.BreakerStartRequests > 100 {
		cfg.BreakerStartRequests = 100
	}
	if cfg.BreakerStartFailureRatio < 0.5 {
		cfg.BreakerStartFailureRatio = 0.5
	} else if cfg.BreakerStartFailureRatio > 0.9 {
		cfg.BreakerStartFailureRatio = 0.9
	}
	if cfg.BreakerInterval < 30 {
		cfg.BreakerInterval = 30
	} else if cfg.BreakerInterval > 300 {
		cfg.BreakerInterval = 300
	}
	if cfg.BreakerTimeout < 5 {
		cfg.BreakerTimeout = 5
	} else if cfg.BreakerTimeout > 60 {
		cfg.BreakerTimeout = 60
	}

	bk := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:        name,
		MaxRequests: cfg.BreakerMaxHalfOpenRequests,
		Interval:    time.Duration(cfg.BreakerInterval) * time.Second,
		Timeout:     time.Duration(cfg.BreakerTimeout) * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			if counts.Requests >= cfg.BreakerStartRequests && failureRatio >= cfg.BreakerStartFailureRatio {
				return true //进入熔断
			}
			return false //不进入熔断
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			log.StdWarning("Breaker", name, "OnStateChange", from.String(), to.String())
		},
	})
	breakers[name] = bk
	return nil
}

//使用熔断器执行函数
func Do(name string, f BreakerDoFunc) (interface{}, error) {
	bk, ok := breakers[name]
	if !ok {
		if err := Set(name, BreakerConfig{}); err != nil {
			return nil, err
		}
		bk = breakers[name]
	}

	return bk.Execute(func() (interface{}, error) {
		return f()
	})
}
