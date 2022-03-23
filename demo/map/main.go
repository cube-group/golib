package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/types/jsonutil"
)

func main() {
	fmt.Println(jsonutil.ToString(gin.H{"b": 2, "1": 2, "a": 1,"1a": 1}))
}
