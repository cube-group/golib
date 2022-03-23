package conf

import (
	"errors"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/types/convert"
	"github.com/cube-group/golib/types/slice"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var _listMysql = make(map[string]*gorm.DB)

func initMysql(vip *viper.Viper, ignores []string, use []string) {
	res := vip.GetStringMap("mysql")
	if len(res) == 0 {
		return
	}

	if _, ok := res["address"]; !ok { //多连接
		for name, value := range res {
			if len(use) > 0 && !slice.InArrayString(name, use) {
				log.StdOut("Conf", "Mysql.Ignore", name)
				continue
			}
			if ignores != nil && slice.InArrayString(name, ignores) {
				log.StdOut("Conf", "Mysql.Ignore", name)
				continue
			}
			createMysql(name, value)
		}
	} else { //单连接
		name := "default"
		if ignores != nil && slice.InArrayString(name, ignores) {
			log.StdOut("Conf", "Redis.Ignore", name)
		} else if len(use) > 0 && !slice.InArrayString(name, use) {
			log.StdOut("Conf", "Redis.Ignore", name)
		} else {
			createMysql(name, res)
		}
	}
}

func createMysql(name string, cfg interface{}) {
	maps := cfg.(map[string]interface{})
	config := &gorm.Config{}
	if convert.MustBool(maps["logmode"]) { //info
		config.Logger = logger.Default.LogMode(logger.Info)
	} else { //silent
		config.Logger = logger.Default.LogMode(logger.Silent)
	}
	conn, err := gorm.Open(mysql.Open(maps["address"].(string)), config)
	if err != nil {
		log.StdFatal("Conf", "Mysql.Open", name, err.Error())
		return
	}
	sqlDB, err := conn.DB()
	if err != nil {
		log.StdFatal("Conf", "Mysql.Conn", name, err.Error())
		return
	}
	if maxIdle, ok := maps["maxidle"]; ok {
		sqlDB.SetMaxIdleConns(convert.MustInt(maxIdle))
	}
	if maxOpen, ok := maps["maxopen"]; ok {
		sqlDB.SetMaxOpenConns(convert.MustInt(maxOpen))
	}
	_listMysql[name] = conn
	log.StdOut("Conf", "Mysql", name)
}

func DB(name ...string) *gorm.DB {
	realName := "default"
	if len(name) > 0 {
		realName = name[0]
	}
	if i, ok := _listMysql[realName]; ok {
		return i
	}
	return nil
}

//错误是否为
func DBNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
