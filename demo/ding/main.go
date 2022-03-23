package main

import (
	"fmt"
	"github.com/cube-group/golib/net/ding"
)

func main() {
	fmt.Println(ding.Ding("https://dingtalk.com/xxx", "hello", "world"))
}
