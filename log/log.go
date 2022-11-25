package log

import (
	"flog"
	"io"
)

// global _log
var _log = flog.New()

func Debug() *flog.Entry {
	return _log.Debug(1)
}
func Info() *flog.Entry {
	return _log.Info(1)
}
func Warn() *flog.Entry {
	return _log.Warn(1)
}
func Error() *flog.Entry {
	return _log.Error(1)
}
func Fatal() *flog.Entry {
	return _log.Fatal(1)
}

func SetOutput(w io.Writer) {
	flog.SetOutput(w)
}
