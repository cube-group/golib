package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func After() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("after")
		c.Next()
	}
}
