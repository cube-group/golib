package main

import (
	"fmt"
	"github.com/cube-group/golib/aliyun/analysis"
	"time"
)

func main() {
	//获取sso的分析
	a := &analysis.LogReader{Name: "analysis-demo", ProjectName: "huabei2-record", AnyStore: true}
	a.ConsumeCount = 1                      //日志消费能力
	shardCursors := make(map[int]string, 0) //shard-cursor记录

	for {
		//模拟断开重连
		a.ConsumeLogs(
			0,
			func(logs []map[string]string, shardId int, nextCursor string, err error) {
				fmt.Println(shardId, nextCursor)
				shardCursors[shardId] = nextCursor
				if logs != nil {
					for _, t := range logs {
						fmt.Println(t)
					}
				}
				a.ConsumeCount = 0 //停止消费
			},
			shardCursors,
		)

		time.Sleep(15 * time.Second)
	}
}
