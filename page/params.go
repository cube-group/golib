package page

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/convert"
	"net/url"
)

//获取get参数
func GetParams(c *gin.Context) map[string]string {
	u, _ := url.Parse(c.Request.URL.String())
	q := u.Query()
	params := make(map[string]string)
	for key, values := range q {
		params[key] = values[0]
	}
	return params
}

//获取页数
func GetPage(c *gin.Context) uint {
	return convert.MustUint(c.DefaultQuery("page", "1"))
}

//获取分页大小
func GetPageSize(c *gin.Context) uint {
	return convert.MustUint(c.DefaultQuery("pagesize", c.DefaultQuery("pageSize", "20")))
}
