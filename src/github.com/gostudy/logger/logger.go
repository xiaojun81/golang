package logger

import (
	"fmt"
)

var log LogInterface


//file,初始化一个文件日志实例
//consoe，初始化一个console日志实例
func InitLogger(name string, config map[string]string ) (err error){
	switch name {
	case "file":
		log, err = NewFileLogger(config)
	case "console":
		log, err = NewConsoleLog(config)
	default:
		err = fmt.Errorf("unsupport logger name:%s", name)
	}
	return
}

func Debug(format string, args ...interface{})  {
	log.Debug(format,args...)
}

func Info(format string, args ...interface{})  {
	log.Info(format,args...)
}

func Trace(format string, args ...interface{}) {
	log.Trace(format, args...)
}


func Warn(format string, args ...interface{})  {
	log.Warn(format,args...)
}

func Error(format string, args ...interface{}) {
	log.Error(format, args...)
}
func Fatal(format string, args ...interface{})  {
	log.Fatal(format, args...)
}