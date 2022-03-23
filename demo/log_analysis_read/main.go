package main

import (
	"encoding/json"
	"fmt"
	"github.com/cube-group/golib/aliyun/analysis"
	"time"
)

func main() {
	//获取sso的分析
	a := &analysis.LogReader{AliyunAccessKeyId: "", AliyunAccessKeySecret: "", Name: "form"}
	query := "* and content.name:linyang and content.phone:15901214776 and content.ext.examArea.name:10 "
	res, err := a.GetLogs(query, time.Now().Unix()-10000, time.Now().Unix(), 10, 0)
	bytes, _ := json.Marshal(res)
	fmt.Println(err)
	fmt.Println(string(bytes))
}
