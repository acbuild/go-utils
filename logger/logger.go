package logger

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Logger is used to log messages at different levels.
type Logger struct {
	threshold      Level
	output         io.Writer
	prefixDateTime bool
}

// New creates a new Logger instance with the specified threshold level.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout, prefixDateTime: false}

	for _, configFunc := range opts {
		configFunc(lgr)
	}

	return lgr
}

// Debugf formats and prints a message if the log level is Debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	l.logf(LevelDebug, format, args...)
}

func (l *Logger) Infof(format string, args ...any) {
	l.logf(LevelInfo, format, args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	l.logf(LevelWarn, format, args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	l.logf(LevelError, format, args...)
}

func (l *Logger) logf(loglevel Level, format string, args ...any) {
	if l.threshold > loglevel {
		return
	}
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.prefixDateTime {
		format = time.Now().Format(time.UnixDate) + " - " + format
	}

	fmt.Printf(format+"\n", args...)
}
