package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/conf"
)

func init() {
	var core = conf.Init(conf.Template{
		AppYamlPath:         "./",
		AppYamlPathChildren: map[string]string{"abc": "child.yaml"},
		AppYamlIgnoreRedis:  []string{"default"},
		AppYamlIgnoreMysql:  []string{"default"},
	})
	fmt.Println(core.Viper("abc").GetString("abc.efg"))
	fmt.Println(core.Viper().GetString("server.name"))
	fmt.Println(core.Viper("abc").GetStringMap("abc"))
}

func main() {
	//fmt.Println(viper.Get("server"))            //config get map
	fmt.Println(viper.GetString("server.name")) //config get string
	//fmt.Println(viper.GetStringSlice("nsq.d"))  //config get string slice
	//fmt.Println(viper.GetInt("mysql.maxOpen"))  //config get int

	//conf.Redis().Set("a", 1, time.Hour)     //redis set
	//res, _ := conf.Redis().Get("a").Bytes() //redis get
	//fmt.Println(string(res))
	//
	//var a int64
	//conf.DB().Count(&a) //db count
	//
	//indexNames, err := conf.ES().IndexNames() //es select index
	//if err != nil {
	//	log.StdFatal("es", err.Error())
	//}
	//fmt.Println(indexNames)
	//fmt.Println(conf.DB("tss"))
	fmt.Println(conf.Redis("common"))
}
