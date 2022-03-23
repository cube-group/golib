package analysis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/env"
	"github.com/cube-group/golib/types/jsonutil"
	"github.com/cube-group/golib/types/times"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

type LogFileJson struct {
	//分析名称
	Name string
	//本地日志目录
	LogPath string
	//唯一uuid属性
	_ruid string
	//当前正在操作日志的指针实例
	_logger *log.Logger
	//当前正在写入日志的文件指针
	_currentFile *os.File
}

//初始化文件类日志系统
func (t *LogFileJson) initLog() error {
	if t.Name == "" {
		return errors.New("LogName is nil.")
	}
	if t.LogPath == "" {
		t.LogPath = "/data/log/"
	}
	now := time.Now()
	newName := fmt.Sprintf(
		"%s/%s.analysis.%s",
		path.Dir(t.LogPath),
		now.Format("2006-01-02"),
		t.Name,
	)
	if t._ruid == "" {
		t._ruid = md5.MD5(fmt.Sprintf("%v-%v", rand.Intn(10000), time.Now().Nanosecond()))
	}

	f, err := os.OpenFile(newName, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			if f2, err := os.Create(newName); err == nil {
				t._currentFile = f2
				goto SUCCESS
			}
		}
	} else {
		t._currentFile = f
		goto SUCCESS
	}

	t._logger = log.New(os.Stdout, "", 0)
	return nil

SUCCESS:
	t._logger = log.New(t._currentFile, "", 0)
	return nil
}

func (t *LogFileJson) Puts(values []map[string]interface{}) error {
	if values == nil {
		return errors.New("log values is nil.")
	}
	if err := t.initLog(); err != nil {
		return err
	}

	for _, item := range values {
		if err := t.Put(item); err != nil {
			return err
		}
	}

	return nil
}

//记录本地文本类日志
func (t *LogFileJson) Put(v gin.H) error {
	if v == nil {
		return errors.New("log values is nil.")
	}
	if err := t.initLog(); err != nil {
		return err
	}
	if _, ok := v["env"]; !ok {
		v["env"] = env.GetString("APP_MODE", "")
	}
	if _, ok := v["uuid"]; !ok {
		v["uuid"] = t._ruid
	}
	if _, ok := v["time"]; !ok {
		v["time"] = time.Now().Unix()
	}
	if _, ok := v["micro"]; !ok {
		v["micro"] = time.Now().UnixNano() / 1e6
	}
	if _, ok := v["date"]; !ok {
		v["date"] = times.FormatDatetime(time.Now())
	}
	t._logger.Println(jsonutil.ToString(v))
	return nil
}
