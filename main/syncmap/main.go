package main

import (
	"fmt"
	"sync"
)

func main(){
	var maps sync.Map
	maps.Store("b",2)
	maps.Store("c",3)
	maps.Store("a",1)
	maps.Store("1",1)
	maps.Range(func(key, value interface{}) bool {
		fmt.Println(key,":",value)
		return true
	})
}
