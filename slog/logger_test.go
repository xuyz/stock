package slog

import (
	"os"
	"testing"
)

func t1(l Logger) {
	l.Debug("in t1")
}

func t2(l Logger) {
	l.Info("in t2")
}
func t3(l Logger) {
	l.Warn("in t3")
}

func TestLogger(t *testing.T) {
	l := NewLogger(os.Stdout, DebugLevel)
	l.Info("start ...")
	t1(l.WithField("t1", "hello"))
	t2(l.WithField("t2", "hello"))
	t3(l.WithField("t3", "hello").WithField("t4", "hello"))
}
