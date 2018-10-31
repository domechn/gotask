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

func TestAddToTaskList(t *testing.T) {
	for i := 0; i < 100; i++ {
		tk := NewTask(time.Second*5, func() {
			fmt.Println(time.Now().String())
		})
		go AddToTaskList(tk)
	}
	time.Sleep(time.Second * 20)
}

func TestAppendToTaskList(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func() {
			ticker := time.NewTimer(time.Second * 2)
			for {
				select {
				case <-ticker.C:
					fmt.Println(time.Now().String())
				}
			}
		}()
	}
	time.Sleep(time.Second * 20)

}
