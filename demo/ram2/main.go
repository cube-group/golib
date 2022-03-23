package main

import (
	"fmt"
	"github.com/cube-group/golib/cache"
	"time"
)

func main() {
	content := `{"id":61,"state":1,"type":0,"name":"题库付费批改","ak":"42839cd6f40c675269661373e19ad297","sk":"636c41cebae796045c8725e479be4cbc","gid":1,"update_time":"2020-03-23T17:31:49+08:00","create_time":"2020-03-23T17:31:49+08:00","ms":"","is_update_key":false,"is_update_state":false}`
	ramer := cache.NewRam()
	ramer.Set("key", []byte(content), time.Minute)

	fmt.Println(ramer.DebugInfo())

	time.Sleep(time.Hour)
}
