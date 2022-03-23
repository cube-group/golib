package g

import (
	"html/template"
)

//模板全局函数
var ViewFunc = template.FuncMap{
	"demo": demo,
}

func demo(v string) string {
	return v + "-viewfunc"
}
