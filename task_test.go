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
	tk := NewTask(time.Second, func() {
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

func TestDayTask(t *testing.T) {
	var p []int
	tks, _ := NewDayTasks([]string{"00:00:00", "12:00:00", "08:00:00"}, func() {
		p = append(p, 1)
	})
	AddToTaskList(tks...)
	for _, tk := range tks {
		fmt.Println(tk.ExecuteTime())
	}
}

func TestMonthTask(t *testing.T) {
	var tks []Tasker
	tks, _ = NewMonthTasks([]string{"1 00:00:00", "1 1:00:00", "3 08:00:00"}, func() {

	})
	for _, tk := range tks {
		fmt.Println(tk.ExecuteTime())
	}
	time.Sleep(time.Second)
}
