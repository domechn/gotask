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

var dayCase = []struct {
	name  string
	param string
	want  error
}{
	{
		name:  "case1",
		param: "10.1",
		want:  ErrParseTime,
	}, {
		name:  "case2",
		param: "10:10",
		want:  ErrParseTime,
	}, {
		name:  "case3",
		param: "10:10:20",
		want:  nil,
	},
}

func TestDayParse_Parse(t *testing.T) {
	for _, v := range dayCase {
		p := newTimeParser(dailyParseType)
		_, err := p.Parse(v.param)
		if err != v.want {
			t.Errorf("name:%s appears error:%+v , want:%+v\n", v.name, err, v.want)
		}
	}
}

var monthCase = []struct {
	name  string
	param string
	want  error
}{
	{
		name:  "case1",
		param: "10.1",
		want:  ErrParseTime,
	}, {
		name:  "case2",
		param: "10:10:20",
		want:  ErrParseTime,
	}, {
		name:  "case3",
		param: "20 10:10:20",
		want:  nil,
	}, {
		name:  "case4",
		param: "31 10:10:20",
		want:  nil,
	},
}

func TestMonthParse_Parse(t *testing.T) {
	for _, v := range monthCase {
		p := newTimeParser(monthlyParseType)
		_, err := p.Parse(v.param)
		if err != v.want {
			t.Errorf("name:%s appears error:%+v , want:%+v\n", v.name, err, v.want)
		}
	}
}

func Test_yearlyParse_Parse(t *testing.T) {
	var tt time.Time
	tests := []struct {
		name    string
		args    string
		wantRes time.Time
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "case1",
			args:    "1-10 12:12:12",
			wantRes: time.Date(time.Now().Year(), 1, 10, 12, 12, 12, 0, time.Now().Location()),
		},
		{
			name:    "case2",
			args:    "02-29 12:12:12",
			wantRes: time.Date(2020, 2, 29, 12, 12, 12, 0, time.Now().Location()),
		},
		{
			name:    "case3",
			args:    "33-01 12:12:12",
			wantErr: true,
			wantRes: tt,
		},
		{
			name:    "case2",
			args:    "03-31 12:12:12",
			wantRes: time.Date(time.Now().Year(), 3, 31, 12, 12, 12, 0, time.Now().Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &yearlyParse{}
			gotRes, err := p.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("yearlyParse.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("yearlyParse.Parse() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
