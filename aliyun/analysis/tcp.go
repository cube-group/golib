package analysis

import (
	"encoding/json"
	"fmt"
	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/crypt/md5"
	"github.com/cube-group/golib/env"
	"github.com/cube-group/golib/types/convert"
	"github.com/cube-group/golib/types/jsonutil"
	"github.com/cube-group/golib/types/times"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type LogTcpJson struct {
	AliyunEndpoint        string
	AliyunAccessKeyId     string
	AliyunAccessKeySecret string

	//分析名称
	Name string
	//项目群
	ProjectName string
	//是否为任意logstore 默认为必须为analysis-xx
	AnyStore bool
	//唯一uuid属性
	_ruid string
	//当前正在操作日志的指针实例
	_producer *producer.Producer
	_client   *sls.Client
}

//初始化文件类日志系统
func (t *LogTcpJson) initLog() error {
	if t._producer != nil {
		return nil
	}

	if t.Name == "" {
		return errors.New("LogName is nil.")
	}
	if t.ProjectName == "" {
		t.ProjectName = PROJECT_NAME
	}
	if t._ruid == "" {
		t._ruid = md5.MD5(fmt.Sprintf("%v-%v", rand.Intn(10000), time.Now().Nanosecond()))
	}
	if t.AliyunEndpoint == "" {
		t.AliyunEndpoint = "cn-beijing.log.aliyuncs.com"
	}
	if t.AliyunAccessKeyId == "" || t.AliyunAccessKeySecret == "" {
		return errors.New("aliyun log access key config is nil")
	}
	if !t.AnyStore {
		t.Name = "analysis-" + t.Name
	}
	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.AllowLogLevel = "error"
	producerConfig.Endpoint = t.AliyunEndpoint
	producerConfig.AccessKeyID = t.AliyunAccessKeyId
	producerConfig.AccessKeySecret = t.AliyunAccessKeySecret
	t._producer = producer.InitProducer(producerConfig)
	t._producer.Start()
	return nil
}

func (t *LogTcpJson) initLog2() error {
	if t._client != nil {
		return nil
	}

	if t.AliyunAccessKeyId == "" || t.AliyunAccessKeySecret == "" {
		return errors.New("aliyun log access key config is nil")
	}
	if t.AliyunEndpoint == "" {
		t.AliyunEndpoint = "cn-beijing.log.aliyuncs.com"
	}
	t._client = &sls.Client{
		Endpoint:    t.AliyunEndpoint,
		AccessKeyID: t.AliyunAccessKeyId, AccessKeySecret: t.AliyunAccessKeySecret,
	}
	return nil
}

//初始化文件类日志系统
func (t *LogTcpJson) Close() error {
	if t._producer != nil {
		t._producer.SafeClose()
	}
	if t._client != nil {
		t._client.Close()
	}
	return nil
}

//记录Fast
func (t *LogTcpJson) Put(values map[string]interface{}) error {
	return t.Puts([]map[string]interface{}{values})
}

//记录Normal
func (t *LogTcpJson) Put2(values map[string]interface{}) error {
	return t.Puts2([]map[string]interface{}{values})
}

//记录本地文本类日志
func (t *LogTcpJson) Puts2(params []map[string]interface{}) error {
	if params == nil {
		return errors.New("log values is nil.")
	}
	if err := t.initLog2(); err != nil {
		return err
	}

	logs := []*sls.Log{}
	for _, item := range params {
		contents := []*sls.LogContent{}
		if _, ok := item["env"]; !ok {
			item["env"] = env.GetString("APP_MODE", "")
		}
		if _, ok := item["app"]; !ok {
			item["app"] = env.GetString("APP_NAME", "")
		}
		if _, ok := item["uuid"]; !ok {
			item["uuid"] = t._ruid
		}
		if _, ok := item["time"]; !ok {
			item["time"] = time.Now().Unix()
		}
		if _, ok := item["micro"]; !ok {
			item["micro"] = time.Now().UnixNano() / 1e6
		}
		if _, ok := item["date"]; !ok {
			item["date"] = times.FormatDatetime(time.Now())
		}
		for k, v := range item {
			switch reflect.ValueOf(v).Kind() {
			case reflect.Invalid:
				contents = append(contents, &sls.LogContent{Key: proto.String(k), Value: proto.String("")})
			case reflect.String, reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64:
				contents = append(contents, &sls.LogContent{Key: proto.String(k), Value: proto.String(convert.MustString(v))})
			case reflect.Float64:
				contents = append(contents, &sls.LogContent{Key: proto.String(k), Value: proto.String(strconv.FormatFloat(v.(float64), 'f', -1, 64))})
			case reflect.Float32:
				contents = append(contents, &sls.LogContent{Key: proto.String(k), Value: proto.String(strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32))})
			case reflect.Map, reflect.Slice, reflect.Array:
				bytes, _ := json.Marshal(v)
				contents = append(contents, &sls.LogContent{Key: proto.String(k), Value: proto.String(string(bytes))})
			default:
				return errors.New(k + "的值类型不支持")
			}
		}
		logs = append(
			logs,
			&sls.Log{Time: proto.Uint32(uint32(time.Now().Unix())), Contents: contents},
		)
	}

	logGroup := &sls.LogGroup{
		Topic:  proto.String(""),
		Source: proto.String(""),
		Logs:   logs,
	}

	logStoreName := t.Name
	if !t.AnyStore {
		logStoreName = "analysis-" + t.Name
	}

	if err := t._client.PutLogs(
		t.ProjectName,
		logStoreName,
		logGroup,
	); err != nil {
		return err
	}
	return nil
}

//使用最新的日志生产者日志压缩率更高
func (t *LogTcpJson) Puts(params []map[string]interface{}) error {
	if params == nil {
		return errors.New("log values is nil.")
	}
	if err := t.initLog(); err != nil {
		return err
	}

	for _, item := range params {
		if _, ok := item["env"]; !ok {
			item["env"] = env.GetString("APP_MODE", "")
		}
		if _, ok := item["app"]; !ok {
			item["app"] = env.GetString("APP_NAME", "")
		}
		if _, ok := item["uuid"]; !ok {
			item["uuid"] = t._ruid
		}
		if _, ok := item["time"]; !ok {
			item["time"] = time.Now().Unix()
		}
		if _, ok := item["micro"]; !ok {
			item["micro"] = time.Now().UnixNano() / 1e6
		}
		if _, ok := item["date"]; !ok {
			item["date"] = times.FormatDatetime(time.Now())
		}
		newItem := make(map[string]string)
		for k, v := range item {
			switch reflect.ValueOf(v).Kind() {
			case reflect.Invalid:
				newItem[k] = ""
			case reflect.String, reflect.Int, reflect.Int64, reflect.Uint, reflect.Uint64:
				newItem[k] = convert.MustString(v)
			case reflect.Float64:
				newItem[k] = convert.MustString(strconv.FormatFloat(v.(float64), 'f', -1, 64))
			case reflect.Float32:
				newItem[k] = convert.MustString(strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32))
			case reflect.Map, reflect.Slice, reflect.Array:
				newItem[k] = jsonutil.ToString(v)
			default:
				return errors.New(k + "的值类型不支持")
			}
		}
		if err := t._producer.SendLog(
			t.ProjectName, t.Name, "", "", producer.GenerateLog(uint32(time.Now().Unix()), newItem),
		); err != nil {
			return err
		}
	}
	return nil
}
