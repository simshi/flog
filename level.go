package flog

import (
	"fmt"
	"strconv"
	"strings"
)

type Level int8

const (
	LEVEL_DISABLED = iota - 1
	LEVEL_FATAL
	LEVEL_ERROR
	LEVEL_WARN
	LEVEL_INFO
	LEVEL_DEBUG
)

func (l Level) String() string {
	switch l {
	case LEVEL_DISABLED:
		return "DISABLED"
	case LEVEL_FATAL:
		return "FATAL"
	case LEVEL_ERROR:
		return "ERROR"
	case LEVEL_WARN:
		return "WARN"
	case LEVEL_INFO:
		return "INFO"
	case LEVEL_DEBUG:
		return "DEBUG"
	default:
		return strconv.Itoa(int(l))
	}
}

// ParseLevel converts a level string into a zerolog Level value.
// returns an error if the input string does not match known values.
func ParseLevel(levelStr string) (Level, error) {
	switch strings.ToLower(levelStr) {
	case "disabled":
		return LEVEL_DISABLED, nil
	case "fatal":
		return LEVEL_FATAL, nil
	case "error":
		return LEVEL_ERROR, nil
	case "warn":
		return LEVEL_WARN, nil
	case "info":
		return LEVEL_INFO, nil
	case "debug":
		return LEVEL_DEBUG, nil
	default:
		return LEVEL_INFO, fmt.Errorf("unknown level '%s'", levelStr)
	}
}
