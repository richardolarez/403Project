// logger/logger.go

package logger

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// LogLevel represents the severity/ type of log entry
type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
)

// LogEntry represents the entry to be logged.
type LogEntry struct {
	Level     LogLevel `json:"level"`
	Timestamp string   `json:"timestamp"`
	Message   string   `json:"message"`
	// Add more fields here if needed
}

type Logger struct {
	mu     sync.Mutex
	logDir string
	logs   []LogEntry
}

// Creates a new Logger instance
func newLogger(logDir string) *Logger {
	return &Logger{
		logDir: logDir,
	}
}

func (l *Logger) Log(level LogLevel, message string, additionalData map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Get current time
	timestamp := time.Now().Format("2023-01-01 15:03:02")

	logEntry := LogEntry{
		Timestamp: timestamp,
		Level:     level,
		Message:   message,
	}

	for key, value := range additionalData {
		logEntry[key] = value
	}

	logEntryJSON, err := json.Marshal(logEntry)
	if err != nil {
		return
	}

	err = l.saveLogEntryToFle(logEntryJSON)
	if err != nil {
		return
	}

	// Print the log for testing
	// fmt.Println(string(logEntryJSON))
}

func (l *Logger) saveLogEntryToFile(logEntryJSON []byte) error {
	// Write the log entry to a log file
	filename := l.logDir + "/logs.json"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(logEntryJSON)
	if err != nil {
		return err
	}

	_, err = file.WriteString("\n")
	if err != nil {
		return err
	}

	return nil
}
