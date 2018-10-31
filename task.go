// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
    "github.com/satori/go.uuid"
    "time"
)

type Task struct {
	id string

	executeTime time.Time

	interval time.Duration

	do func()

	stopC     chan struct{}
	isStarted int32
}

func NewTask(t time.Duration, do func()) *Task {
	uid, _ := uuid.NewV4()
	return &Task{
		id:       uid.String(),
		do:       do,
		interval: t,
		executeTime:time.Now(),
	}
}
