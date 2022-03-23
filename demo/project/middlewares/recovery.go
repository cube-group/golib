package middlewares

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/log"
	"os"
	"time"
)

func Recovery(c *gin.Context) {
	defer func() {
		if e := recover(); e != nil {
			if dsn := os.Getenv("APP_SENTRY"); dsn != "" {
				sentry.CaptureException(errors.New(fmt.Sprintf("%v", e)))
				sentry.Flush(time.Second)
			} else {
				log.StdErr("Global", "Recover", e)
			}
			ginutil.JsonError(c, fmt.Sprintf("%v", e), nil, 999)
		}
	}()

	c.Next()
}
