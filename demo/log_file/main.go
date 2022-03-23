package main

import (
	"github.com/cube-group/golib/log"
	"time"
)

func main() {
	log.AppName = "demo" //定义日志应用名称
	log.Path = "."       //定义日志被分片存储的目录
	for {
		log.FileLogInfo("adf", "123", "0", "adsf", "") //普通日志
		go time.Sleep(time.Millisecond * 2)
	}

	//log.FileLogErr("adf", "123", "0", "adsf", "")  //错误日志
}
