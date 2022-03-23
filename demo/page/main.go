package main

import (
	"fmt"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/page"
	"log"
)

type Status struct {
	Phone       string `gorm:""`
	AgreementId uint   `gorm:""`
}

func (t *Status) TableName() string {
	return "occ_agreement_status"
}
func main() {
	conf.Init(conf.Template{
		AppYamlPath: "https://eoffcn-software.oss-cn-beijing.aliyuncs.com/application.yaml",
	})

	var finds []*Status
	res, err := page.Vue(nil, &finds, conf.DB())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res,finds)
}
