package main

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/net/breaker"
	"github.com/cube-group/golib/net/curl"
	"time"
)

func init() {
	if err := breaker.Set("scene1", breaker.BreakerConfig{
		BreakerStartRequests:     10,  //进入熔断判断的失败请求次数
		BreakerStartFailureRatio: 0.6, //进入熔断判断的失败率
		BreakerInterval:          30,  //30秒内请求次数>=10且失败率>=0.6则进入熔断
		BreakerTimeout:           5,   //5秒内从熔断状态恢复
	}); err != nil {
		log.StdFatal("Breaker", "Init", err)
	}
}

func main() {
	for {
		res, err := breaker.Do("scene1", func() (interface{}, error) {
			//todo 当然该函数内也可以操作mysql、redis、elasticsearch等
			resp, err := curl.Get("http://localhost:8080/dashboard", req.Param{"a": 1}, req.Header{"a": "1"})
			if err != nil {
				return "", err
			} else {
				return resp.String(), nil
			}
		})
		fmt.Println(res, err)

		time.Sleep(time.Second)
	}
}
