package ginutil

import "github.com/gin-gonic/gin"

type IController interface {
	Init(group *gin.RouterGroup)
}

func Group(engine *gin.Engine, i IController, relativePath string, handlers ...gin.HandlerFunc) {
	i.Init(engine.Group(relativePath, handlers...))
}
