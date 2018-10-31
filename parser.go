// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package gotask

type parserType string

const (
	dayParseType   = parserType("day")
	monthParseType = parserType("month")
)

type Parser interface {
	// Parse 解析定时执行的时间
	Parse(string) (string, error)
}

func newTimeParser(pt parserType) Parser {
	switch pt {
	case dayParseType:
		return newDayParse()
	case monthParseType:
		return newMonthParse()
	}
	return nil
}
