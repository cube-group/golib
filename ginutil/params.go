// Author: chenqionghe
// Time: 2018-10
// json编码

package ginutil

import "github.com/gin-gonic/gin"

//获取请求参数
func RequestParams(c *gin.Context) map[string]interface{} {
	return map[string]interface{}{
		"get":  c.Request.URL.Query(),
		"post": c.Request.PostForm,
	}
}

//将query参数转换为map
func RequestQueryMap(c *gin.Context) map[string]interface{} {
	values := make(map[string]interface{})
	for k, v := range c.Request.URL.Query() {
		if v == nil || len(v) == 0 {
			values[k] = ""
		} else {
			values[k] = v[0]
		}
	}
	return values
}
