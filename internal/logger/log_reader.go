// logger/logreader.go

package logger

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
)

// A log entry in the log file
type Log struct {
	Level          LogLevel               `json:"level"`
	Timestamp      string                 `json:"timestamp"`
	Message        string                 `json:"message"`
	AdditionalData map[string]interface{} `json:"additional_data"`
}

// LogReader reads and processes the log files.
type LogReader struct {
	logDir       string
	levelFilter  LogLevel
	searchFilter string
}

// NewLogReader creates a new LogReader instance. Must search by LogLevel (see logger.go for levels)
// and input a search parameter.
func NewLogReader(logDir string, levelFilter LogLevel, searchFilter string) *LogReader {
	return &LogReader{logDir: logDir, levelFilter: levelFilter, searchFilter: searchFilter}
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
		if logR.passesFilters(&log) {
			logs = append(logs, log)
		}
	}
	return logs, nil
}

// Checks if a log entry passes the filters.
func (logR *LogReader) passesFilters(log *Log) bool {
	if logR.levelFilter != 0 && log.Level != logR.levelFilter {
		return false
	}

	if logR.searchFilter != "" && !strings.Contains(log.Message, logR.searchFilter) {
		return false
	}

	return true
}
