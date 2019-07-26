// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_refreshExecuteTime(t *testing.T) {
	type fields struct {
		executeTime     time.Time
		nextExecuteTime func(time.Time) time.Time
		taskType        int
	}
	tests := []struct {
		name     string
		fields   fields
		wantTime time.Time
	}{
		{
			name: "yaer",
			fields: fields{
				executeTime:     time.Date(2020, 2, 28, 12, 12, 12, 0, time.Now().Location()),
				nextExecuteTime: yearlyNextFunc,
				taskType:        yearly,
			},
			wantTime: time.Date(2021, 2, 28, 12, 12, 12, 0, time.Now().Location()),
		},
		{
			name: "yaer29",
			fields: fields{
				executeTime:     time.Date(2020, 2, 29, 12, 12, 12, 0, time.Now().Location()),
				nextExecuteTime: yearlyNextFunc,
				taskType:        yearly,
			},
			wantTime: time.Date(2024, 2, 29, 12, 12, 12, 0, time.Now().Location()),
		},
		{
			name: "month",
			fields: fields{
				executeTime:     time.Date(2020, 2, 29, 12, 12, 12, 0, time.Now().Location()),
				nextExecuteTime: monthlyNextFunc,
				taskType:        monthly,
			},
			wantTime: time.Date(2020, 3, 29, 12, 12, 12, 0, time.Now().Location()),
		},
		{
			name: "month31",
			fields: fields{
				executeTime:     time.Date(2020, 3, 31, 12, 12, 12, 0, time.Now().Location()),
				nextExecuteTime: monthlyNextFunc,
				taskType:        monthly,
			},
			wantTime: time.Date(2020, 5, 31, 12, 12, 12, 0, time.Now().Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			task := &Task{
				executeTime:     tt.fields.executeTime,
				nextExecuteTime: tt.fields.nextExecuteTime,
				taskType:        tt.fields.taskType,
			}
			task.refreshExecuteTime()
			if !reflect.DeepEqual(task.ExecuteTime(), tt.wantTime) {
				t.Errorf("yearlyParse.Parse() = %v, want %v", task.ExecuteTime(), tt.wantTime)
			}
		})
	}
}
