package context

import (
	"context"
	"github.com/cube-group/golib/uuid"
	"time"
)

type RoutineContext struct {
	Id      string
	Cancel  context.Context
	Timeout context.Context

	cancelFunc context.CancelFunc
}

//创建一个协程关联实例
//timeout 超时秒数
func NewRoutineContext(timeout int64) *RoutineContext {
	cancelCtx, cancelFunc := context.WithCancel(context.Background())
	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	return &RoutineContext{
		uuid.GetUUID(),
		cancelCtx,
		timeoutCtx,
		cancelFunc,
	}
}

//是否被删除
func (this *RoutineContext) IsCancel() <-chan struct{} {
	return this.Cancel.Done()
}

//是否超时
func (this *RoutineContext) IsTimeout() <-chan struct{} {
	return this.Timeout.Done()
}

//强制执行关闭
func (this *RoutineContext) Close() {
	this.cancelFunc()
}
