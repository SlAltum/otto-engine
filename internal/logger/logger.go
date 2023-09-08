package logger

import (
	"fmt"
)

type Logger struct {
	prefix string
}

func NewLogger(prefix string) *Logger {
	logger := &Logger{
		prefix: prefix,
	}
	return logger
}

func (lg *Logger) Log(level Level, format string, args ...interface{}) {
	formatStr := fmt.Sprintf("[%s]%s: %s\n", GetLvStr(level), lg.prefix, format)
	fmt.Printf(formatStr, args...)
}
