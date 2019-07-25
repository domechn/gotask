// Package gotask Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"errors"
	"sort"
	"sync"
	"time"
)

var (
	// ErrTaskNotFound cannot found task by id
	ErrTaskNotFound = errors.New("task does not exist")
	// ErrTaskTypeInValid only polling task can modify interval
	ErrTaskTypeInValid = errors.New("this type does not support modifying the execution interval")
)

// Tasks 任务列表
type Tasks []*Task

func (s Tasks) Len() int      { return len(s) }
func (s Tasks) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Tasks) Less(i, j int) bool {
	return s[i].ExecuteTime().Before(s[j].ExecuteTime())
}

type taskList struct {
	// all tasks
	taskers Tasks
}

type intervalChange struct {
	task     *Task
	interval time.Duration
}

var (
	tasks   *taskList
	editC   = make(chan interface{})
	removeC = make(chan string)
	wg      = &sync.WaitGroup{}
)

func init() {
	tasks = &taskList{}
	go doAllTask()
}

// AddToTaskList add the task to the execution list
func AddToTaskList(ts ...*Task) {
	for _, t := range ts {
		if t == nil {
			continue
		}
		wg.Add(1)
		editC <- t
	}
}

func (tl *taskList) addToTaskList(t *Task) {
	tl.taskers = append(tl.taskers, t)
	wg.Done()
}

// Remove remove corresponding tasks through the id of task
func Remove(id string) {
	removeC <- id
}

func (tl *taskList) remove(id string) {
	for k, v := range tl.taskers {
		if v.id == id {
			tl.taskers = append(tl.taskers[:k], tl.taskers[k+1:]...)
		}
	}
}

// ChangeInterval changes the interval between the tasks specified by the ID,
// Apply only to polling tasks.
func ChangeInterval(id string, interval time.Duration) error {
	tsk, err := tasks.getConcurrent(id)
	if err != nil {
		return err
	}
	if tsk.taskType != polling {
		return ErrTaskTypeInValid
	}
	editC <- &intervalChange{
		task:     tsk,
		interval: interval,
	}

	return nil
}

// Pause the running task, it only returns error
// when can not found task
func Pause(id string) error {
	tsk, err := tasks.getConcurrent(id)
	if err != nil {
		return err
	}
	tsk.pause()
	return nil
}

// Resume the paused task, it only returns error
// when can not found task
func Resume(id string) error {
	tsk, err := tasks.getConcurrent(id)
	if err != nil {
		return err
	}
	tsk.resume()
	return nil
}

func changeInterval(i *intervalChange) {
	i.task.setInterval(i.interval)
}

func (tl *taskList) get(id string) *Task {
	for _, v := range tl.taskers {
		if v.id == id {
			return v
		}
	}
	return nil
}

func (tl *taskList) getConcurrent(id string) (*Task, error) {
	tsk := tasks.get(id)
	if tsk == nil {
		wg.Wait()
		tsk = tasks.get(id)
		if tsk == nil {
			return nil, ErrTaskNotFound
		}
	}
	return tsk, nil
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
			case edit := <-editC:
				now = time.Now()
				timer.Stop()
				if t, ok := edit.(*Task); ok {
					tasks.addToTaskList(t)
				} else if ic, ok := edit.(*intervalChange); ok {
					changeInterval(ic)
				}
			case id := <-removeC:
				tasks.remove(id)
			}
			break
		}
	}
}

func doNestedTask() {
	for _, v := range tasks.taskers {
		if v.ExecuteTime().Before(time.Now()) {
			if !v.isPaused() {
				go v.do()
			}
			v.refreshExecuteTime()
		} else {
			return
		}
	}
}
