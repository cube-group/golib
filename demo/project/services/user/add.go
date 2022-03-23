package user

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/cube-group/golib/conf"
	"time"
)

func Add(c *gin.Context) interface{} {
	var val ValidatorAdd
	if err := c.ShouldBindWith(&val, binding.Form); err != nil {
		return nil
	}

	if err := conf.DB().First(&models.User{}, "username=?", val.Name).Error; err == nil {
		return gin.H{"detail": "name已存在"}
	}

	saver := &models.User{
		Username: val.Name,
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}
	if err := conf.DB().Save(&saver).Error; err != nil {
		return nil
	}

	return saver
}
