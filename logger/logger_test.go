package logger_test

import logger "github.com/acbuild/go-utils/logger"

func ExampleLogger_Debugf() {
	debugLogger := logger.New(logger.LevelDebug, nil)
	debugLogger.Debugf("Hello %s", "World")

}
