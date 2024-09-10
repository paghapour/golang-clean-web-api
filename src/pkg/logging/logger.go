package logging

import "github.com/paghapour/golang-clean-web-api/config"

type Logger interface{
	Init()

	Info(Category, SubCategory, msg string, extra map[ExtraKey]interface{})
	Infof(template string, args ...interface{})

	Warn(Category, SubCategory, msg string, extra map[ExtraKey]interface{})
	Warnf(template string, args ...interface{})

	Error(err error, cat Category, SubCategory, msg string, extra map[ExtraKey]interface{})
	Errorf(err error, template string, args ...interface{})

	Fatal(err error, cat Category, SubCategory, msg string, extra map[ExtraKey]interface{})
	Fatalf(err error, template string, args ...interface{})
}

func NewLogger(cfg *config.Config) Logger{
	return nil
}

// file <- filebeat -> elasticsearch -> kibana