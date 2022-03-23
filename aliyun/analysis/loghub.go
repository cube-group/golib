package analysis

import (
	"github.com/aliyun/aliyun-log-go-sdk"
	"github.com/pkg/errors"
	"github.com/cube-group/golib/types/convert"
	"time"
)

const (
	PROJECT_NAME = "huabei2-slb"
)

//cloud client
type LogReader struct {
	AliyunEndpoint        string
	AliyunAccessKeyId     string
	AliyunAccessKeySecret string

	//分析名称
	Name string
	//项目群
	ProjectName string
	//是否为任意logstore 默认为必须为analysis-xx
	AnyStore bool
	//消费并发能力，默认10，为0时自动停止消费函数
	ConsumeCount int

	_client      *sls.Client
}

//获取调用阿里云日志的client
func (t *LogReader)aliClient() (*sls.Client, error) {
	if t.AliyunEndpoint == "" {
		t.AliyunEndpoint = "cn-beijing.log.aliyuncs.com"
	}
	c := &sls.Client{
		Endpoint:    t.AliyunEndpoint,
		AccessKeyID: t.AliyunAccessKeyId, AccessKeySecret: t.AliyunAccessKeySecret,
	}
	if c.AccessKeyID == "" || c.AccessKeySecret == "" {
		return nil, errors.New("aliyun log access key config is nil")
	}
	return c, nil
}

func (t *LogReader) Close() error {
	if t._client != nil {
		return t._client.Close()
	}
	return nil
}

func (t *LogReader) logStoreName() (string, error) {
	if t.Name == "" {
		return "", errors.New("name is nil.")
	}
	if t.ProjectName == "" {
		t.ProjectName = PROJECT_NAME
	}
	if t._client == nil {
		c, err := t.aliClient()
		if err != nil {
			return "", err
		}
		t._client = c
	}
	if t.ConsumeCount == 0 {
		t.ConsumeCount = 10
	} else if t.ConsumeCount > 100 {
		t.ConsumeCount = 100
	}

	logStoreName := t.Name
	if !t.AnyStore {
		logStoreName = "analysis-" + t.Name
	}
	return logStoreName, nil
}

//设置索引
func (t *LogReader) SetKeyIndex(key string, indexType string) error {
	if indexType != "long" && indexType != "text" && indexType != "json" {
		return errors.New("indexType is err.")
	}
	logStoreName, err := t.logStoreName()
	if err != nil {
		return err
	}

	indexKeys := map[string]sls.IndexKey{}
	if indexType == "long" {
		indexKeys[key] = sls.IndexKey{
			Token:         []string{" "},
			CaseSensitive: false,
			Type:          indexType,
		}
	} else if indexType == "text" {
		indexKeys[key] = sls.IndexKey{
			Token:         []string{",", ":", " "},
			CaseSensitive: false,
			Type:          "text",
		}
	} else {
		return errors.New("indexType not support")
	}

	index := sls.Index{
		Keys: indexKeys,
		Line: &sls.IndexLine{
			Token:         []string{",", ":", " "},
			CaseSensitive: false,
			IncludeKeys:   []string{},
			ExcludeKeys:   []string{},
		},
	}
	err = t._client.CreateIndex(t.ProjectName, logStoreName, index)
	if err != nil {
		return err
	}
	return nil
}

//获取日志
func (t *LogReader) GetLogs(query string, start, end, maxLineNum, offset int64) ([]map[string]string, error) {
	logStoreName, err := t.logStoreName()
	if err != nil {
		return nil, err
	}

	resp, err := t._client.GetLogs(
		t.ProjectName,
		logStoreName,
		"",
		start, end, query, maxLineNum, offset, true,
	)
	if err != nil {
		return nil, err
	}
	return resp.Logs, nil
}

//标准日志消费回调函数
//logs 批量日志
//shardId 本批次日志所属分片
//nextCursor 本批次日志所属分片下一个即将消费的日志游标
//err 日志消费是否失败
type PullLogsHandler func(logs []map[string]string, shardId int, nextCursor string, err error)

//消费日志
//start 开始时间戳，如果是从日志库最开始则写0
//handler 日志消费回调函数
//shardCursors 不同shard的cursor记录,如果是从头消费则可为nil
func (t *LogReader) ConsumeLogs(start int64, handler PullLogsHandler, shardCursors map[int]string) {
	logStoreName, err := t.logStoreName()
	if err != nil {
		handler(nil, 0, "", err)
		return
	}

CONTINUE:
	var shards []*sls.Shard
	if t.ConsumeCount == 0 { //终止检测
		goto END
	}

	shards, err = t._client.ListShards(t.ProjectName, logStoreName)
	if err != nil {
		handler(nil, 0, "", err)
		return
	}

	for _, shard := range shards {
		if t.ConsumeCount == 0 { //终止检测
			goto END
		}

		beginCursor := ""
		if shardCursors != nil {
			if cursor, ok := shardCursors[shard.ShardID]; ok {
				beginCursor = cursor
			}
		}
		if beginCursor == "" {
			from := "begin"
			if start > 0 {
				beginCursor = convert.MustString(start)
			}
			c, err := t._client.GetCursor(t.ProjectName, logStoreName, shard.ShardID, from)
			if err != nil {
				continue
			}
			beginCursor = c
		}

		if t.ConsumeCount == 0 { //终止检测
			goto END
		}
		endCursor, err := t._client.GetCursor(t.ProjectName, logStoreName, shard.ShardID, "end")
		if err != nil {
			continue
		}

		nextCursor := beginCursor
		for nextCursor != endCursor {
			gl, nc, err := t._client.PullLogs(t.ProjectName, logStoreName, shard.ShardID, nextCursor, endCursor, t.ConsumeCount)
			if err != nil {
				break
			}
			nextCursor = nc
			logs := make([]map[string]string, 0)
			if gl != nil {
				for _, lg := range gl.LogGroups {
					//for _, tag := range lg.LogTags {
					//	handler(tag)
					//}
					for _, l := range lg.Logs {
						maps := map[string]string{}
						for _, c := range l.Contents {
							maps[c.GetKey()] = c.GetValue()
						}
						logs = append(logs, maps)
					}
				}
			}
			handler(logs, shard.ShardID, nextCursor, nil)
		}

		if shardCursors == nil {
			shardCursors = map[int]string{}
		}
		shardCursors[shard.ShardID] = nextCursor
	}

	time.Sleep(time.Second)
	goto CONTINUE

END:
	return
}
