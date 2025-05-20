package logger

// Level represents the severity level of a log message.
type Level byte

const (
	// LevelDebug is used for debugging purposes.
	LevelDebug Level = iota
	// LevelInfo is used for informational messages.
	LevelInfo
	// LevelWarn is used for warning messages.
	LevelWarn
	// LevelError is used to trace errors.
	LevelError
)
