package main

import (
	"fmt"
	"github.com/cube-group/golib/crypt/base64"
	"github.com/cube-group/golib/types"
	"github.com/cube-group/golib/types/convert"
)

func main() {
	types.World()
	fmt.Println(convert.MustString(1234))
	base64.Base64Encode("abc")
}