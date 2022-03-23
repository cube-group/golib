package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Base("./"), path.Dir("./"))
	fmt.Println(path.Base("./abc"), path.Dir("./abc"))
}
