package main

import (
	"fmt"
	"sync"
	"time"
)

type A struct {
	locker sync.RWMutex
}

var a = new(A)
var v string

func main() {
	go do()

	a.locker.Lock()
	time.Sleep(time.Second * 5)
	v = "abc"
	a.locker.Unlock()
	time.Sleep(time.Hour)
}

func do() {
	time.Sleep(time.Second)
	defer a.locker.RUnlock()
	a.locker.RLock()
	fmt.Println(v)
}
