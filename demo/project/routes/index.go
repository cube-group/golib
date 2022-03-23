package routes

import (
	"app/controllers"
	"app/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
)

func Init(engine *gin.Engine) {
	//before middleware
	engine.Use(middlewares.Before())

	//controllers
	ginutil.Group(
		engine,
		&controllers.UserController{},
		"/user",
		middlewares.Auth(),
	)
	ginutil.Group(
		engine,
		&controllers.DemoController{},
		"/demo",
	)

	//after middleware
	engine.Use(middlewares.After())
}
