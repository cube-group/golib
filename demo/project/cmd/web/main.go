package web

import (
	"app/library/g"
	"app/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/env"
	"github.com/cube-group/golib/ginutil/server"
	"github.com/cube-group/golib/log"
)

func Init() {
	server.Create(server.Config{
		AppName:            log.AppName,
		SentryDsn:          env.GetString("SENTRY_DSN", viper.GetString("sentry.dsn")),
		Address:            viper.GetString("server.address"),
		FuncController:     uses,
		FuncMap:            g.ViewFunc,
		UseRedisSession:    true,
		StaticRoot:         "public",
		StaticRelativePath: "/public", //例如/public/assets/js/common.js
		HtmlPattern:        "views/*/**",
	})
}

func uses(engine *gin.Engine) {
	//todo register routes
	routes.Init(engine)
}
