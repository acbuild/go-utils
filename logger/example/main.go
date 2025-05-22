package main

import (
	"os"
	"time"

	"github.com/acbuild/go-utils/logger"
)

func main() {
	lgr := logger.New(logger.LevelInfo, logger.WithOutput(os.Stdout), logger.WithPrefixDateTime())

	lgr.Infof("Hello %s", "world")
	lgr.Debugf("I am a debug message")
	lgr.Errorf("I am an error message")
	lgr.Warnf("I am warning message")
	lgr.Infof("I am an info messag %s", time.Now())
}
