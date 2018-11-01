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

var dayCase = []struct {
	name  string
	param string
	want  error
}{
	{
		name:  "case1",
		param: "10.1",
		want:  errParseTime,
	}, {
		name:  "case2",
		param: "10:10",
		want:  errParseTime,
	}, {
		name:  "case3",
		param: "10:10:20",
		want:  nil,
	},
}

func TestDayParse_Parse(t *testing.T) {
	for _, v := range dayCase {
		p := newTimeParser(dayParseType)
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
		want:  errParseTime,
	}, {
		name:  "case2",
		param: "10:10:20",
		want:  errParseTime,
	}, {
		name:  "case3",
		param: "20 10:10:20",
		want:  nil,
	},
}

func TestMonthParse_Parse(t *testing.T) {
	for _, v := range monthCase {
		p := newTimeParser(monthParseType)
		_, err := p.Parse(v.param)
		if err != v.want {
			t.Errorf("name:%s appears error:%+v , want:%+v\n", v.name, err, v.want)
		}
	}
	ts := time.Date(2018, 2, 28, 0, 0, 0, 0, time.Now().Location())
	ts = ts.AddDate(0, 1, 0)
	fmt.Println(ts)

}
