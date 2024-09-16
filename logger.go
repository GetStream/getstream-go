package getstream

import (
	"io"
	"log"
	"os"
	"sync"
)

// LogLevel represents the severity of a log message
type LogLevel int

const (
	// LogLevelDebug is the lowest severity
	LogLevelDebug LogLevel = iota
	// LogLevelInfo is for general information
	LogLevelInfo
	// LogLevelWarn is for warning messages
	LogLevelWarn
	// LogLevelError is for error messages
	LogLevelError
)

// Logger is a custom logger with log levels
type Logger struct {
	logger *log.Logger
	level  LogLevel
	mu     sync.Mutex
}

// NewLogger creates a new Logger instance
func NewLogger(out io.Writer, prefix string, flag int, level LogLevel) *Logger {
	return &Logger{
		logger: log.New(out, prefix, flag),
		level:  level,
	}
}

// SetLevel sets the logging level
func (l *Logger) SetLevel(level LogLevel) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.level = level
}

// Debug logs a debug message
func (l *Logger) Debug(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelDebug {
		l.logger.Printf("[DEBUG] "+format, v...)
	}
}

// Info logs an info message
func (l *Logger) Info(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelInfo {
		l.logger.Printf("[INFO] "+format, v...)
	}
}

// Warn logs a warning message
func (l *Logger) Warn(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelWarn {
		l.logger.Printf("[WARN] "+format, v...)
	}
}

// Error logs an error message
func (l *Logger) Error(format string, v ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.level <= LogLevelError {
		l.logger.Printf("[ERROR] "+format, v...)
	}
}

// DefaultLogger is the default logger instance
var DefaultLogger = NewLogger(os.Stderr, "", log.LstdFlags, LogLevelInfo)

// SetDefaultLogger sets the default logger
func SetDefaultLogger(logger *Logger) {
	DefaultLogger = logger
}

// SetDefaultLogLevel sets the log level for the default logger
func SetDefaultLogLevel(level LogLevel) {
	DefaultLogger.SetLevel(level)
}

// Debug logs a debug message using the default logger
func Debug(format string, v ...interface{}) {
	DefaultLogger.Debug(format, v...)
}

// Info logs an info message using the default logger
func Info(format string, v ...interface{}) {
	DefaultLogger.Info(format, v...)
}

// Warn logs a warning message using the default logger
func Warn(format string, v ...interface{}) {
	DefaultLogger.Warn(format, v...)
}

// Error logs an error message using the default logger
func Error(format string, v ...interface{}) {
	DefaultLogger.Error(format, v...)
}
