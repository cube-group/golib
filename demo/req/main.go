package main

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/cube-group/golib/log"
)

func main() {
	resp, err := req.Get("https://xx.com", req.Param{"key1": "value1"})
	if err != nil {
		log.StdFatal("Req", "Err", err)
	}
	var result map[string]string
	if err := resp.ToJSON(&result); err != nil {
		log.StdFatal("Req", "ToJson", "Err", err)
	}
	fmt.Println(result)
}
