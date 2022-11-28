package flog

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(skip int) *Entry {
	if gLevel < LEVEL_DEBUG {
		return nil
	}
	e := gEntryPool.Get().(*Entry)
	return e.Init(LEVEL_DEBUG, skip+1)
}

func (l *Logger) Info(skip int) *Entry {
	if gLevel < LEVEL_INFO {
		return nil
	}
	e := gEntryPool.Get().(*Entry)
	return e.Init(LEVEL_INFO, skip+1)
}

func (l *Logger) Warn(skip int) *Entry {
	if gLevel < LEVEL_WARN {
		return nil
	}
	e := gEntryPool.Get().(*Entry)
	return e.Init(LEVEL_WARN, skip+1)
}

func (l *Logger) Error(skip int) *Entry {
	if gLevel < LEVEL_ERROR {
		return nil
	}
	e := gEntryPool.Get().(*Entry)
	return e.Init(LEVEL_ERROR, skip+1)
}

func (l *Logger) Fatal(skip int) *Entry {
	if gLevel < LEVEL_FATAL {
		return nil
	}
	e := &ExitEntry{
		gEntryPool.Get().(*Entry),
	}
	return e.Init(LEVEL_FATAL, skip+1)
}
