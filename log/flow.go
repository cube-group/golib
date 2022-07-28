package log

import (
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/ginutil"
	"github.com/cube-group/golib/types/jsonutil"
	"github.com/cube-group/golib/types/times"
	"github.com/cube-group/golib/uuid"
	"time"
)

const (
	LevelInfo         = "INFO"
	LevelWarn         = "WARN"
	LevelError        = "ERROR"                     //会触发报警
	LevelFatal        = "FATAL"                     //会触发报警
	LFlowIDKey        = "x-flow-id"          //上下文流程id暂存key
	LTcpLogApiAddress = "http://log.inner/business" //集群内网投递地址
)

type LLevel string

type LContent struct {
	Flow     string
	FlowID   string
	Duration int64
	Route    string
	Uid      string
	Code     int
	Msg      string
	Ext      gin.H
	Level    LLevel
	Context  *gin.Context
}

func (t *LContent) ToString() string {
	var now = time.Now()
	res := gin.H{
		"app":      getAppName(),
		"route":    t.Route,
		"uid":      t.Uid,
		"code":     t.Code,
		"msg":      t.Msg,
		"ext":      jsonutil.ToString(t.Ext),
		"level":    t.Level,
		"duration": t.Duration,
		"time":     now.Unix(),
		"micro":    now.UnixNano() / 1e6,
		"date":     times.FormatDatetime(now),
	}
	if t.Flow != "" {
		if t.FlowID == "" && t.Context != nil {
			t.FlowID = FlowID(t.Context)
		}
		res["flow"] = t.Flow
		res["flowid"] = t.FlowID
	}
	return jsonutil.ToString(res)
}

//从gin web上下文中获取流程链ID
func FlowID(c *gin.Context) (res string) {
	if c != nil {
		key, ok := c.Get(LFlowIDKey)
		if ok {
			res = key.(string)
		}
		if res != "" {
			return
		}

		res = ginutil.Input(c, LFlowIDKey)
		if res == "" {
			res = c.GetHeader(LFlowIDKey)
		}
	}
	if res == "" {
		res = md5.MD5(uuid.GetUUID())
	}
	if c != nil {
		c.Set(LFlowIDKey, res)
	}
	return
}

//写入标准日志
func Push(v LContent) {
	go filePush(v.ToString())
}
