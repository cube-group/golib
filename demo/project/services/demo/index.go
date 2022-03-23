package demo

import (
	"app/models"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
)

func GetRequestURI(c *gin.Context) string {
	return c.Request.RequestURI
}

func SetRequestURI(c *gin.Context) error {
	db := conf.DB()

	d := &models.Demo{Name: c.Query("p")}
	d.No = "yes"
	if err := db.Save(d).Error; err != nil {
		return err
	}

	return nil
}
