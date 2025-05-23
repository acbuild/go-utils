package logger

import (
	"fmt"
	"io"
	"os"
)

// Option defines a functional option to our logger.
type Option func(*Logger)

// WithOutput returns a configuration function that sets the output of logs.
func WithOutput(output io.Writer) Option {
	return func(l *Logger) {
		l.output = output
	}
}

// WithPrefixDateTime returns a configuration function that sets a prefix for DateTime.
func WithPrefixDateTime() Option {
	return func(l *Logger) {
		l.prefixDateTime = true
	}
}

// WithFileOutput returns a configuration function that sets the logger's output to a specified file.
func WithFileOutput(filePath string) Option {
	return func(l *Logger) {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
			return
		}
		l.output = file
	}
}
