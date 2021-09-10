// Package assertion 断言类，如果传入值存在，则触发对应的 log
package assertion

import "github.com/sirupsen/logrus"

// Panic 恐慌
func Panic(err interface{}) {
	judge(err, logrus.Panic)
}

// Error 错误
func Error(err interface{}) {
	judge(err, logrus.Error)
}

// Warn 警告
func Warn(err interface{}) {
	judge(err, logrus.Warn)
}

func judge(err interface{}, cb func(args ...interface{})) {
	if err != nil {
		cb(err)
	}
}
