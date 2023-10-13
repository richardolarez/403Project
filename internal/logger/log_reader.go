// logger/logreader.go

package logger

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// A log entry in the log file
type Log struct {
	Level          logger.logLevel        `json:"level"`
	Timestamp      string                 `json:"timestamp"`
	Message        string                 `json:message"`
	AdditionalData map[string]interface{} `json:"additional_data"`
}

// LogReader reads and processes the log files.
type LogReader struct {
	logDir string
}

// NewLogReader creates a new LogReader instance.
func NewLogReader(logDir string) *LogReader {
	return &LogReader{logDir: logDir}
}

// Reads the log entries from the log file.
func (logR *LogReader) ReadLogs() ([]Log, error) {
	filePath := filepath.Join(logR.logDir, "database.json")
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var logs []Log
	decoder := json.NewDecoder(file)

	for decoder.More() {
		var log Log
		if err := decoder.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}
