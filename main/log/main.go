package main

import (
	"github.com/cube-group/golib/log"
	"time"
)

func main(){
	log.AppName = "demo"
	log.Path = "/data/log/"
	log.FileLogInfo("adf","123","0","adsf","")
	log.FileLogInfo("adf","123","0","adsf","")
	log.FileLogInfo("adf","123","0","adsf","")
	log.FileLogInfo("adf","123","0","adsf","")
	time.Sleep(time.Hour)
}
