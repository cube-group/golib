package mysync

import "sync"

type Slice struct {
	locker sync.Mutex
	list   []interface{}
}

func (t *Slice) Push(v interface{}) {
	defer t.locker.Unlock()
	t.locker.Lock()
	t.list = append(t.list, v)
}

func (t *Slice) Length() int {
	return len(t.list)
}

