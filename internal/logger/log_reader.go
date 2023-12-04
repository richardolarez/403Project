package logger

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

type Log struct {
	Level          LogLevel               `json:"level"`
	Timestamp      string                 `json:"timestamp"`
	Message        string                 `json:"message"`
	AdditionalData map[string]interface{} `json:"additional_data"`
}

type LogReader struct {
	logDir       string
	levelFilter  LogLevel
	searchFilter string
	startTime    time.Time
	endTime      time.Time
}

func NewLogReader(logDir string, levelFilter LogLevel, searchFilter string, startTime, endTime time.Time) *LogReader {
	return &LogReader{logDir: logDir, levelFilter: levelFilter, searchFilter: searchFilter, startTime: startTime, endTime: endTime}
}

func (logR *LogReader) ReadLogs() ([]Log, error) {
	file, err := os.Open(logR.logDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var database struct {
		Logs []Log `json:"logs"`
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&database)
	if err != nil {
		return nil, err
	}

	var filteredLogs []Log
	for _, log := range database.Logs {
		if logR.passesFilters(&log) {
			filteredLogs = append(filteredLogs, log)
		}
	}

	return filteredLogs, nil
}


func (logR *LogReader) passesFilters(log *Log) bool {
	if logR.levelFilter != 0 && log.Level != logR.levelFilter {
		return false
	}

	if logR.searchFilter != "" && !strings.Contains(log.Message, logR.searchFilter) {
		return false
	}

	logTime, err := time.Parse("2006-01-02 15:04:05", log.Timestamp)
	if err != nil || (logR.startTime != time.Time{} && logTime.Before(logR.startTime)) || (logR.endTime != time.Time{} && logTime.After(logR.endTime)) {
		return false
	}

	return true
}
