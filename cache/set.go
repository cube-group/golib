package cache

import (
	"encoding/json"
	"strings"
	"sync"
	"time"
)

type RamItem []interface{}

type Ram struct {
	locker  sync.Mutex
	maps    *sync.Map //数据存储器
	forever bool      //数据是否为永久保存
}

func NewRam(forever ... bool) *Ram {
	i := &Ram{maps: new(sync.Map)}
	if len(forever) > 0 {
		i.forever = forever[0]
	}
	if !i.forever {
		go i.check()
	}
	return i
}

//心跳检测
func (t *Ram) check() {
	for {
		time.Sleep(time.Second)
		now := time.Now().Unix()
		var leaved bool
		var deleted bool
		t.maps.Range(func(key, value interface{}) bool {
			if now > value.(RamItem)[1].(int64) {
				t.maps.Delete(key)
				deleted = true
			} else {
				leaved = true
			}
			return true
		})
		if deleted && !leaved {
			t.locker.Lock()
			t.maps = new(sync.Map)
			t.locker.Unlock()
			//runtime.GC()
			time.Sleep(time.Second * 3)
		}
	}
}

//删除
func (t *Ram) Delete(key string) {
	t.maps.Delete(key)
}

//设置数据
func (t *Ram) Set(key string, value interface{}, expire time.Duration) {
	t.maps.Store(key, RamItem{value, time.Now().Add(expire).Unix()})
}

//获取数据
func (t *Ram) Get(key string) interface{} {
	if v, ok := t.maps.Load(key); ok {
		return v.(RamItem)[0]
	}
	return nil
}

//获取debug缓存信息
func (t *Ram) DebugInfo() interface{} {
	arr := make([]string, 0)
	t.maps.Range(func(key, value interface{}) bool {
		bytes, _ := json.Marshal(value)
		arr = append(arr, string(bytes))
		return true
	})
	return strings.Join(arr, "\n")
}
