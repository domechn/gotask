// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"fmt"
	"sync"
	"time"
)

// Task Polling tasks
type Task struct {
	sync.RWMutex

	id string

	executeTime time.Time

	interval time.Duration

	do func()
}

// NewTask create a new polling task
func NewTask(t time.Duration, do func()) (Tasker, error) {
	if t < time.Millisecond {
		return nil, fmt.Errorf("the execution interval is too short")
	}
	idStr := RandStringBytesMaskImprSrc(20)
	return &Task{
		id:          idStr,
		do:          do,
		interval:    t,
		executeTime: time.Now().Add(t),
	}, nil
}

// ExecuteTime gets the next execution time
func (t *Task) ExecuteTime() time.Time {
	t.RLock()
	defer t.RUnlock()
	return t.executeTime
}

// SetInterval modify execution interval
func (t *Task) SetInterval(td time.Duration) {
	t.Lock()
	defer t.Unlock()
	t.interval = td
	t.changeExecuteTime(td)
}

func (t *Task) changeExecuteTime(td time.Duration) {
	t.executeTime = time.Now().Add(td)
}

// RefreshExecuteTime refresh execution interval
func (t *Task) RefreshExecuteTime() {
	t.Lock()
	t.Unlock()
	t.executeTime = t.executeTime.Add(t.interval)
}

// ID return taskID
func (t *Task) ID() string {
	t.RLock()
	defer t.RUnlock()
	return t.id
}

// Do return Task Function
func (t *Task) Do() func() {
	t.RLock()
	defer t.RUnlock()
	return t.do
}

// DailyTask run every day
type DailyTask struct {
	id string

	executeTime time.Time

	do func()
}

// NewDailyTask create a new daily task
func NewDailyTask(tm string, do func()) (Tasker, error) {
	idStr := RandStringBytesMaskImprSrc(20)
	pt := newTimeParser(dayParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = begin.Add(time.Hour * 24)
	}
	return &DailyTask{
		id:          idStr,
		do:          do,
		executeTime: begin,
	}, nil
}

// NewDailyTasks create new daily tasks
func NewDailyTasks(tms []string, do func()) ([]Tasker, error) {
	var ts []Tasker
	for _, tm := range tms {
		dt, err := NewDailyTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, dt)
	}
	return ts, nil
}

// ID returns task id
func (d *DailyTask) ID() string {
	return d.id
}

// ExecuteTime returns executeTime
func (d *DailyTask) ExecuteTime() time.Time {
	return d.executeTime
}

// RefreshExecuteTime change excuteTime
func (d *DailyTask) RefreshExecuteTime() {
	d.executeTime = d.executeTime.Add(time.Hour * 24)
}

// Do daily task
func (d *DailyTask) Do() func() {
	return d.do
}

// MonthTask create monthly task
type MonthTask struct {
	id string

	executeTime time.Time

	do func()
}

// NewMonthTask initialize a function that executes each month
func NewMonthTask(tm string, do func()) (Tasker, error) {
	idStr := RandStringBytesMaskImprSrc(20)
	pt := newTimeParser(monthParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = begin.AddDate(0, 1, 0)
	}
	return &MonthTask{
		id:          idStr,
		do:          do,
		executeTime: begin,
	}, nil
}

// NewMonthTasks initialize a function that executes each month
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

// ID return task id
func (m *MonthTask) ID() string {
	return m.id
}

// ExecuteTime return excuteTime
func (m *MonthTask) ExecuteTime() time.Time {
	return m.executeTime
}

// RefreshExecuteTime change executeTime
func (m *MonthTask) RefreshExecuteTime() {
	m.executeTime = m.executeTime.AddDate(0, 1, 0)
}

// Do monthly task
func (m *MonthTask) Do() func() {
	return m.do
}
