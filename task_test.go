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
	tk, _ := NewDayTask("00:00:00", func() {
		p = append(p, 1)
	})
	AddToTaskList(tk)
	fmt.Println(tk.executeTime)
}

func TestMonthTask(t *testing.T) {

}
