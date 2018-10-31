// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
    "errors"
    "strings"
)

var (
    errParseTime = errors.New("parse time error")
)

type dayParse struct {
}

func newDayParse() Parser {
    return &dayParse{}
}

// Parse 接收格式 "hh:mm:ss"
func (p *dayParse) Parse(s string) (string, error) {
    ss := strings.Split(s,":")
    if len(ss) != 3 {
        return "",errParseTime
    }
    return "",nil
}

type monthParse struct {
}

func newMonthParse() Parser {
    return &monthParse{}
}

// Parse 接收格式 dd mm:hh:ss   dd为每月几号，如果需要每月最后一天 dd=-1
func (p *monthParse) Parse(s string) (string, error) {
    ss := strings.SplitN(s," ",2)
    s1 := strings.Split(ss[1],":")
    if len(s1) != 3 {
        return "",errParseTime
    }
    return "",nil
}



