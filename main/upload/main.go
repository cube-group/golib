package main

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/cube-group/golib/crypt/md5"
	"time"
)

func main() {
	requestUrl := "http://ms.xx.com/cdn/upload/token"
	t := time.Now().Unix()
	sign := md5.MD5(fmt.Sprintf("secret=%s&t=%d", "SFASDFASF!@#!#!#!@#SDFASFASFAF", t))
	params := req.Param{
		"t":    t,
		"sign": sign,
	}
	resp, _ := req.Get(requestUrl, params)
	fmt.Println(resp.String())
}
