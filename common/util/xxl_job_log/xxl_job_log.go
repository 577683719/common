package xxl_job_log

import (
	"fmt"
	"log"
)

var Log = &Logger{}

// xxl.Logger接口实现
type Logger struct{}

func (l *Logger) Info(format string, a ...interface{}) {
	fmt.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}

func (l *Logger) Error(format string, a ...interface{}) {
	log.Println(fmt.Sprintf("自定义日志 - "+format, a...))
}
