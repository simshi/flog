package flog

import (
	"io"
	"os"
)

var globalWriter io.Writer = os.Stdout

func SetOutput(w io.Writer) {
	if w != nil {
		globalWriter = w
	} else {
		globalWriter = io.Discard
	}
}
