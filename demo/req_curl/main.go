package main

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/cube-group/golib/net/curl"
	"time"
)

func main() {
	//使用默认curl实例发送请求
	curl.ReqLog = true                //设置日志
	curl.ReqTimeout = 5 * time.Second //设置超时时间
	resp, err := curl.Get("http://localhost:8089/failed", req.Param{"a": 1}, req.Header{"a": "1"})
	fmt.Println(resp, err)
}
