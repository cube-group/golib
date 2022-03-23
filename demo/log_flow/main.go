package main

import (
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/log"
	"net/http"
)

func main() {
	log.Path = "/Users/linyang/"
	log.AppName = "demo"

	server.Create(server.Config{
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.GET("/1", func(c *gin.Context) {
		//TODO 业务处理
		log.Push(log.LContent{ //TODO 以本地Log日志模式写入日志
			Flow:    "test",
			Route:   "1",
			Ext:     gin.H{"route": 1},
			Context: c,
		})
		go req.Get("http://127.0.0.1:8080/2", req.Header{log.LFlowIDKey: log.FlowID(c)})
	})
	engine.GET("/2", func(c *gin.Context) {
		log.Push(log.LContent{ //TODO 以协程TCP模式写入日志
			Flow:    "test",
			Route:   "2",
			Ext:     gin.H{"route": 2},
			Context: c,
		})
		c.String(http.StatusOK, "ok")
	})
}
