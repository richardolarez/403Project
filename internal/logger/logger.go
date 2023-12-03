// logger/logger.go

package logger

import (
	"encoding/json"
	"io"
	"os"
	"sync"
	"time"
)

// LogLevel represents the severity/ type of log entry
type LogLevel int

const (
	Debug   LogLevel = 1
	Info    LogLevel = 2
	Warning LogLevel = 3
	Error   LogLevel = 4
)

// LogEntry represents the entry to be logged.
type LogEntry struct {
	Level          LogLevel               `json:"level"`
	Timestamp      string                 `json:"timestamp"`
	Message        string                 `json:"message"`
	Additionaldata map[string]interface{} `json:"additional_data"`
	// Add more fields here if needed
}

type Logger struct {
	mu     sync.Mutex
	logDir string
}

// Creates a new Logger instance
func NewLogger(logDir string) *Logger {
	// fmt.Print("NewLogger entered")
	return &Logger{
		logDir: logDir,
	}
}

func (l *Logger) Log(level LogLevel, message string, additionalData map[string]interface{}) {
	// fmt.Println("Log entered")
	l.mu.Lock()
	defer l.mu.Unlock()

	// Get current time
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	logEntry := LogEntry{
		Timestamp:      timestamp,
		Level:          level,
		Message:        message,
		Additionaldata: additionalData,
	}

	logEntryJSON, err := json.Marshal(logEntry)
	if err != nil {
		return
	}

	err = l.saveLogEntryToFile(logEntryJSON)
	if err != nil {
		return
	}

	//Print the log for testing
	// fmt.Println(string(logEntryJSON))
}

type LogArray struct {
	Logs []LogEntry `json:"logs"`
}

func (l *Logger) saveLogEntryToFile(logEntryJSON []byte) error {
	// fmt.Println("saveLogEntryToFile entered")

	// Open the log file or create it if it doesn't exist and append to it
	file, err := os.OpenFile("./db/logs.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode the existing logs from the file
	var logArray LogArray
	err = json.NewDecoder(file).Decode(&logArray)
	if err != nil && err != io.EOF {
		return err
	}

	// Append the new log entry to the logs array
	var logEntry LogEntry
	err = json.Unmarshal(logEntryJSON, &logEntry)
	if err != nil {
		return err
	}
	logArray.Logs = append(logArray.Logs, logEntry)

	// Write the updated logs array to the file
	file.Seek(0, 0)
	file.Truncate(0)
	err = json.NewEncoder(file).Encode(logArray)
	if err != nil {
		return err
	}

	return nil
}
