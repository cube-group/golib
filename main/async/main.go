package main

import (
	"fmt"
	"github.com/cube-group/golib/routine"
	"time"
)

func main(){
	routine.Async(func() {
		fmt.Println("a")
		panic("b")
	})

	time.Sleep(time.Minute)
}
