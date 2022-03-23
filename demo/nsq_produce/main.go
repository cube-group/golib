package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	nsq2 "github.com/nsqio/go-nsq"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/nsq"
	"time"
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
	producer, err := nsq.NewNsqProducer(viper.GetStringSlice("nsq.d"))
	if err != nil {
		log.StdFatal("nsq", "produce", err)
	}
	defer producer.Close()

	msg := gin.H{
		"name": "demo", "date": time.Now().String(),
	}
	//正常生产消息
	fmt.Println(producer.One(
		viper.GetString("nsq.topic"),
		msg,
		0, //是否延迟消费
	))
	//生产多个消息
	fmt.Println(producer.Multiple(
		viper.GetString("nsq.topic"),
		[]interface{}{msg, msg},
	))
}
