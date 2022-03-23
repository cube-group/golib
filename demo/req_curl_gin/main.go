package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/net/curl"
	"time"
)

func main() {
	curl.ReqLog = true                //设置日志
	curl.ReqTimeout = 5 * time.Second //设置超时时间
	server.Create(server.Config{
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.GET("/info", info)
}

func info(c *gin.Context) {
	resp, err := curl.Get("http://localhost:8089/failed", req.Param{"a": 1}, c) //会自动修改上下文中的链id然后通过curl类库传递下去
	fmt.Println(resp, err)
}
