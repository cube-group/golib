package routine

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
	"github.com/cube-group/golib/log"
	"os"
	"time"
)

type Func func()

//异步执行函数且进行错误捕获
func Async(f Func) {
	go Sync(f)
}

//同步执行函数且进行错误捕获
func Sync(f Func) {
	func() {
		defer func() {
			if e := recover(); e != nil {
				log.StdErr("Async", e)
			}
		}()
		f()
	}()
}

func Go(f Func) {
	func() {
		defer func() {
			if e := recover(); e != nil {
				sentryDsn := viper.GetString("sentry.dsn")
				if sentryDsn == "" {
					sentryDsn = os.Getenv("SENTRY_DSN")
				}
				if sentryDsn != "" {
					sentry.CaptureException(errors.New(fmt.Sprintf("%v", e)))
					sentry.Flush(time.Second)
				} else {
					log.StdErr("Go", e)
				}
			}
		}()
		f()
	}()
}
