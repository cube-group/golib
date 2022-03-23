package conf

import (
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/imroc/req"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
)

var oncer sync.Once

//初始化配置模板
type Template struct {
	CpuNum              int               //应用程序所使用的cpu核数，默认为1核
	ReqTimeout          int64             //应用程序内对外发送请求的超时时间，单位：秒
	AppYamlIgnoreMysql  []string          //yaml中mysql忽略链接名称
	AppYamlIgnoreRedis  []string          //yaml中redis忽略链接名称
	AppYamlIgnoreEs     []string          //yaml中es忽略链接名称
	AppYamlUseMysql     []string          //yaml中mysql使用的链接名称
	AppYamlUseRedis     []string          //yaml中redis使用的链接名称
	AppYamlUseEs        []string          //yaml中es使用的链接名称
	AppYamlPath         string            //应用程序的application.yaml主配置文件路径
	AppYamlPathChildren map[string]string //应用程序的application.yaml子配置文件路径群
}

type Core struct {
	viper         *viper.Viper
	viperChildren map[string]*viper.Viper
}

// Get Viper Instance
func (t *Core) Viper(name ...string) *viper.Viper {
	var viperName string
	if len(name) > 0 {
		viperName = name[0]
	}
	if viperName == "" {
		return t.viper
	}
	if i, ok := t.viperChildren[viperName]; ok {
		return i
	}
	return nil
}

func Init(cfg Template) *Core {
	var err error
	var core = &Core{viper: viper.GetViper(), viperChildren: map[string]*viper.Viper{}}
	oncer.Do(func() {
		//req
		if cfg.ReqTimeout > 0 {
			req.SetTimeout(time.Duration(cfg.ReqTimeout) * time.Second)
		}

		//numCpu
		if cfg.CpuNum > 0 {
			runtime.GOMAXPROCS(cfg.CpuNum)
		}

		//file
		if cfg.AppYamlPath != "" || cfg.AppYamlPathChildren != nil {
			if err = initConfigFile(cfg, core); err != nil {
				log.StdFatal("Conf", "initConfigFile", err)
			}
			if core.viper == nil {
				return
			}
			//redis
			initRedis(core.viper, cfg.AppYamlIgnoreRedis, cfg.AppYamlUseRedis)
			//mysql
			initMysql(core.viper, cfg.AppYamlIgnoreMysql, cfg.AppYamlUseMysql)
			//elastic search
			initElasticSearch(core.viper, cfg.AppYamlIgnoreEs, cfg.AppYamlUseEs)
			//log name
			if logAppName := core.viper.GetString("server.name"); logAppName != "" {
				log.AppName = logAppName
			}
		}
	})

	return core
}
