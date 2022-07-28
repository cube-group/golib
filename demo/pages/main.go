package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/page"
)

type Status struct {
	Phone       string `gorm:""`
	AgreementId uint   `gorm:""`
}

func (t *Status) TableName() string {
	return "occ_agreement_status"
}

func init() {
	conf.Init(conf.Template{
		AppYamlPath: "https://xx.com/application.yaml",
	})
}

func main() {
	server.Create(server.Config{
		FuncController: routes,
	})
}

func routes(engine *gin.Engine) {
	engine.GET("/list", list)
}

func list(c *gin.Context) {
	var finds []*Status
	res, err := page.List(c, &finds, conf.DB().Where("phone=?", "12345678901"))
	ginutil.JsonAuto(c, "ok", err, res)
}
