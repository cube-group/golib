package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/jsonutil"
	"reflect"
)

type A struct {
	Name string
}

func main() {
	list:=[]interface{}{
		1,
		jsonutil.ToString(gin.H{"a": 1}),
		jsonutil.ToString(1),
		gin.H{"a": 1},
		map[string]interface{}{"a":1},
	}
	for _,i:=range list{
		fmt.Println(reflect.TypeOf(i).String())
	}

}
