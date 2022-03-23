package ms

import (
	"fmt"
	"github.com/imroc/req"
	"github.com/cube-group/golib/crypt/md5"
	"strings"
	"time"
)

type MsReq struct {
	ak string
	sk string
}

func NewMsReq(ak, sk string) *MsReq {
	return &MsReq{
		ak: ak,
		sk: sk,
	}
}

func (t *MsReq) authInfo() req.QueryParam {
	timestamp := time.Now().Unix()
	return req.QueryParam{
		"ak":   t.ak,
		"t":    timestamp,
		"sign": md5.MD5(fmt.Sprintf("%s%d%s", t.ak, timestamp, t.sk)),
	}
}

func (t *MsReq) Do(method, url string, vs ...interface{}) (*req.Resp, error) {
	method = strings.ToUpper(method)
	return req.Do(method, url, append(vs, t.authInfo())...)
}
