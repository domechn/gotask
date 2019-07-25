// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"fmt"
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	var p []int
	tk, _ := NewTask(time.Second, func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	select {
	case <-time.After(time.Second*1 + time.Millisecond*100):
		if len(p) != 1 {
			t.Errorf("TestTask() fail , need len : %d , actually len : %d", 1, len(p))
		}
		return
	}
}

func TestDailyTask(t *testing.T) {
	tks, _ := NewDailyTasks([]string{"00:00:00", "12:00:00", "08:00:00"}, func() {

	})
	AddToTaskList(tks...)
	for _, tk := range tks {
		fmt.Println(tk.ExecuteTime())
		fmt.Println(tk.ID())
		tk.refreshExecuteTime()
		if tk.ExecuteTime().Day()-time.Now().Day() < 1 {
			t.Errorf("RefreshExecuteTime() appears error")
		}
		tk.do()
	}
}

func init() {
	tasks = &taskList{}
	go doAllTask()
}

func TestChangeInterval(t *testing.T) {
	var p []int
	tk, _ := NewTask(time.Second*100, func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	ChangeInterval(tk.ID(), time.Millisecond*200)
	tc := time.After(time.Second)
	select {
	case <-tc:
		if len(p) == 0 {
			t.Errorf("TestTask() fail , need len : 1, actually len : %d", len(p))
		}
		return
	}
}

func TestChangeInterval2(t *testing.T) {
	var p []int
	tk, _ := NewDailyTask("20:20:20", func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	err := ChangeInterval(tk.ID(), time.Millisecond*200)
	if err != ErrTaskTypeInValid {
		t.Errorf("ChangeInterval2 expect get err: %v, actually: %v", ErrTaskTypeInValid, err)
	}
}

func TestMonthlyTask(t *testing.T) {
	var tks []*Task
	tks, _ = NewMonthlyTasks([]string{"1 00:00:00", "1 1:00:00", "3 08:00:00"}, func() {
		fmt.Println(time.Now())
	})
	for _, tk := range tks {
		tk.refreshExecuteTime()
		a := int(time.Now().Month()) - int(tk.ExecuteTime().Month())
		if a > -2 {
			if a > 10 {
				t.Errorf("RefreshExecuteTime() appears error")
			}
		}
		tk.do()
	}
	time.Sleep(time.Second)
}

func TestPauseAndResume(t *testing.T) {
	var p []int
	tk, _ := NewTask(time.Millisecond*90, func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	Pause(tk.ID())
	time.Sleep(time.Millisecond * 300)
	if len(p) != 0 {
		t.Errorf("TestTask() fail , need len : 0, actually len : %d", len(p))
	}
	Resume(tk.ID())
	time.Sleep(time.Millisecond * 200)
	if len(p) == 0 {
		t.Errorf("TestTask() fail , need len : !=0, actually len : %d", len(p))
	}
}

func TestRemove(t *testing.T) {
	var p []int
	tk, _ := NewTask(time.Millisecond*90, func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	time.Sleep(time.Millisecond * 100)
	Remove(tk.ID())
	time.Sleep(time.Millisecond * 100)
	if tasks.get(tk.ID()) != nil {
		t.Errorf("Remove() failed")
	}
}
