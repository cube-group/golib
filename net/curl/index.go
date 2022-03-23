// Author: linyang
// Time: 2021-01
// 统一的http请求类
package curl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/log"
	"github.com/cube-group/golib/types/jsonutil"
	"github.com/cube-group/golib/urls"
	"github.com/cube-group/golib/uuid"
	"strings"
	"time"
)

const (
	REQ_TRACING_HEADER    = "x-tracing"
	REQ_TRACING_LOG_ROUTE = "x-tracing"
)

//默认curl超时时间
var ReqTimeout = 5 * time.Second

//是否开启日志自动记录
var ReqLog bool

func init() {
	req.SetTimeout(ReqTimeout)
}

func Req(requestUrl, method string, params ...interface{}) (*req.Resp, error) {
	var reqHeader req.Header
	var reqParam req.Param
	var reqBody string
	var requestChainId string
	for _, vv := range params {
		switch v := vv.(type) {
		case req.Header:
			reqHeader = v
		case req.Param:
			reqParam = v
		case string:
			reqBody = v
		case *gin.Context:
			requestChainId = ginutil.Input(v, REQ_TRACING_HEADER)
			if requestChainId == "" {
				requestChainId = v.GetHeader(REQ_TRACING_HEADER)
			}
		}
	}

	//准备请求链id
	if requestChainId == "" {
		requestChainId = md5.MD5(fmt.Sprint(uuid.GetUUID(), time.Now().String()))
	}
	requestUrl = urls.UrlAddQuery(requestUrl, gin.H{REQ_TRACING_HEADER: requestChainId})
	//send
	var startTime = time.Now().UnixNano()
	var r *req.Resp
	var err2 error
	switch strings.ToLower(method) {
	case "get":
		r, err2 = req.Get(requestUrl, params...)
	case "post":
		r, err2 = req.Post(requestUrl, params...)
	case "put":
		r, err2 = req.Put(requestUrl, params...)
	case "path":
		r, err2 = req.Patch(requestUrl, params...)
	case "delete":
		r, err2 = req.Delete(requestUrl, params...)
	default:
		err2 = fmt.Errorf("invalid method: %s ", method)
	}
	if ReqLog {
		defer func() { //business log
			go log.FileLogInfo(
				REQ_TRACING_LOG_ROUTE,
				requestChainId,
				0,
				fmt.Sprint(err2),
				jsonutil.ToString(map[string]interface{}{
					"url":      requestUrl,
					"method":   method,
					"header":   reqHeader,
					"params":   reqParam,
					"body":     reqBody,
					"duration": (time.Now().UnixNano() - startTime) / 1e6,
				}),
			)
		}()
	}
	if err2 != nil {
		return nil, err2
	}
	return r, nil
}

func Get(url string, params ...interface{}) (*req.Resp, error) {
	return Req(url, "get", params...)
}

func Post(url string, params ...interface{}) (*req.Resp, error) {
	return Req(url, "post", params...)
}
func Put(url string, params ...interface{}) (*req.Resp, error) {
	return Req(url, "put", params...)
}

func Patch(url string, params ...interface{}) (*req.Resp, error) {
	return Req(url, "patch", params...)
}

func Delete(url string, params ...interface{}) (*req.Resp, error) {
	return Req(url, "delete", params...)
}
