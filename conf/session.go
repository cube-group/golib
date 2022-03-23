package conf

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/ginutil/session"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/viperutil"
)

func InitRedisSession(engine *gin.Engine) error {
	res := viper.GetStringMap("session")
	if len(res) == 0 {
		log.StdWarning("RedisSession", "No Viper Config")
		return errors.New("RedisSession No Viper Config")
	}

	address := viper.GetString("session.address")
	if address == "" {
		log.StdWarning("RedisSession", "session.address invalid")
		return errors.New("RedisSession No Viper Config Address")
	}
	sid := viper.GetString("session.sid")
	if sid == "" {
		log.StdWarning("RedisSession", "session.sid Recommended setting")
	}
	maxAge := viperutil.GetInt("session.max_age", "session.maxAge")
	if maxAge == 0 {
		maxAge = 24 * 3600
		log.StdWarning("RedisSession", "session.max_age|maxAge Recommended setting, the default is 24 hours")
	}
	secret := viper.GetString("session.secret")
	if secret == "" {
		log.StdWarning("RedisSession", "session.secret Setting is recommended for safety")
	}
	password := viper.GetString("session.password")
	if password == "" {
		log.StdWarning("RedisSession", "session.password Setting is recommended for safety")
	}
	poolSize := viperutil.GetInt("pool_size", "poolSize", "poolsize")
	if poolSize == 0 {
		poolSize = 40
		log.StdWarning("RedisSession", "session.pool_size|poolSize Recommended setting, the default is 40")
	}
	db := viper.GetInt("session.db")
	engine.Use(session.GinHandler(&session.Options{
		SessionName:   sid,
		MaxAge:        maxAge,
		Secret:        secret,
		RedisAddress:  address,
		RedisPassword: password,
		RedisDB:       db,
		RedisPoolSize: poolSize,
	}))
	return nil
}

func Session(c *gin.Context) session.IRedisSession {
	return session.Session(c)
}
