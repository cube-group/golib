package main

import (
	"fmt"
	nsq2 "github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/nsq"
)

type Handler struct {
}

func (t *Handler) HandleMessage(message *nsq2.Message) error {
	fmt.Println(string(message.Body), message.Attempts)
	return nil //说明ack success
}

func init() {
	conf.Init(conf.Template{
		AppYamlPath: "./",
	})
}

func main() {
	nsq.NewNsqConsumer(nsq.NsqConsumeConfig{
		Topic:       viper.GetString("nsq.topic"),
		Channel:     viper.GetString("nsq.channel"),
		MaxAttempts: uint16(viper.GetUint("nsq.max_attempts")),
		Nsqds:       viper.GetStringSlice("nsq.d"),
		Nsqlookupds: viper.GetStringSlice("nsq.lookupd"),
		Handler:     new(Handler),
	})
}
