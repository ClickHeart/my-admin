package bootstrap

import (
	"strings"

	"github.com/donnie4w/go-logger/logger"
)

func LogInit(level string) {
	attrformat := &logger.AttrFormat{
		SetBodyFmt: func(l logger.LEVELTYPE, bs []byte) []byte {
			//处理日志末尾换行符
			if size := len(bs); bs[size-1] == '\n' {
				bs = append(bs[:size-1], []byte("\x1b[0m\n")...)
			} else {
				bs = append(bs, []byte("\x1b[0m\n")...)
			}
			switch l {
			case logger.LEVEL_DEBUG:
				return append([]byte("\x1b[34m"), bs...)
			case logger.LEVEL_INFO:
				return append([]byte("\x1b[32m"), bs...)
			case logger.LEVEL_WARN:
				return append([]byte("\x1b[33m"), bs...)
			case logger.LEVEL_ERROR:
				return append([]byte("\x1b[31m"), bs...)
			case logger.LEVEL_FATAL:
				return append([]byte("\x1b[41m"), bs...)
			default:
				return bs
			}
		},
	}
	logger.SetOption(&logger.Option{
		Level:      getLevel(level),
		AttrFormat: attrformat,
		Console:    true,
		Stacktrace: logger.LEVEL_WARN,
		FileOption: &logger.FileTimeMode{
			Filename:   "log/naive_admin.log",
			Maxbuckup:  3,
			IsCompress: false,
			Timemode:   logger.MODE_MONTH,
		},
	})
}

func getLevel(level string) (l logger.LEVELTYPE) {
	level = strings.ToLower(level)
	if level == "debug" {
		l = logger.LEVEL_DEBUG
	} else if level == "info" {
		l = logger.LEVEL_INFO
	} else if level == "warn" {
		l = logger.LEVEL_WARN
	} else if level == "error" {
		l = logger.LEVEL_ERROR
	} else if level == "fatal" {
		l = logger.LEVEL_FATAL
	} else if level == "off" {
		l = logger.LEVEL_OFF
	} else {
		l = logger.LEVEL_ALL
	}
	return
}
