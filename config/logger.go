package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(writer io.Writer, prefix string) *Logger {
	if writer == nil { writer = os.Stdout }

	return &Logger {
		debug:   log.New(writer, prefix+"DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:    log.New(writer, prefix+"INFO: ", log.Ldate|log.Ltime),
		warning: log.New(writer, prefix+"WARNING: ", log.Ldate|log.Ltime),
		err:     log.New(writer, prefix+"ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		writer:  writer,
	}
}

// Non-Formatted Logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Println(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
}

// Formatted Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
}
