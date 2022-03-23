package main

import (
	"flag"
	"fmt"
)

func main() {
	var cmd string
	flag.StringVar(&cmd, "cmd", "web", "cmd")
	flag.Parse()
	fmt.Println(cmd)
}
