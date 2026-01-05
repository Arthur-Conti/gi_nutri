package configs

import (
	"fmt"
	"log"
	"os"
)

type Logger struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	warnLog  *log.Logger
}

var defaultLogger *Logger

func init() {
	defaultLogger = &Logger{
		infoLog:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLog: log.New(os.Stderr, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		warnLog:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func GetLogger() *Logger {
	return defaultLogger
}

func (l *Logger) Info(format string, v ...interface{}) {
	l.infoLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(format string, v ...interface{}) {
	l.errorLog.Output(2, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(format string, v ...interface{}) {
	l.warnLog.Output(2, fmt.Sprintf(format, v...))
}

func Info(format string, v ...interface{}) {
	defaultLogger.Info(format, v...)
}

func Error(format string, v ...interface{}) {
	defaultLogger.Error(format, v...)
}

func Warn(format string, v ...interface{}) {
	defaultLogger.Warn(format, v...)
}

