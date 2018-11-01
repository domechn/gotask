// Package gotask Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"fmt"
	"sort"
	"time"
)

// Tasks 任务列表
type Tasks []Tasker

func (s Tasks) Len() int      { return len(s) }
func (s Tasks) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Tasks) Less(i, j int) bool {
	return s[i].ExecuteTime().Before(s[j].ExecuteTime())
}

type taskList struct {
	// 所有任务列表
	taskers Tasks
}

var (
	tasks *taskList
	addC  = make(chan Tasker)
	stopC = make(chan string)
)

func init() {
	tasks = &taskList{}
	go doAllTask()
}

// AddToTaskList 加入任务列表
func AddToTaskList(ts ...Tasker) {
	for _, t := range ts {
		if t == nil {
			return
		}
		addC <- t
	}
}

func (tl *taskList) addToTaskList(t Tasker) {
	tl.taskers = append(tl.taskers, t)
}

// Stop 通过task的id停止对应task
func Stop(id string) {
	stopC <- id
}

func (tl *taskList) stop(id string) {
	for k, v := range tl.taskers {
		if v.ID() == id {
			tl.taskers = append(tl.taskers[:k], tl.taskers[k+1:]...)
		}
	}
}

func ChangeInterval(id string, interval time.Duration) error {
	tsk := tasks.get(id)
	if tsk != nil {
		t, ok := tsk.(*Task)
		if !ok {
			return fmt.Errorf("该类型不支持修改执行间隔")
		}
		t.SetInterval(interval)
	}
	return nil
}

func (tl *taskList) get(id string) Tasker {
	for _, v := range tl.taskers {
		if v.ID() == id {
			return v
		}
	}
	return nil
}

func doAllTask() {
	var timer *time.Timer

	var now time.Time
	for {
		sort.Sort(tasks.taskers)

		now = time.Now()

		if len(tasks.taskers) == 0 {
			timer = time.NewTimer(time.Hour * 100000)
		} else {
			sub := tasks.taskers[0].ExecuteTime().Sub(now)
			if sub < 0 {
				sub = 0
			}
			timer = time.NewTimer(sub)
		}

		for {
			select {
			case now = <-timer.C:
				doNestedTask()
			case t := <-addC:
				now = time.Now()
				timer.Stop()
				tasks.addToTaskList(t)
			case id := <-stopC:
				tasks.stop(id)
			}
			break
		}
	}
}

func doNestedTask() {
	for _, v := range tasks.taskers {
		if v.ExecuteTime().Before(time.Now()) {
			go v.Do()()
			v.RefreshExecuteTime()
		} else {
			return
		}
	}
}
