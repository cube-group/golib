package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/ginutil/session"
	"github.com/cube-group/golib/net/curl"
	"math/rand"
	"net/http"
	"time"
)

func init() {
	conf.Init(conf.Template{
		AppYamlPath: "./",
	})
}

func main() {
	server.Create(server.Config{
		StaticRoot: "public",
		//UseRedisSession: true,
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.Any("/timeout", func(context *gin.Context) {
		time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
		context.JSON(http.StatusOK, gin.H{"a": time.Now().String()})
	})
	engine.GET("/nojson", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})
	engine.GET("/dashboard", dashboard)
	engine.GET("/session_set", sessionSet)
	engine.GET("/session_get", sessionGet)
	engine.GET("/abc/:path", func(c *gin.Context) {
		c.String(http.StatusOK,c.Param("path"))
	})
	engine.GET("/abcd/*path", func(c *gin.Context) {
		c.String(http.StatusOK,c.Param("path"))
	})

	gp := engine.Group("/user")
	gp.GET(".", func(context *gin.Context) {

	})
	gp.DELETE("/:id", func(context *gin.Context) {

	})
	gp.GET("/:id", func(context *gin.Context) {

	})

}

func dashboard(c *gin.Context) {
	curl.Get("/abc", c) //curl tracing
	ginutil.JsonAuto(c, c.Request.RequestURI, nil, nil)
}

func sessionSet(c *gin.Context) {
	username, err := session.Session(c).Get("username")
	ginutil.JsonAuto(c, "ok", err, username)
}

func sessionGet(c *gin.Context) {
	ginutil.JsonAuto(c, "", nil, session.Session(c).Set("username", time.Now().String()))
}
