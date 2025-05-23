package logger_test

import (
	"bytes"
	"testing"

	logger "github.com/acbuild/go-utils/logger"
)

func TestDebugfNotExpected(t *testing.T) {
	var buf bytes.Buffer

	debugLogger := logger.New(logger.LevelError, logger.WithOutput(&buf))
	debugLogger.Debugf("Hello World")

	if buf.Len() > 0 {
		t.Errorf("Expected no output.\n%s", buf.String())
	}
}

func TestDebugfExpected(t *testing.T) {
	var buf bytes.Buffer

	debugLogger := logger.New(logger.LevelDebug, logger.WithOutput(&buf))
	debugLogger.Debugf("Hello World")

	if buf.Len() == 0 {
		t.Errorf("Expected output.\n%s", buf.String())
	}
}
