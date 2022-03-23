package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	dataFormat "github.com/cube-group/golib/data-format"
)

func main() {
	a := map[string]interface{}{
		"a": 123,
		"b": "xxx",
		"c": map[string]interface{}{
			"d": 456,
			"e": "fff",
		},
		"f": "456",
		"g": "",
	}
	b, _ := jsoniter.Marshal(a)
	format := dataFormat.NewMapDataFormatFromJson(b)
	format.FormatInt(111, "a", "f", "g")
	fmt.Println(format.ToJsonString())
}
