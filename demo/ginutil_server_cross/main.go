package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/middlewares"
	"github.com/cube-group/golib/ginutil/server"
)

func init() {
	conf.Init(conf.Template{
		AppYamlPath: "./",
	})
}

func main() {
	server.Create(server.Config{
		FuncController:  routes,
	})
}

func routes(engine *gin.Engine) {
	engine.GET("/dashboard", middlewares.CrossDomain(middlewares.CrossDomainConfig{}), dashboard)
}

func dashboard(c *gin.Context) {
	ginutil.JsonAuto(c, c.Request.RequestURI, nil, nil)
}
