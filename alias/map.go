package alias

import (
	"runtime"
	"sync"
	"time"
)

type dict struct {
	_map             sync.Map
	_expire          sync.Map
	cleaningDuration time.Duration // 清除失效key的间隔时间
	cleaningTime     int64         // 上次完成清除的时间
	lastWriteTime    int64         // 最后一条记录过期时间
	len              int           // key的总量(包括失效未清除的)
}

func newDict() *dict {
	em := &dict{}
	go func() {
		em.cleaningTime = time.Now().UnixNano()
		for {
			now := time.Now().UnixNano()
			n := now - em.cleaningTime
			time.Sleep(em.cleaningDuration - time.Duration(n))
			now = time.Now().UnixNano()
			if now > em.lastWriteTime {
				em._map = sync.Map{}
				em._expire = sync.Map{}
				em.len = 0
				runtime.GC()
			}
			em.cleaningTime = time.Now().UnixNano()
		}
	}()
	return em
}

func (e *dict) SetCleaningDuration(d time.Duration) {
	e.cleaningDuration = d
}

func (e *dict) set(key, value interface{}, expired time.Duration) error {
	e._map.Store(key, value)
	t := time.Now().Add(expired).UnixNano()
	e._expire.Store(key, &t)
	e.len += 1
	e.lastWriteTime = t
	return nil
}

func (e *dict) del(key interface{}) {
	e._map.Delete(key)
	e._expire.Delete(key)
	e.len -= 1
}

func (e *dict) get(key interface{}) (interface{}, bool) {
	v, ok := e._expire.Load(key)
	if ok && time.Now().UnixNano() > *v.(*int64) {
		e.del(key)
		return nil, false
	}
	return e._map.Load(key)
}

func (e *dict) All() sync.Map {
	return e._map
}

type ExpiredMap struct {
	m []*dict
}

func NewExpireMap() *ExpiredMap {
	em := &ExpiredMap{
		m: []*dict{newDict(), newDict()},
	}
	return em
}

func (e *ExpiredMap) SetCleaningDuration(d time.Duration) {
	e.m[0].cleaningDuration = d
	e.m[1].cleaningDuration = d
}

func (e *ExpiredMap) Set(key, value interface{}, expired time.Duration) error {
	if _, ok := e.m[0]._map.Load(key); ok {
		e.m[0].len -= 1
		return e.m[0].set(key, value, expired)
	}
	if _, ok := e.m[1]._map.Load(key); ok {
		e.m[1].len -= 1
		return e.m[1].set(key, value, expired)
	}
	if e.m[0].len <= 0 || e.m[1].len > 0 {
		return e.m[0].set(key, value, expired)
	} else {
		e.m = []*dict{e.m[1], e.m[0]}
		return e.m[0].set(key, value, expired)
	}
}

func (e *ExpiredMap) Del(key interface{}) {
	e.m[0].del(key)
	e.m[1].del(key)
}

func (e *ExpiredMap) Get(key interface{}) (interface{}, bool) {
	res, ok := e.m[0].get(key)
	if ok {
		return res, ok
	}
	return e.m[1].get(key)
}

func (e *ExpiredMap) All() sync.Map {
	return e.m[0].All()
}
