package controllers

import (
	"app/services/user"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/ginutil"
	"net/http"
)

type UserController struct {
	ginutil.IController
}

func (t *UserController) Init(group *gin.RouterGroup) {
	group.GET(".", t.userList)
	group.POST(".", t.userAdd)

	g := group.Group("/detail")//此处可见中间件
	g.GET("/:id", t.userGet)
}

func (t *UserController) userList(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": user.List(c)})
}

func (t *UserController) userAdd(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": "success", "data": user.Add(c)})
}

func (t *UserController) userGet(c *gin.Context) {
	res,err:=user.Get(c)
	c.JSON(http.StatusOK, gin.H{"code": 0, "msg": err, "data": res})
}
