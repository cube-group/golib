package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil/middlewares"
)

func main() {
	engine := gin.Default()
	engine.Use(middlewares.RateLimiter(1, ""))
	engine.Use(func(context *gin.Context) {
		context.Status(206)
	})
	engine.GET("/", func(context *gin.Context) {
		//context.String(http.StatusOK, "hello")
	})
	engine.Run(":8000")
}
