package user

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil"
)

func List(c *gin.Context) interface{} {
	var val ValidatorList
	if err := c.ShouldBindQuery(&val); err != nil {
		return nil
	}

	username, _ := c.Get("username")

	var total int64
	if err := conf.DB().Model(&models.User{}).Where("username=?", username).Count(&total).Error; err != nil {
		return nil
	}

	var users []*models.User
	paginate := ginutil.GetPageNation(c, int(total))
	if err := conf.DB().Limit(paginate.PageSize).Offset(paginate.PageOffset).Order("id DESC").Find(&users).Error; err != nil {
		return nil
	}

	return users
}
