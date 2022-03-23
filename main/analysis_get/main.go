package main

import (
	"fmt"
	"github.com/cube-group/golib/aliyun/analysis"
	"time"
)

func main() {
	//获取sso的分析
	a := &analysis.LogReader{Name: "analysis-demo", ProjectName: "huabei2-record", AnyStore: true}
	query := "* | select *"
	res, _ := a.GetLogs(query, time.Now().Unix()-10000, time.Now().Unix(), 2, 0)
	fmt.Println(len(res))
	//bytes, _ := json.Marshal(res)
	//fmt.Println(string(bytes))
}
