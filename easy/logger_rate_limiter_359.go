package easy

import (
	"interview_go/common"
)

// https://leetcode.com/problems/logger-rate-limiter/
type LoggerRateLimiter struct {
}

type Logger struct {
	logs    map[string]int
	timeout int
}

func LoggerConstructor() Logger {
	return Logger{logs: make(map[string]int), timeout: 10}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {

	if v, ok := l.logs[message]; ok && v > timestamp {
		return false
	}

	l.logs[message] = timestamp + l.timeout

	return true
}

func (l *LoggerRateLimiter) Test() {
	logger := LoggerConstructor()
	e := true
	r := logger.ShouldPrintMessage(1, "foo")
	common.PrintBool(r, e)
	e = true
	r = logger.ShouldPrintMessage(2, "bar")
	common.PrintBool(r, e)
	e = false
	r = logger.ShouldPrintMessage(3, "foo")
	common.PrintBool(r, e)
	e = false
	r = logger.ShouldPrintMessage(8, "bar")
	common.PrintBool(r, e)
	e = false
	r = logger.ShouldPrintMessage(10, "foo")
	common.PrintBool(r, e)
	e = true
	r = logger.ShouldPrintMessage(11, "foo")
	common.PrintBool(r, e)
}
