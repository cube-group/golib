package main

import (
	"flag"
	"fmt"
	"github.com/nsqio/go-nsq"
	nsq2 "github.com/cube-group/golib/nsq"
	"log"
	"time"
)

type HandleMessage struct {
}

func (this *HandleMessage) HandleMessage(message *nsq.Message) error {
	fmt.Println(string(message.Body), message.ID, message.Attempts, message.Timestamp)
	message.Finish()
	return nil
}

var nsqd string
var nsqlookupd string

func main() {
	flag.StringVar(&nsqd, "nsqd", "", "nsqd address")
	flag.StringVar(&nsqlookupd, "nsqlookupd", "", "nsqlookupd address")
	flag.Parse()

	nsqds := make([]string, 0)
	if nsqd != "" {
		nsqds = append(nsqds, nsqd)
	}
	nsqlookupds := make([]string, 0)
	if nsqlookupd != "" {
		nsqlookupds = append(nsqlookupds, nsqlookupd)
	} else {
		log.Fatal("nsqlookupd is nil.")
	}
	i, err := nsq2.NewNsqConsumer(nsq2.NsqConsumeConfig{
		Topic:               "nsq-golib",
		Channel:             "default",
		Nsqds:               nsqds,
		Nsqlookupds:         nsqlookupds,
		LookupdPollInterval: 3 * time.Second,
		Handler:             new(HandleMessage),
	})
	if err != nil {
		log.Fatal(err)
	}
	for {
		fmt.Println(i.Stats())
		time.Sleep(time.Hour)
	}
}
