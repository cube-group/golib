package cmd

import (
	"app/cmd/job"
	"app/cmd/web"
	"flag"
	"fmt"
	"github.com/cube-group/golib/conf"
	"github.com/cube-group/golib/env"
)

func init() {
	err := conf.Init(conf.Template{
		CpuNum:      env.GetInt("CPU_NUM", 1, 1, 16),
		ReqTimeout:  env.GetInt64("REQ_TIMEOUT", 5, 3, 15),
		AppYamlPath: env.GetString("APP_YAML_PATH", "."),
	})
	fmt.Println("here according to err, judge whether it is fatal", err)
}

//facade
func Execute() {
	var cmd string
	flag.StringVar(&cmd, "cmd", "web", "program start type web|job")
	flag.Parse()
	switch cmd {
	case "job":
		job.Init()
	case "web+job":
		go job.Init()
		web.Init()
	default:
		web.Init()
	}
}
