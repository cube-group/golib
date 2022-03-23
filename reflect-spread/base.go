package reflectSpread

import (
	"reflect"
	"strconv"
)

func GetReflectValue(m interface{}) (res []reflect.Value) {
	val := reflect.ValueOf(m)
	if val.Kind() == reflect.Ptr {
		// 如果是指针获取其真实数据
		val = val.Elem()
	}
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			v := val.Index(i)
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			res = append(res, v)
		}
	} else {
		res = append(res, val)
	}
	return
}

func GetFieldByTag(v reflect.Type, tagName, tagVal string) string {
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if t, ok := f.Tag.Lookup(tagName); ok {
			if t == tagVal {
				return f.Name
			}
		}
	}
	return ""
}

func InterfaceToString(v interface{}) string {
	switch val := v.(type) {
	case int:
		return strconv.Itoa(val)
	case int8:
		return strconv.Itoa(int(val))
	case int16:
		return strconv.Itoa(int(val))
	case int32:
		return strconv.Itoa(int(val))
	case int64:
		return strconv.FormatInt(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'E', -1, 32)
	case float64:
		return strconv.FormatFloat(val, 'E', -1, 64)
	case string:
		return val
	default:
		return ""
	}
}
