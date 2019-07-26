// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

import (
	"time"
)

type parserType string

const (
	dailyParseType   = parserType("day")
	monthlyParseType = parserType("month")
	yearlyParseType  = parserType("year")
)

// Parser to pase the timestring
type Parser interface {
	// Parse 解析定时执行的时间
	Parse(string) (time.Time, error)
}

func newTimeParser(pt parserType) Parser {
	switch pt {
	case dailyParseType:
		return newdailyParse()
	case monthlyParseType:
		return newmonthlyParse()
	case yearlyParseType:
		return newyearlyParse()
	}
	return nil
}
