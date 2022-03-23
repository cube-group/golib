package middlewares

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/log"
	"os"
	"time"
)

func Recovery(appName, sentryDsn string) gin.HandlerFunc {
	log.StdOut("MiddleWare", "Recovery Init")

	if sentryDsn != "" {
		log.StdOut("Sentry", sentryDsn, sentry.Init(sentry.ClientOptions{
			Environment:  "APP_MODE=" + os.Getenv("APP_MODE"),
			Release:      appName + "@" + os.Getenv("version"),
			IgnoreErrors: []string{"write tcp"}, //过滤掉你不感兴趣的
			Dsn:          sentryDsn,
		}))
	}

	return func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				if sentryDsn != "" {
					sentry.CaptureException(errors.New(fmt.Sprintf("%v", e)))
					sentry.Flush(time.Second)
				}
				log.StdErr("Recovery.Gin", e)
			}
		}()

		c.Next()
	}
}
