package main

import (
	"flag"
	"fmt"
	"github.com/cube-group/golib/nsq"
	"log"
	"time"
)

var nsqd string
var topic string

func main() {
	flag.StringVar(&nsqd, "nsqd", "", "nsqd address")
	flag.StringVar(&topic, "topic", "", "topic name")
	flag.Parse()

	nsqds := make([]string, 0)
	if nsqd != "" {
		nsqds = append(nsqds, nsqd)
	}else{
		log.Fatal("nsqd is nil.")
	}
	if topic == "" {
		log.Fatal("topic is nil.")
	}

	i, err := nsq.NewNsqProducer(nsqds)
	if err != nil {
		log.Fatal("produce", err)
	}
	fmt.Println(i.One(topic, map[string]interface{}{"time": time.Now().Unix()}, 0))
}
