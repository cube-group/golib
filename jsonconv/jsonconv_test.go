package jsonconv

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestJsonCamelCase_MarshalJSON(t *testing.T) {
	type Category struct {
		ID            uint32   `json:"id"`
		Name          string   `json:"name"`
		ParentID      uint32   `json:"parent_id"`
		Level         uint32   `json:"level"`
		TeacherIdList []uint32 `json:"teacher_id_list"`
	}
	var a = Category{}
	res, _ := json.Marshal(a)
	log.Println("正常格式", string(res))

	camelCase, _ := json.Marshal(JsonCamelCase{a})
	log.Println("驼峰格式", string(camelCase))
}

func TestJsonSnakeCase_MarshalJSON(t *testing.T) {

	type Person struct {
		LightWeightBaby string
	}
	type Test struct {
		MyName         string `json:"myName"`
		HelloWorld     string `json:"helloWorld"`
		Person         Person
		HelloWorldABCD string
	}
	var a, b Test
	a.MyName = `{"aName":"bName"}`
	b.MyName = "helloWorldABCD"
	list := []Test{b, a}
	res, _ := json.Marshal(list)
	log.Println("正常无tag格式", string(res))

	snakeRes, _ := json.Marshal(JsonSnakeCase{list})
	log.Println("下划线格式", string(snakeRes))
}

//驼峰转下划线
func TestCase2Camel(t *testing.T) {
	fmt.Println(Case2Camel("light_weigh_baby"))
}

//驼峰转下划线
func TestCamel2Case(t *testing.T) {
	fmt.Println(Camel2Case("helloWorld"))
}
