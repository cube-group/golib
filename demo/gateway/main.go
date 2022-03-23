package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/middlewares"
	"github.com/cube-group/golib/ginutil/server"
)

func main() {
	server.Create(server.Config{
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.Use(middlewares.Gateway()) //如果不是来自于api网关的转发则会直接报错中断
	engine.GET("/request", request)
}

func request(c *gin.Context) {
	ginutil.JsonAuto(c, c.Request.RequestURI, nil, nil)
}
