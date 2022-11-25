package flog

import "strconv"

type Level int8

const (
	LEVEL_DISABLED = -1
	LEVEL_FATAL    = iota - 1
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
// func ParseLevel(levelStr string) (Level, error) {
// 	switch strings.ToLower(levelStr) {
// 	}
// 	return
// }
