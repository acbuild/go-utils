# Logger Library

A lightweight and flexible logging library for Go, designed to support multiple log levels, customizable output destinations, and optional date/time prefixes.

## Features

- Log messages at different levels: `Debug`, `Info`, `Warn`, `Error`.
- Customizable output (e.g., `os.Stdout`, files, or any `io.Writer`).
- Optional date/time prefix for log messages.
- Functional options for easy configuration.

## Installation

To use this library, add it to your project using `go get`:

```bash
go get github.com/acbuild/go-utils/logger
```

## Usage

### Basic Example
```go
package main

import (
    "github.com/acbuild/go-utils/logger"
)

func main() {
    lgr := logger.New(logger.LevelInfo)

    lgr.Infof("Hello %s", "world")
    lgr.Debugf("This debug message will not appear")
    lgr.Warnf("This is a warning message")
    lgr.Errorf("This is an error message")
}

```

### Advanced Example
```go
package main

import (
    "github.com/acbuild/go-utils/logger"
    "os"
)

func main() {
    // Create a logger with date/time prefix and file output
    lgr := logger.New(
        logger.LevelDebug,
        logger.WithDateTimePrefix(),
        logger.WithFileOutput("app.log"),
    )

    defer lgr.Close() // Ensure the file is closed when done

    lgr.Debugf("This is a debug message")
    lgr.Infof("This is an info message")
    lgr.Warnf("This is a warning message")
    lgr.Errorf("This is an error message")
}
```

## Functional Options
The library provides functional options to customize the logger:

- `WithOutput(io.Writer)`: Set the output destination (e.g., os.Stdout, a file, or a buffer).
- `WithDateTimePrefix()`: Prefix each log message with the current date and time.
- `WithFileOutput(filePath string)`: Write log messages to a specified file.

## Log Levels
The logger supports the following log levels:

- `LevelDebug`: Logs all messages.
- `LevelInfo`: Logs Info, Warn, and Error messages.
- `LevelWarn`: Logs Warn and Error messages.
- `LevelError`: Logs only Error messages.

## Testing

Run the tests using:
```
go test
```