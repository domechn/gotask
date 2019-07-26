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

const (
	idLen       = 20
	dayInterval = time.Hour * 24
)

const (
	polling = iota
	daily
	monthly
	yearly
)

var (
	monthlyNextFunc = func(tm time.Time) time.Time {
		step := 1
	PASS:
		newTime := tm.AddDate(0, step, 0)
		if newTime.Day() != tm.Day() {
			step++
			// some months may not include this day
			goto PASS
		}
		return newTime
	}
	yearlyNextFunc = func(tm time.Time) time.Time {
		step := 1
	PASS:
		newTime := tm.AddDate(step, 0, 0)
		if newTime.Month() != tm.Month() {
			step++
			goto PASS
		}
		return newTime
	}
)

// Task Polling tasks
type Task struct {
	sync.RWMutex

	id              string
	executeTime     time.Time
	nextExecuteTime func(time.Time) time.Time
	paused          bool
	taskType        int

	do func()
}

// NewTask create a new polling task
func NewTask(t time.Duration, do func()) (*Task, error) {
	if t < time.Millisecond {
		return nil, fmt.Errorf("the execution interval is too short")
	}
	idStr := RandStringBytesMaskImprSrc(idLen)
	return &Task{
		id: idStr,
		do: do,
		nextExecuteTime: func(tm time.Time) time.Time {
			return tm.Add(t)
		},
		executeTime: time.Now().Add(t),
		taskType:    polling,
	}, nil
}

// NewDailyTask create a new daily task
func NewDailyTask(tm string, do func()) (*Task, error) {
	idStr := RandStringBytesMaskImprSrc(idLen)
	pt := newTimeParser(dailyParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = begin.Add(dayInterval)
	}
	return &Task{
		id:          idStr,
		do:          do,
		executeTime: begin,
		nextExecuteTime: func(tm time.Time) time.Time {
			return tm.Add(dayInterval)
		},
		taskType: daily,
	}, nil
}

// NewDailyTasks create new daily tasks
func NewDailyTasks(tms []string, do func()) ([]*Task, error) {
	var ts []*Task
	for _, tm := range tms {
		dt, err := NewDailyTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, dt)
	}
	return ts, nil
}

// NewMonthlyTask initialize a function that executes each month
func NewMonthlyTask(tm string, do func()) (*Task, error) {
	idStr := RandStringBytesMaskImprSrc(idLen)
	pt := newTimeParser(monthlyParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = monthlyNextFunc(begin)
	}
	return &Task{
		id:              idStr,
		do:              do,
		executeTime:     begin,
		nextExecuteTime: monthlyNextFunc,
		taskType:        monthly,
	}, nil
}

// NewMonthlyTasks initialize a function that executes each month
func NewMonthlyTasks(tms []string, do func()) ([]*Task, error) {
	var ts []*Task
	for _, tm := range tms {
		mt, err := NewMonthlyTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, mt)
	}
	return ts, nil
}

// NewYearlyTask initialize a function that executes each month
func NewYearlyTask(tm string, do func()) (*Task, error) {
	idStr := RandStringBytesMaskImprSrc(idLen)
	pt := newTimeParser(yearlyParseType)
	begin, err := pt.Parse(tm)
	if err != nil {
		return nil, err
	}
	if begin.Before(time.Now()) {
		begin = yearlyNextFunc(begin)
	}
	return &Task{
		id:              idStr,
		do:              do,
		executeTime:     begin,
		nextExecuteTime: yearlyNextFunc,
		taskType:        yearly,
	}, nil
}

// NewYearlyTasks initialize a function that executes each month
func NewYearlyTasks(tms []string, do func()) ([]*Task, error) {
	var ts []*Task
	for _, tm := range tms {
		mt, err := NewYearlyTask(tm, do)
		if err != nil {
			return nil, err
		}
		ts = append(ts, mt)
	}
	return ts, nil
}

// ExecuteTime gets the next execution time
func (t *Task) ExecuteTime() time.Time {
	t.RLock()
	defer t.RUnlock()
	return t.executeTime
}

// SetInterval modify execution interval only for polling task
func (t *Task) setInterval(td time.Duration) {
	t.Lock()
	defer t.Unlock()
	t.nextExecuteTime = func(tm time.Time) time.Time {
		return tm.Add(td)
	}
	t.executeTime = time.Now().Add(td)
}

// refreshExecuteTime refresh execution interval
func (t *Task) refreshExecuteTime() {
	t.Lock()
	defer t.Unlock()
	t.executeTime = t.nextExecuteTime(t.executeTime)
}

// ID return taskID
func (t *Task) ID() string {
	return t.id
}

// pause the runnning task
func (t *Task) pause() {
	t.Lock()
	defer t.Unlock()
	t.paused = true
}

func (t *Task) isPaused() bool {
	t.RLock()
	defer t.RUnlock()
	return t.paused
}

// resume the paused task
func (t *Task) resume() {
	t.Lock()
	defer t.Unlock()
	t.paused = false
}
