package main

import (
	"github.com/cube-group/golib/log"
	"time"
)

func main() {
	log.StdOut("Main", "Hello", time.Now())
	log.StdWarning("Main", "Hello", time.Now())
	log.StdErr("Main", "Hello", time.Now())
	log.StdFatal("Main", "Hello", time.Now()) //会直接退出进程
}
