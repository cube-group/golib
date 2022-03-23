package main

import (
	"fmt"
)

type A struct {
	Name string
}
func main() {
	var i *A
	defer fmt.Println("=>1", i)
	defer func() {
		fmt.Println("=>2", i)
	}()
	defer func(v *A) {
		fmt.Println("=>3", v)
	}(i)
	i = &A{Name:"A"}
	defer fmt.Println("=>4", i)
	defer func() {
		fmt.Println("=>5", i)
	}()
	return
}
