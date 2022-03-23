package main

import (
	"fmt"
	"github.com/cube-group/golib/cache"
	"time"
)

func main() {
	ram := cache.NewRam()
	ram.Set("a", time.Now().String(), 3*time.Second)
	time.Sleep(time.Second)
	fmt.Println("a:", ram.Get("a"))
	fmt.Println("length:", ram.Length())
	time.Sleep(4 * time.Second)
	fmt.Println("a:", ram.Get("a"))
	time.Sleep(time.Second)
	ram.Close()
	fmt.Println("length:", ram.Length())
}
