package logger

import (
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

var _level = map[string]log.Level{
	"INFO":  log.InfoLevel,
	"DEBUG": log.DebugLevel,
	"WARN":  log.WarnLevel,
	"PANIC": log.PanicLevel,
	"TRACE": log.TraceLevel,
	"FATAL": log.FatalLevel,
	"ERROR": log.ErrorLevel,
}

type Logger struct {
	*log.Entry
}

type Fields log.Fields

// New : param "level" could be INFO, DEBUG, WARN, PANIC, TRACE, FATAL or ERROR. Otherwise must be setted DEBUG
func New(environment, level string) Logger {
	logger := log.New()
	logger.SetFormatter(&log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
	})
	logLevel, exist := _level[strings.ToUpper(level)]
	if !exist {
		logger.SetLevel(log.DebugLevel)
	}
	logger.SetLevel(logLevel)
	return Logger{logger.WithFields(log.Fields{"environment": environment})}
}

func (l Logger) WithFields(f Fields) Logger {
	return l.WithFields(f)
}
