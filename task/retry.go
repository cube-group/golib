package task

import (
	"fmt"
	"time"
)

//重试函数
func Retry(name string, tryTimes int, sleep time.Duration, callback func() error) (err error) {
	for i := 1; i <= tryTimes; i++ {
		err = callback()
		if err == nil {
			return nil
		}
		time.Sleep(sleep)
	}
	err = fmt.Errorf("[%v]失败，共重试%d次, 最近一次错误:%s \n", name, tryTimes, err)
	return err

}

//重试，限制时间
func RetryDurations(name string, max time.Duration, sleep time.Duration, callback func() error) (err error) {
	t0 := time.Now()
	i := 0
	for {
		err = callback()
		if err == nil {
			return
		}
		delta := time.Now().Sub(t0)
		if delta > max {
			fmt.Printf("[%v]失败，超过最大时间%s, 共重试%d次，最近一次错误: %s \n", name, max, i, err)
			return err
		}
		time.Sleep(sleep)
		i++
		fmt.Printf("[%v]失败，第%v次重试， 错误信息:%s \n", name, i, err)
	}
}
