package main

import (
	"fmt"
	"github.com/cube-group/golib/bas"
)

func main() {
	token, useTime, err := bas.NewBasService().GetTokenFromHttp("1", "42839cd6f40c675269661373e19ad297", 15)
	fmt.Println(token)
	fmt.Println(useTime, err)
}
