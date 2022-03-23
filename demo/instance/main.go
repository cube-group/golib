package main

import "fmt"

type My struct {
	Name string
}

func  main(){
	var m My
	do(&m)
}

func do(m *My){
	m.Name = "Adsf"
	fmt.Println(m.Name)
}
