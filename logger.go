package flog

import (
	"sync"
)

var ep = sync.Pool{
	New: func() interface{} {
		return new(Entry)
	},
}

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(skip int) *Entry {
	e := ep.Get().(*Entry)
	return e.Init(LEVEL_DEBUG, skip+1)
}

func (l *Logger) Info(skip int) *Entry {
	e := ep.Get().(*Entry)
	return e.Init(LEVEL_INFO, skip+1)
}

func (l *Logger) Warn(skip int) *Entry {
	e := ep.Get().(*Entry)
	return e.Init(LEVEL_WARN, skip+1)
}

func (l *Logger) Error(skip int) *Entry {
	e := ep.Get().(*Entry)
	return e.Init(LEVEL_ERROR, skip+1)
}

func (l *Logger) Fatal(skip int) *Entry {
	e := &ExitEntry{
		ep.Get().(*Entry),
	}
	return e.Init(LEVEL_FATAL, skip+1)
}
