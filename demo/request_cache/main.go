package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/cube-group/golib/cache"
	"github.com/cube-group/golib/conf"
	"time"
)

func main() {
	conf.Init(conf.Template{
		AppYamlPath: "https://eoffcn-software.oss-cn-beijing.aliyuncs.com/application.yaml",
	})

	cacheKey := "aaa"
	conf.Redis().Del(cacheKey)

	var res interface{}
	if err := cache.RequestWithCache(cacheKey, time.Minute, A, &res); err != nil {
		log.Fatal("err", err)
	}

	fmt.Println(res)
	fmt.Println(conf.Redis().Get(cacheKey).Result())
}

func A() (interface{}, error) {
	return gin.H{"w":1}, nil
}
