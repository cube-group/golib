package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/urls"
)

func main(){
	var requestUrl = "https://abc.com:8088/aa/bb/cc.html?a=1#asdfasdf"
	fmt.Println(urls.AddQuery(requestUrl,"a=123&b=1&c=2&d=3"))
	fmt.Println(urls.AddQueryMap(requestUrl,gin.H{"w":2,"h":"asdfa"}))
}
