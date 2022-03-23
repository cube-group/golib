package ginSpread

import (
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type ResStruct struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type MSG string

var Code = map[int]string{
	0:     "OK",
	404:   "请求的接口不存在",
	10001: "参数校验不通过",
	10002: "请求无返回结果",
	20001: "鉴权失败",
	20002: "该KEY无请求权限",
	20003: "鉴权已过期",
	20006: "服务器错误",
}

func ShouldBind(c *gin.Context, obj interface{}) error {
	var err error
	if strings.Contains(c.GetHeader("Content-Type"), "application/json") {
		err = c.ShouldBindJSON(obj)
	} else {
		err = c.ShouldBind(obj)
	}
	return err
}

func Success(c *gin.Context, args ...interface{}) {
	out := &ResStruct{
		Code: 0,
		Msg:  Code[0],
	}
	if len(args) > 2 {
		args = args[:2]
	}
	for _, v := range args {
		switch v := v.(type) {
		case MSG:
			if out.Msg == Code[0] {
				out.Msg = string(v)
			} else {
				out.Data = v
			}
		default:
			out.Data = v
		}
	}
	c.Set("code", out.Code)
	c.Set("msg", out.Msg)
	c.Set("ext", out.Data)
	c.JSON(200, out)
	c.Abort()
	return
}

func Error(c *gin.Context, args ...interface{}) {
	out := &ResStruct{
		Code: 20006,
		Msg:  Code[20006],
	}
	if len(args) > 3 {
		args = args[:3]
	}
	for _, v := range args {
		switch v := v.(type) {
		case int:
			out.Code = v
			out.Msg = Code[v]
		case error:
			log.Print(v)
			if out.Msg == "" || out.Msg == Code[out.Code] {
				out.Msg = v.Error()
			} else {
				out.Data = v.Error()
			}
		case string:
			if out.Msg == "" || out.Msg == Code[out.Code] {
				out.Msg = v
			} else {
				out.Data = v
			}
		default:
			out.Data = v
		}
	}
	c.Set("code", out.Code)
	c.Set("msg", out.Msg)
	c.Set("ext", out.Data)
	c.JSON(200, out)
	c.Abort()
	return
}

func Auto(c *gin.Context, data interface{}, err error) {
	if err != nil {
		Error(c, err.Error())
		return
	}
	Success(c, data)
	return
}
