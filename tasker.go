// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"time"
)

// Tasker 任务接口类
type Tasker interface {

	// ID 返回执行函数的ID
	ID() string

	// ExecuteTime 获取下一次执行时间
	ExecuteTime() time.Time

	// RefreshExecuteTime 刷新执行时间
	RefreshExecuteTime()

	// Do 返回执行函数
	Do() func()
}
