package main

import (
	"fmt"
	"github.com/cube-group/golib/task"
	"time"
)

func main() {
	a := task.NewAsync()
	a.Run(
		func() interface{} {
			time.Sleep(time.Second)
			return "a"
		}, func() interface{} {
			return "b"
		},
	)
	fmt.Println(a.Results)
}
