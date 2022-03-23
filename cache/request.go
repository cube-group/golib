package cache

import (
	"encoding/json"
	"errors"
	"github.com/cube-group/golib/conf"
	"time"
)

type Request func() (interface{}, error)

//请求自动添加缓存
func RequestWithCache(cacheKey string, cacheExpire time.Duration, f Request, target interface{}, security ...bool) error {
	isSecurity := false

	if cacheExpire > 0 { //必须存在缓存时效
		if len(security) > 0 && security[0] {
			isSecurity = true
		}
		cacheRes, err := conf.Redis().Get(cacheKey).Bytes()
		if err == nil {
			if string(cacheRes) == "SECURITY_NIL" { //命中击穿
				return errors.New("security nil")
			}
			if err := json.Unmarshal(cacheRes, target); err == nil {
				return nil
			}
		}
	}

	res, err := f()
	if err != nil {
		if isSecurity { //防止缓存击穿
			go conf.Redis().Set(cacheKey, "SECURITY_NIL.", 15*time.Second)
		}
		return err
	}

	var savedBytes []byte
	switch res.(type) {
	case []byte:
		savedBytes = res.([]byte)
	case string:
		savedBytes = []byte(res.(string))
	default:
		bytes, err := json.Marshal(res)
		if err != nil {
			if isSecurity { //防止缓存击穿
				go conf.Redis().Set(cacheKey, "SECURITY_NIL.", 15*time.Second)
			}
			return err
		}
		savedBytes = bytes
	}

	if cacheExpire > 0 {//必须存在缓存时效
		go conf.Redis().Set(cacheKey, savedBytes, cacheExpire)
	}

	json.Unmarshal(savedBytes, target)
	return nil
}
