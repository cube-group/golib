package controllers

import (
	"app/services/demo"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"net/http"
)

type DemoController struct {
	ginutil.IController
}

func (t *DemoController) Init(group *gin.RouterGroup) {
	group.GET(".", t.index)
	group.POST(".", t.post)
}

func (t *DemoController) index(c *gin.Context) {
	c.HTML(http.StatusOK, "demo/index.html", gin.H{"uri": demo.GetRequestURI(c)})
}

func (t *DemoController) post(c *gin.Context) {
	if err := demo.SetRequestURI(c); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 3000, "msg": err.Error(), "data": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": nil})
	}
}
