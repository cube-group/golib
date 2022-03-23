package main

import (
	"fmt"
	"github.com/cube-group/golib/compress/zip"
	"os"
)

func main() {
	list := []string{"a.txt", "application.yaml"}
	files := make([]*os.File, 0)
	for _, v := range list {
		file, err := os.Open(v)
		if err != nil {
			continue
		}
		files = append(files, file)
	}
	dist := "./dist.zip"
	fmt.Println(zip.Compress(files, dist))
}
