package http

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"reflect"
)

func GetErrorInfo(response interface{}, err error) ([]byte, error) {
	if err != nil {
		return nil, err
	}
	resp := reflect.ValueOf(response)
	if resp.Kind() == reflect.Ptr {
		resp = resp.Elem()
	}
	code := resp.FieldByName("StatusCode").Int()
	body := resp.FieldByName("Body").Interface().(io.ReadCloser)
	defer body.Close()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("StatusCode %d \n%s", code, err.Error()))
	}
	if code < 200 || code > 299 {
		return nil, errors.New(string(b))
	}
	return b, nil
}
