package ding

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"time"
)

var prefix = ""
var timeout = 2 * time.Second
var header = req.Header{
	"Content-Type": "application/json",
}

//发送钉钉报警
func Ding(ding string, params ...interface{}) error {
	param := req.Param{
		"msgtype": "text",
		"text":    req.Param{"content": prefix + fmt.Sprintf("%s", fmt.Sprint(params...))},
		"at":      req.Param{"isAtAll": true},
	}
	req.SetTimeout(timeout)
	r, err := req.Post(ding, header, req.BodyJSON(&param))
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := r.ToJSON(&result); err != nil {
		return err
	}
	if result["errcode"] != 0 {
		return errors.New(fmt.Sprint(result))
	}
	return nil
}
