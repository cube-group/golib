package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/session"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//todo write something
		if s := session.Session(c); s != nil {
			username, err := s.Get("username")
			if err == nil && username != "" {
				c.Set("username", username)
				c.Next()
				return
			}
		}

		ginutil.JsonError(c, "no login", nil)
		c.Abort()
	}
}
