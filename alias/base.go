package alias

import (
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type MSF map[string]interface{}
type MSS map[string]string
type DateTime time.Time //日期时间
type Date time.Time     //日期
type AnyInt int         //允许解析整形字符串
type AnyString string   //允许解析数值为字符串

func (dt *DateTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*dt).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

func (dt *DateTime) UnmarshalJSON(text []byte) error {
	stamp, err := time.ParseInLocation("2006-01-02 15:04:05", string(text[1:20]), time.Local)
	if err != nil {
		return err
	}
	reflect.ValueOf(dt).Elem().Set(reflect.ValueOf(DateTime(stamp)))
	return err
}

func (dt *DateTime) String() string {
	return time.Time(*dt).Format("2006-01-02 15:04:05")
}

func (d *Date) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(*d).Format("2006-01-02"))
	return []byte(stamp), nil
}

func (d *Date) UnmarshalJSON(text []byte) error {
	stamp, err := time.ParseInLocation("2006-01-02", string(text[1:11]), time.Local)
	if err != nil {
		return err
	}
	reflect.ValueOf(d).Elem().Set(reflect.ValueOf(Date(stamp)))
	return err
}

func (d *Date) String() string {
	return time.Time(*d).Format("2006-01-02")
}

func (ai *AnyInt) UnmarshalJSON(text []byte) error {
	l := len(text)
	if text[0] == '"' && text[l-1] == '"' {
		text = text[1 : l-1]
	}
	if len(text) == 0 {
		return nil
	}
	v, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}
	reflect.ValueOf(ai).Elem().Set(reflect.ValueOf(AnyInt(v)))
	return err
}

func (t *AnyString) UnmarshalJSON(text []byte) error {
	l := len(text)
	if text[0] == '"' && text[l-1] == '"' {
		text = text[1 : l-1]
	}
	reflect.ValueOf(t).Elem().Set(reflect.ValueOf(AnyString(text)))
	return nil
}
