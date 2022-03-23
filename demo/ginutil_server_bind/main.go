package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/server"
)

func main() {
	server.Create(server.Config{
		StaticRoot: "public",
		//UseRedisSession: true,
		FuncController: routes,
	})
}

type val struct {
	Name string `form:"name" binding:"required"`
	Uid  int    `form:"uid" binding:"required"`
}

func routes(engine *gin.Engine) {
	engine.POST("/", middle, index)
}

func middle(c *gin.Context) {
	fmt.Println("middle")
	var val val
	ginutil.ShouldBind(c, &val)
	fmt.Println(val)
}

func index(c *gin.Context) {
	fmt.Println("index")
	var val val
	c.ShouldBind(&val)
	fmt.Println(val)
}
