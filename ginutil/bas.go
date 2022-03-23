package ginutil

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/bas/consts"
	"github.com/cube-group/golib/crypt/md5"
)

//检测信任通道
//检测是否来自统一认证系统
func Gid(c *gin.Context) (string, string, error) {
	ak := c.GetHeader(consts.AUTH_FORWARDED_AK)
	gid := c.GetHeader(consts.AUTH_FORWARDED_GID)
	md5V := c.GetHeader(consts.AUTH_FORWARDED_MD5)
	timestamp := c.GetHeader(consts.AUTH_FORWARDED_TIME)
	if gid == "" || ak == "" || timestamp == "" || md5V == "" {
		return "", "", errors.New("distrusted")
	}
	return gid, ak, nil
}

//严格检测信任通道
//严格检测是否来自统一认证系统
func GidStrict(c *gin.Context) (string, string, error) {
	ak := c.GetHeader(consts.AUTH_FORWARDED_AK)
	gid := c.GetHeader(consts.AUTH_FORWARDED_GID)
	md5V := c.GetHeader(consts.AUTH_FORWARDED_MD5)
	timestamp := c.GetHeader(consts.AUTH_FORWARDED_TIME)
	if gid == "" || ak == "" || timestamp == "" || md5V == "" {
		return "", "", errors.New("distrusted")
	}
	if md5.MD5(fmt.Sprintf("gid=%s&t=%s", gid, timestamp)) != md5V {
		return "", "", errors.New("timestamp invalid.")
	}
	return gid, ak, nil
}
