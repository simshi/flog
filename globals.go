package flog

import (
	"io"
	"os"
	"sync"
)

var (
	gWriter io.Writer = os.Stdout
	gLevel  Level     = LEVEL_INFO
)

var gEntryPool = sync.Pool{
	New: func() interface{} {
		return new(Entry)
	},
}

var gNopEntry = NopEntry{}

func SetOutput(w io.Writer) {
	if w != nil {
		gWriter = w
	} else {
		gWriter = io.Discard
	}
}

func SetLevel(lvl Level) {
	gLevel = lvl
}
