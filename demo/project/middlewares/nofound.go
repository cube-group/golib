package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"net/http"
)

func NotFound(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusNotFound, "errors/404.html", nil)
	} else {
		ginutil.JsonError(c, "404！您访问的页面未找到！", nil)
	}
}