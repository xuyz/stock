package slog

import (
	"io"

	"github.com/sirupsen/logrus"
)

type Level uint32

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `os.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
)

type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	WithField(key string, value interface{}) Logger
}

type logger struct {
	log *logrus.Entry
}

func NewLogger(out io.Writer, lvl Level) Logger {
	ll := &logrus.Logger{
		Formatter: new(logrus.TextFormatter),
		Hooks:     make(logrus.LevelHooks),
		Out:       out,
		Level:     logrus.Level(lvl),
	}
	logrus.Warnf("[NewLogger] %v", ll)
	l := logrus.NewEntry(ll)

	log := &logger{log: l}

	return log
}

func (l *logger) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logger) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *logger) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logger) WithField(key string, value interface{}) (log Logger) {
	log = &logger{
		log: l.log.WithField(key, value),
	}
	return
}
