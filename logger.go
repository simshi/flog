package flog

type Logger struct {
}

func New() *Logger {
	return &Logger{}
}

func (l *Logger) Debug(skip int) IEntry {
	if gLevel < LEVEL_DEBUG {
		return &gNopEntry
	}
	e := gEntryPool.Get().(IEntry)
	return e.Init(LEVEL_DEBUG, skip+1)
}

func (l *Logger) Info(skip int) IEntry {
	if gLevel < LEVEL_INFO {
		return &gNopEntry
	}
	e := gEntryPool.Get().(IEntry)
	return e.Init(LEVEL_INFO, skip+1)
}

func (l *Logger) Warn(skip int) IEntry {
	if gLevel < LEVEL_WARN {
		return &gNopEntry
	}
	e := gEntryPool.Get().(IEntry)
	return e.Init(LEVEL_WARN, skip+1)
}

func (l *Logger) Error(skip int) IEntry {
	if gLevel < LEVEL_ERROR {
		return &gNopEntry
	}
	e := gEntryPool.Get().(IEntry)
	return e.Init(LEVEL_ERROR, skip+1)
}

func (l *Logger) Fatal(skip int) IEntry {
	if gLevel < LEVEL_FATAL {
		return &gNopEntry
	}
	e := &ExitEntry{
		gEntryPool.Get().(*Entry),
	}
	e.Init(LEVEL_FATAL, skip+1)
	return e
}
