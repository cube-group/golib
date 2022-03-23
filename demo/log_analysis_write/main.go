package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/aliyun/analysis"
)

func main() {
	//以本地日志模式间接写入日志
	a := &analysis.LogFileJson{Name: "sso"}
	a.Put(map[string]interface{}{
		"env": "asdfa",
		"a":   1,
	})

	//以tcp模式直接写入日志
	//查询的JSON值一定要为text否则难以查询
	//todo 记得配置aliyun log ak sk
	//todo LogTcpJson实例可以长期保存使用
	b := &analysis.LogTcpJson{
		AliyunAccessKeyId:     "",
		AliyunAccessKeySecret: "",
		AnyStore:              true,
		ProjectName:           "huabei2-form",
		Name:                  "sso",
	}
	defer b.Close()
	err := b.Put(map[string]interface{}{
		"uid":  "111111",
		"type": "demo",
		"env":  "local",
		"ext": gin.H{
			"formid": "www-main-123456", //form id
			"examArea": []interface{}{
				map[string]interface{}{"name": "10"},
				map[string]interface{}{"name": "20"},
			},
		},
	})

	fmt.Println(err)
}
