// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
	// ErrParseTime time format failed
	ErrParseTime = errors.New("parse time error")
)

var defaultValue = time.Time{}

type dailyParse struct {
}

func newdailyParse() Parser {
	return &dailyParse{}
}

// Parse 接收格式 "hh:mm:ss",返回begintime
func (p *dailyParse) Parse(s string) (time.Time, error) {
	tm, err := time.Parse("15:04:05", s)
	if err != nil {
		return defaultValue, ErrParseTime
	}
	now := time.Now()
	t := time.Date(now.Year(), now.Month(), now.Day(), tm.Hour(), tm.Minute(), tm.Second(), 0, now.Location())
	return t, nil
}

type monthlyParse struct {
}

func newmonthlyParse() Parser {
	return &monthlyParse{}
}

// Parse 接收格式 dd hh:mm:ss   dd为每月几号，如果需要每月最后一天 dd=-1
func (p *monthlyParse) Parse(s string) (res time.Time, err error) {
	s2 := strings.SplitN(s, " ", 2)
	if len(s2) != 2 {
		err = ErrParseTime
		return
	}
	var dt time.Time
	dt, err = newdailyParse().Parse(s2[1])
	if err != nil {
		return
	}
	var day int
	if day, err = strconv.Atoi(s2[0]); err != nil {
		err = ErrParseTime
		return
	}
	if day < 0 || day > 31 {
		err = ErrParseTime
		return
	}
	now := time.Now()
	step := 0
	t := time.Date(now.Year(), now.Month(), day, dt.Hour(), dt.Minute(), dt.Second(), 0, now.Location())
	if t.Day() != day {
		step++
		t = t.AddDate(0, step, 0)
	}
	return t, nil
}

type yearlyParse struct {
}

func newyearlyParse() Parser {
	return &yearlyParse{}
}

// Parse 接收格式 MM-dd hh:mm:ss   dd为每月几号，如果需要每月最后一天 dd=-1
func (p *yearlyParse) Parse(s string) (res time.Time, err error) {
	ss := strings.SplitN(s, "-", 2)
	if len(ss) != 2 {
		err = ErrParseTime
		return
	}
	var mt time.Time
	mt, err = newmonthlyParse().Parse(ss[1])
	if err != nil {
		return
	}
	month, err := strconv.Atoi(ss[0])
	if err != nil {
		err = ErrParseTime
		return
	}
	if month < 1 || month > 12 {
		err = ErrParseTime
		return
	}

	if month != 2 && mt.Day() != 29 {
		return time.Date(mt.Year(), time.Month(month), mt.Day(), mt.Hour(), mt.Minute(), mt.Second(), 0, mt.Location()), nil
	}

	year := mt.Year()
	for !(year%4 == 0 && year%100 != 0) {
		year++
	}
	res = time.Date(year, time.Month(month), mt.Day(), mt.Hour(), mt.Minute(), mt.Second(), 0, mt.Location())
	return
}
