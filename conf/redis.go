package conf

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/types/convert"
	"github.com/cube-group/golib/types/slice"
)

var _listRedis = make(map[string]*redis.Client)

func initRedis(vip *viper.Viper, ignores []string, use []string) {
	res := vip.GetStringMap("redis")
	if len(res) == 0 {
		return
	}

	if _, ok := res["address"]; !ok { //多连接
		for name, value := range res {
			if len(use) > 0 && !slice.InArrayString(name, use) {
				log.StdOut("Conf", "Redis.Ignore", name)
				continue
			}
			if ignores != nil && slice.InArrayString(name, ignores) {
				log.StdOut("Conf", "Redis.Ignore", name)
				continue
			}
			createRedis(name, value)
		}
	} else { //单连接
		name := "default"
		if ignores != nil && slice.InArrayString(name, ignores) {
			log.StdOut("Conf", "Redis.Ignore", name)
		} else if len(use) > 0 && !slice.InArrayString(name, use) {
			log.StdOut("Conf", "Redis.Ignore", name)
		} else {
			createRedis(name, res)
		}
	}
}

func createRedis(name string, values interface{}) {
	cfg := values.(map[string]interface{})
	conn := redis.NewClient(&redis.Options{
		Addr:     convert.MustString(cfg["address"]),
		Password: convert.MustString(cfg["password"]),
		DB:       convert.MustInt(cfg["db"]),
		PoolSize: convert.MustInt(cfg["poolsize"]),
	})
	if err := conn.Ping().Err(); err != nil {
		log.StdFatal("Conf", "Redis", name, err.Error())
	} else {
		_listRedis[name] = conn
		log.StdOut("Conf", "Redis", name)
	}
}

func Redis(name ...string) *redis.Client {
	realName := "default"
	if len(name) > 0 {
		realName = name[0]
	}
	if i, ok := _listRedis[realName]; ok {
		return i
	}
	return nil
}

func GetRedisObject(cacheKey string, target interface{}, name ...string) error {
	redisName := "default"
	if len(name) > 0 {
		redisName = name[0]
	}
	bytes, err := Redis(redisName).Get(cacheKey).Bytes()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bytes, target); err != nil {
		return err
	}
	return nil
}
