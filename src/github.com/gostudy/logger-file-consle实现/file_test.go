package logger

import (
	"testing"
)

func TestFileLogger(t *testing.T)  {   //测试函数，单元测试
	logger := NewFileLogger(0, "C:/logs/", "test")  //初始化

	logger.Debug("user id[%d] is come from chian", 32442)
	logger.Warn("test warn log")
	logger.Error("test Error log")
	logger.Fatal("test fatal log")
	logger.Close()
}


func TestConsoleLogger(t *testing.T)  {   //测试函数，单元测试
	logger := NewConsoleLogger(0)  //初始化

	logger.Debug("user id[%d] is come from chian", 32442)
	logger.Warn("test warn console log")
	logger.Error("test Error console log")
	logger.Fatal("test fatal console log")
	logger.Close()
}