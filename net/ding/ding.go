package ding

import (
	"errors"
	"fmt"
	"github.com/imroc/req"
	"time"
)

//发送钉钉报警
func Ding(url string, params ...interface{}) error {
	param := req.Param{
		"msgtype": "text",
		"text":    req.Param{"content": fmt.Sprintf("%s", fmt.Sprint(params...))},
		"at":      req.Param{"isAtAll": true},
	}
	req.SetTimeout(2 * time.Second)
	r, err := req.Post(url, req.Header{"Content-Type": "application/json"}, req.BodyJSON(&param))
	if err != nil {
		return err
	}

	var result map[string]interface{}
	if err := r.ToJSON(&result); err != nil {
		return err
	}
	if fmt.Sprintf("%v", result["errcode"]) != "0" {
		return errors.New(fmt.Sprint(result))
	}
	return nil
}
