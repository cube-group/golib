package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/types/jsonutil"
)

func init() {
	conf.Init(conf.Template{
		AppYamlPath: "./",
	})
}

type validator struct {
	Name2 string `form:"name" binding:"required"`
}

type validator2 struct {
	Name string `form:"name" binding:"required"`
}

func main() {
	server.Create(server.Config{
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.Use(mid)
	engine.POST("/", dashboard)
}

func mid(c *gin.Context) {
	var val validator
	var err = ginutil.ShouldBind(c, &val)
	fmt.Println("mid", err)
}

func dashboard(c *gin.Context) {
	var val validator2
	var err = ginutil.ShouldBind(c, &val)
	fmt.Println(jsonutil.ToString(val))
	ginutil.JsonAuto(c, c.Request.RequestURI, err, nil)
}
