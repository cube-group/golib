package main

import (
	"fmt"
	"github.com/cube-group/golib/types/slice"
)

func main(){

	var find uint8 = 3
	var values  = []uint8{2,3,4}

	fmt.Println(slice.InArray(find,values))
}
