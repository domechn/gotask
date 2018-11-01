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

// Task 轮询任务
type Task struct {
	id string

	executeTime time.Time

	interval time.Duration

	do func()
}

// NewTask 新建轮询任务
func NewTask(t time.Duration, do func()) Tasker {
	uid := uuid.NewV4()
	return &Task{
		id:          uid.String(),
		do:          do,
		interval:    t,
		executeTime: time.Now().Add(t),
	}
}

func (t *Task) ExecuteTime() time.Time {
	return t.executeTime
}

func (t *Task) SetInterval(time time.Duration) {
	t.interval = time
}

func (t *Task) RefreshExecuteTime() {
	t.executeTime = t.executeTime.Add(t.interval)
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) Do() func() {
	return t.do
}

// DayTask 日任务
type DayTask struct {
	id string

	executeTime time.Time

	do func()
}

// NewDayTask 新建日任务
func NewDayTask(tm string, do func()) (Tasker, error) {
	uid := uuid.NewV4()
	pt := newTimeParser(dayParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = begin.Add(time.Hour * 24)
	}
	return &DayTask{
		id:          uid.String(),
		do:          do,
		executeTime: begin,
	}, nil
}

func NewDayTasks(tms []string, do func()) ([]Tasker, error) {
	var ts []Tasker
	for _, tm := range tms {
		dt, err := NewDayTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, dt)
	}
	return ts, nil
}

func (d *DayTask) ID() string {
	return d.ID()
}

func (d *DayTask) ExecuteTime() time.Time {
	return d.executeTime
}

func (d *DayTask) RefreshExecuteTime() {
	d.executeTime = d.executeTime.Add(time.Hour * 24)
}

func (d *DayTask) Do() func() {
	return d.Do()
}

// MonthTask 月任务
type MonthTask struct {
	id string

	executeTime time.Time

	do func()
}

// NewMonthTask 初始化一个每月执行的函数
func NewMonthTask(tm string, do func()) (Tasker, error) {
	uid := uuid.NewV4()
	pt := newTimeParser(monthParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = begin.AddDate(0, 1, 0)
	}
	return &MonthTask{
		id:          uid.String(),
		do:          do,
		executeTime: begin,
	}, nil
}

func NewMonthTasks(tms []string, do func()) ([]Tasker, error) {
	var ts []Tasker
	for _, tm := range tms {
		mt, err := NewMonthTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, mt)
	}
	return ts, nil
}

func (m *MonthTask) ID() string {
	return m.id
}

func (m *MonthTask) ExecuteTime() time.Time {
	return m.executeTime
}

func (m *MonthTask) RefreshExecuteTime() {
	m.executeTime = m.executeTime.AddDate(0, 1, 0)
}

func (m *MonthTask) Do() func() {
	return m.do
}
