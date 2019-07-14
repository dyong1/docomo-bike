package logging

import (
	"fmt"
	"io"

	glogger "github.com/google/logger"
)

func New(name string, verbose bool, systemLog bool, logFile io.Writer, useDebugLog bool) Logger {
	return &GLogger{
		glogger:     glogger.Init(name, verbose, systemLog, logFile),
		useDebugLog: useDebugLog,
	}
}

type Logger interface {
	Debug(v ...interface{})
	DebugDepth(depth int, v ...interface{})
	Debugln(v ...interface{})
	Debugf(format string, v ...interface{})
	Info(v ...interface{})
	InfoDepth(depth int, v ...interface{})
	Infoln(v ...interface{})
	Infof(format string, v ...interface{})
	Warning(v ...interface{})
	WarningDepth(depth int, v ...interface{})
	Warningln(v ...interface{})
	Warningf(format string, v ...interface{})
	Error(v ...interface{})
	ErrorDepth(depth int, v ...interface{})
	Errorln(v ...interface{})
	Errorf(format string, v ...interface{})
	Fatal(v ...interface{})
	FatalDepth(depth int, v ...interface{})
	Fatalln(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type GLogger struct {
	glogger     *glogger.Logger
	useDebugLog bool
}

func (l *GLogger) Debug(v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.glogger.Info(fmt.Sprintf("DEBUG : %s", fmt.Sprint(v...)))
}
func (l *GLogger) Debugf(format string, v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.glogger.Infof(fmt.Sprintf("DEBUG : %s", fmt.Sprintf(format, v...)))
}
func (l *GLogger) Debugln(v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.glogger.Infoln(fmt.Sprintf("DEBUG : %s", fmt.Sprintln(v...)))
}
func (l *GLogger) DebugDepth(depth int, v ...interface{}) {
	if !l.useDebugLog {
		return
	}
	l.glogger.InfoDepth(depth, fmt.Sprintf("DEBUG : %s", fmt.Sprint(v...)))
}
func (l *GLogger) Info(v ...interface{}) {
	l.glogger.Info(v...)
}
func (l *GLogger) InfoDepth(depth int, v ...interface{}) {
	l.glogger.InfoDepth(depth, v...)
}
func (l *GLogger) Infoln(v ...interface{}) {
	l.glogger.Infoln(v...)
}
func (l *GLogger) Infof(format string, v ...interface{}) {
	l.glogger.Infof(format, v...)
}
func (l *GLogger) Warning(v ...interface{}) {
	l.glogger.Warning(v...)
}
func (l *GLogger) WarningDepth(depth int, v ...interface{}) {
	l.glogger.WarningDepth(depth, v...)
}
func (l *GLogger) Warningln(v ...interface{}) {
	l.glogger.Warningln(v...)
}
func (l *GLogger) Warningf(format string, v ...interface{}) {
	l.glogger.Warningf(format, v...)
}
func (l *GLogger) Error(v ...interface{}) {
	l.glogger.Error(v...)
}
func (l *GLogger) ErrorDepth(depth int, v ...interface{}) {
	l.glogger.ErrorDepth(depth, v...)
}
func (l *GLogger) Errorln(v ...interface{}) {
	l.glogger.Errorln(v...)
}
func (l *GLogger) Errorf(format string, v ...interface{}) {
	l.glogger.Errorf(format, v...)
}
func (l *GLogger) Fatal(v ...interface{}) {
	l.glogger.Fatal(v...)
}
func (l *GLogger) FatalDepth(depth int, v ...interface{}) {
	l.glogger.FatalDepth(depth, v...)
}
func (l *GLogger) Fatalln(v ...interface{}) {
	l.glogger.Fatalln(v...)
}
func (l *GLogger) Fatalf(format string, v ...interface{}) {
	l.glogger.Fatalf(format, v...)
}
