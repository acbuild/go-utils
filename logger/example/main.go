package main

import (
	"os"

	"github.com/acbuild/go-utils/logger"
)

func main() {
	lgr := logger.New(logger.LevelInfo, logger.WithOutput(os.Stdout), logger.WithPrefixDateTime())

	// Example with logging to file
	//lgr := logger.New(logger.LevelInfo, logger.WithPrefixDateTime(), logger.WithFileOutput("app.log"))
	//defer lgr.Close()

	lgr.Infof("Hello %s", "world")
	lgr.Debugf("I am a debug message")
	lgr.Errorf("I am an error message")
	lgr.Warnf("I am warning message")
}
