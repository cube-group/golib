package main

import (
	"fmt"
	"github.com/cube-group/golib/validator"
	"net/url"
	"path"
)

func main() {

	url1 := "http://abc.ciom/html"
	url2 := ""

	u1, _ := url.Parse(url1)
	u2, _ := url.Parse(url2)
	fmt.Println(u1.Scheme, u2.Scheme)

	fmt.Println(path.IsAbs(url1), path.IsAbs(url2))

	fmt.Println(validator.IsUrl(url1), validator.IsUrl(url2))

}
