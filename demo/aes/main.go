package main

import (
	"fmt"
	"github.com/cube-group/golib/crypt/aes"
	"time"
)

func main() {
	content := "kajldkfjaldjflajf拉科技代理费建安费123123131231231%123jlasld::12313212:123123"
	key := "123456789012345678901234"
	var start = time.Now().UnixNano()
	encodeStr, _ := aes.AesEncrypt([]byte(key), content)
	fmt.Println(encodeStr, (time.Now().UnixNano()-start)/1e3, "微秒")
	start = time.Now().UnixNano()
	decodeStr, _ := aes.AesDecrypt([]byte(key), encodeStr)
	fmt.Println(decodeStr, (time.Now().UnixNano()-start)/1e3, "微秒")
}
