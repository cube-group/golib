package nsq

import (
	"github.com/nsqio/go-nsq"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/log"
	"time"
)

type NsqConsumeConfig struct {
	Topic               string
	Channel             string
	MaxAttempts         uint16
	LookupdPollInterval time.Duration

	Nsqds       []string
	Nsqlookupds []string
	Handler     nsq.Handler //回调对象
}

func NewNsqConsumer(c NsqConsumeConfig) (*nsq.Consumer, error) {
	if c.Topic == "" {
		return nil, errors.New("topic is nil.")
	}
	if c.Channel == "" {
		return nil, errors.New("channel is nil.")
	}
	if c.Nsqds == nil || len(c.Nsqds) == 0 {
		return nil, errors.New("nsqds is nil.")
	}
	if c.Nsqlookupds == nil || len(c.Nsqlookupds) == 0 {
		return nil, errors.New("nsqlookupds is nil.")
	}
	if c.Handler == nil {
		return nil, errors.New("handler is nil.")
	}
	cfg := nsq.NewConfig()
	cfg.MaxInFlight = len(c.Nsqds) // 最大允许向几台NSQD接收消息
	if c.MaxAttempts > 0 {
		cfg.MaxAttempts = c.MaxAttempts //最大重试次数，min:0, max:65535, default5
	}
	if c.LookupdPollInterval > 0 {
		cfg.LookupdPollInterval = c.LookupdPollInterval //设置lookup心跳时间，min:10s, max:10min, default:60s
	}
	i, err := nsq.NewConsumer(c.Topic, c.Channel, cfg)
	if err != nil {
		return nil, err
	}
	i.SetLogger(nil, 0)     //屏蔽系统日志
	i.AddHandler(c.Handler) // 添加消费者接口

	if c.Nsqlookupds != nil && len(c.Nsqlookupds) > 0 {
		//建立NSQLookUpd连接间接消费
		log.StdOut("NSQ", "Consume.ConnectToNSQLookupds", c.Nsqlookupds)
		if err := i.ConnectToNSQLookupds(c.Nsqlookupds); err != nil {
			return nil, err
		}
	} else if c.Nsqds != nil && len(c.Nsqds) > 0 {
		//建立NSQd连接直接消费
		log.StdOut("NSQ", "Consume.ConnectToNSQDs", c.Nsqds)
		if err := i.ConnectToNSQDs(c.Nsqds); err != nil {
			return nil, err
		}
	} else {
		i.Stop()
		return nil, errors.New("no connection used")
	}

	return i, nil
}
