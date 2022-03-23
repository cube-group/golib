package user

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/e"
)

func ContextUsername(c *gin.Context) string {
	var username string
	i, exist := c.Get("username")
	if !exist {
		return ""
	}
	e.TryCatch(func() {
		username = i.(string)
	})
	return username
}
