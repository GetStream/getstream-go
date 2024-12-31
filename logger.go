package getstream

import (
	"io"
	"log"
	"os"
	"sync"
)

// LogLevel represents the severity of a log message.
type LogLevel int

const (
	// LogLevelDebug is the lowest severity.
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is for general information.
	LogLevelInfo
	// LogLevelWarn is for warning messages.
	LogLevelWarn
	// LogLevelError is for error messages.
	LogLevelError
)

// Logger is an interface that clients can implement to provide custom logging.
type Logger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warn(format string, v ...interface{})
	Error(format string, v ...interface{})
}

// DefaultLogger is the default implementation of the Logger interface.
type DefaultLogger struct {
	logger *log.Logger
	level  LogLevel
	mu     sync.Mutex
}

// NewDefaultLogger creates a new DefaultLogger instance.
func NewDefaultLogger(out io.Writer, prefix string, flag int, level LogLevel) *DefaultLogger {
	return &DefaultLogger{
		logger: log.New(out, prefix, flag),
		level:  level,
	}
}

// SetLevel sets the logging level.
func (l *DefaultLogger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Debug logs a debug message.
func (l *DefaultLogger) Debug(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelDebug {
		l.logger.Printf("[DEBUG] "+format, v...)
	}
}

// Info logs an info message.
func (l *DefaultLogger) Info(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelInfo {
		l.logger.Printf("[INFO] "+format, v...)
	}
}

// Warn logs a warning message.
func (l *DefaultLogger) Warn(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelWarn {
		l.logger.Printf("[WARN] "+format, v...)
	}
}

// Error logs an error message.
func (l *DefaultLogger) Error(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelError {
		l.logger.Printf("[ERROR] "+format, v...)
	}
}

// DefaultLoggerInstance is the default logger instance.
var DefaultLoggerInstance Logger = NewDefaultLogger(os.Stderr, "", log.LstdFlags, LogLevelInfo)
