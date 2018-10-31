// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
    "fmt"
    "time"
)

type Tasks []*Task

func (s Tasks) Len() int      { return len(s) }
func (s Tasks) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Tasks) Less(i, j int) bool {
	return s[i].executeTime.Before(s[j].executeTime)
}

type taskList struct {
	// 所有任务列表
	taskers Tasks

	// 最近执行的任务的时间
	nestedTime time.Time
}

var (
	tasks *taskList
	addC  = make(chan *Task)
	stopC = make(chan string)
)

func init() {
	tasks = &taskList{
		nestedTime: time.Now(),
	}
	go doAllTask()
}

// AddToTaskList 加入任务列表
func AddToTaskList(t *Task) {
	addC <- t
}

func (tl *taskList) addToTaskList(t *Task) {
	tl.taskers = append(tl.taskers, t)

	if t.executeTime.Before(tl.nestedTime) && t.executeTime.After(time.Now()) {
		tl.nestedTime = t.executeTime
	}
}

func Stop(id string) {
	stopC <- id
}

func (tl *taskList) stop(id string) {
	for k, v := range tl.taskers {
		if v.id == id {
			tl.taskers = append(tl.taskers[:k], tl.taskers[k+1:]...)
		}
	}
}

func doAllTask() {
	var timer *time.Timer

	var now time.Time
	for {
		// sort.Sort(tasks.taskers)

		now = time.Now()

		if len(tasks.taskers) == 0 {
			timer = time.NewTimer(time.Hour * 100000)
		} else {
			sub := tasks.nestedTime.Sub(now)
			fmt.Println(sub)
			timer = time.NewTimer(-sub)
			// timer = time.NewTimer(time.Second*2)
		}

		for {
			select {
			case now = <-timer.C:
				if tasks.nestedTime.Before(now) {
					doNestedTask()
				}
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
		if v.executeTime.Before(time.Now()) {
			go v.do()
			v.executeTime = v.executeTime.Add(v.interval)
		} else {
			return
		}
	}
}
