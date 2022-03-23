package ginutil

import "github.com/gin-gonic/gin"

//请求是否为ajax
func IsAjax(c *gin.Context) bool {
	res := c.GetHeader("X-Requested-With")
	return res != "" && res != "null"
}
