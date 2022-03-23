package uuid

import (
	"fmt"
	"github.com/cube-group/golib/crypt/md5"
	"math/rand"
	"time"
)

//获取唯一md5id
func GetUUID() string {
	return md5.MD5(fmt.Sprintf("%d%d", rand.Int63n(1000), time.Now().Nanosecond()))
}
