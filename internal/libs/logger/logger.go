package logger

import (
	"fmt"
	"io"

	glogger "github.com/google/logger"
)

func New(name string, verbose bool, systemLog bool, logFile io.Writer, useDebugLog bool) *Logger {
	return &Logger{
		Logger:      glogger.Init(name, verbose, systemLog, logFile),
		useDebugLog: useDebugLog,
	}
}

type Logger struct {
	*glogger.Logger

	useDebugLog bool
}

func (l *Logger) Debug(v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.Logger.Info(fmt.Sprintf("DEBUG : %s", fmt.Sprint(v...)))
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.Logger.Infof(fmt.Sprintf("DEBUG : %s", fmt.Sprintf(format, v...)))
}
func (l *Logger) Debugln(v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.Logger.Infoln(fmt.Sprintf("DEBUG : %s", fmt.Sprintln(v...)))
}
func (l *Logger) DebugDepth(depth int, v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.Logger.InfoDepth(depth, fmt.Sprintf("DEBUG : %s", fmt.Sprint(v...)))
}
