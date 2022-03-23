package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/cube-group/golib/env"
	"github.com/cube-group/golib/types/jsonutil"
	"github.com/cube-group/golib/types/times"
	"os"
	"path"
	"reflect"
	"sync"
	"time"
)

var _locker sync.Mutex    //全局锁
var _currentFile *os.File //当前正在写入日志的文件指针

var AppName = ""        //默认应用名称
var Path = "/data/log/" //默认日志目录
var Extension = "json"  //默认日志文件扩展名

func init() {
	initLog()
}

//初始化文件类日志系统
func initLog() error {
	now := time.Now()
	newName := fmt.Sprintf("%s/golib-%s.%s", path.Dir(Path), times.FormatDate(now), Extension)

	//log ready
	if _currentFile != nil {
		if _currentFile.Name() != newName {
			_locker.Lock()
			_currentFile.Close()
		} else if fileInfo, err := os.Stat(_currentFile.Name()); err != nil {
			_locker.Lock()
			_currentFile.Close()
		} else {
			fileInfo.Size()
			return nil
		}
	} else {
		_locker.Lock()
	}

	//lock log file to ready
	defer _locker.Unlock()
	if f, err := os.OpenFile(newName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644); err != nil {
		return err
	} else {
		_currentFile = f
	}
	return nil
	//
	//	_logger = log.New(os.Stdout, "", 0)
	//	return err
	//
	//SUCCESS:
	//	_logger = log.New(_currentFile, "", 0)
	//	return err
}

func filePush(contents string) {
	if err := initLog(); err != nil {
		fmt.Println(contents)
	} else {
		_locker.Lock()
		defer _locker.Unlock()
		if _, err := _currentFile.WriteString(contents + "\n"); err != nil {
			fmt.Println(err)
			fmt.Println(contents)
		}
	}
}

//记录本地文本类Info日志
func FileLogInfo(route, uid, code, msg, ext interface{}) {
	var now = time.Now()
	var contents = jsonutil.ToString(gin.H{
		"date":  times.FormatDatetime(now),
		"time":  now.Unix(),
		"micro": now.UnixNano() / 1e6,
		"level": LevelInfo,
		"app":   getAppName(),
		"route": route,
		"uid":   uid,
		"code":  code,
		"msg":   fmt.Sprintf("%v", msg),
		"ext":   getExt(ext),
	})
	filePush(contents)
}

//记录本地文件类错误日志
func FileLogErr(route, uid, code, msg, ext interface{}) {
	var now = time.Now()
	var contents = jsonutil.ToString(gin.H{
		"date":  times.FormatDatetime(now),
		"time":  now.Unix(),
		"micro": now.UnixNano() / 1e6,
		"level": LevelError,
		"app":   getAppName(),
		"route": route,
		"uid":   uid,
		"code":  code,
		"msg":   fmt.Sprintf("%v", msg),
		"ext":   getExt(ext),
	})
	filePush(contents)
}

func getExt(ext interface{}) interface{} {
	if reflect.TypeOf(ext).String() == "string" {
		return ext
	} else {
		return jsonutil.ToString(ext)
	}
}

func getAppName() string {
	if AppName == "" {
		return env.GetString("APP_NAME", "nil")
	}
	return AppName
}
