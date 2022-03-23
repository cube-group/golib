package conf

import (
	"github.com/olivere/elastic"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/types/convert"
	"github.com/cube-group/golib/types/slice"
	log2 "log"
	"os"
	"time"
)

var _listElasticSearch = make(map[string]*elastic.Client)

func initElasticSearch(vip *viper.Viper, ignores, use []string) {
	res := vip.GetStringMap("es")
	if len(res) == 0 {
		return
	}

	if _, ok := res["address"]; !ok { //多连接
		for name, value := range res {
			if len(use) > 0 && !slice.InArrayString(name, use) {
				log.StdOut("Conf", "Es.Ignore", name)
				continue
			}
			if ignores != nil && slice.InArrayString(name, ignores) {
				log.StdOut("Conf", "Es.Ignore", name)
				continue
			}
			createElasticSearch(name, value)
		}
	} else { //单连接
		name := "default"
		if ignores != nil && slice.InArrayString(name, ignores) {
			log.StdOut("Conf", "Es.Ignore", name)
		} else if len(use) > 0 && !slice.InArrayString(name, use) {
			log.StdOut("Conf", "Es.Ignore", name)
		} else {
			createElasticSearch(name, res)
		}
	}
}

func createElasticSearch(name string, cfg interface{}) {
	maps := cfg.(map[string]interface{})
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(convert.MustString(maps["address"])),
		elastic.SetSniff(convert.MustBool(maps["sniff"])),
		elastic.SetHealthcheckInterval(10 * time.Second),
		elastic.SetGzip(convert.MustBool(maps["gzip"])),
		elastic.SetErrorLog(log2.New(os.Stderr, "ELASTIC ", log2.LstdFlags)),
		elastic.SetInfoLog(log2.New(os.Stdout, "", log2.LstdFlags)),
	}
	username, hasUsername := maps["username"]
	password, hasPassword := maps["password"]
	if hasUsername && hasPassword {
		options = append(options, elastic.SetBasicAuth(
			convert.MustString(username),
			convert.MustString(password),
		))
	}
	c, err := elastic.NewClient(options...)
	if err != nil {
		log.StdFatal("Conf", "Es", name, err.Error())
		return
	}
	_listElasticSearch[name] = c
	log.StdOut("Conf", "Es", name)
}

func ES(name ...string) *elastic.Client {
	realName := "default"
	if len(name) > 0 {
		realName = name[0]
	}
	if i, ok := _listElasticSearch[realName]; ok {
		return i
	}
	return nil
}
