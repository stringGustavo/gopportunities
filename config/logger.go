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
		debug:   log.New(writer, prefix + "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
		info:    log.New(writer, prefix + "INFO: ", log.Ldate|log.Ltime),
		warning: log.New(writer, prefix + "WARNING: ", log.Ldate|log.Ltime),
		err:     log.New(writer, prefix + "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		writer:  writer,
	}
}

// Non-Formatted Logs
func (log *Logger) Debug(variadic ...interface{}) {
	log.debug.Println(variadic...)
}

func (log *Logger) Info(variadic ...interface{}) {
	log.info.Println(variadic...)
}

func (log *Logger) Warn(variadic ...interface{}) {
	log.warning.Println(variadic...)
}

func (log *Logger) Error(variadic ...interface{}) {
	log.err.Println(variadic...)
}

// Formatted Logs
func (log *Logger) Debugf(format string, variadic ...interface{}) {
	log.debug.Printf(format, variadic...)
}

func (log *Logger) Infof(format string, variadic ...interface{}) {
	log.info.Printf(format, variadic...)
}

func (log *Logger) Warnf(format string, variadic ...interface{}) {
	log.warning.Printf(format, variadic...)
}

func (log *Logger) Errorf(format string, variadic ...interface{}) {
	log.err.Printf(format, variadic...)
}
