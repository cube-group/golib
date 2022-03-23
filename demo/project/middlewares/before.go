package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Before() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before")
		c.Next()
	}
}
