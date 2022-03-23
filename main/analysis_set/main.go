package main

import (
	"fmt"
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
	b := &analysis.LogTcpJson{Name: "form"}
	defer b.Close()
	err := b.Put(map[string]interface{}{
		"content": map[string]interface{}{
			"formid": "www-main-123456", //form id
			"name":   "linyang",
			"phone":  "15901214776",
			"ext": map[string]interface{}{
				"examId": "123",
				"examArea": []interface{}{
					map[string]interface{}{"name": "10"},
					map[string]interface{}{"name": "20"},
				},
			},
		},
	})
	fmt.Println(err)
}
