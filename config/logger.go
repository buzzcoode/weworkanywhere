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

func NewLooger(p string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "DEBUG: ", logger.Flags()),
		info:    log.New(writer, "INFO: ", logger.Flags()),
		warning: log.New(writer, "WARNING: ", logger.Flags()),
		err:     log.New(writer, "ERROR: ", logger.Flags()),
		writer:  writer,
	}
}

// Create No-formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Panicln(v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.info.Panicln(v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.warning.Panicln(v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.err.Panicln(v...)
}

// Create Format Enabled Logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Panicf(format, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Panicf(format, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.warning.Panicf(format, v...)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Panicf(format, v...)
}