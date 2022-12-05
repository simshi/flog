package log

import (
	"io"

	"github.com/simshi/flog"
)

// global _log
var _log = flog.New()

func Debug() flog.IEntry {
	return _log.Debug(1)
}
func Debugf(format string, args ...any) {
	_log.Debug(1).Msgf(format, args...)
}

func Info() flog.IEntry {
	return _log.Info(1)
}
func Infof(format string, args ...any) {
	_log.Info(1).Msgf(format, args...)
}

func Warn() flog.IEntry {
	return _log.Warn(1)
}
func Warnf(format string, args ...any) {
	_log.Warn(1).Msgf(format, args...)
}

func Error() flog.IEntry {
	return _log.Error(1)
}
func Errorf(format string, args ...any) {
	_log.Error(1).Msgf(format, args...)
}

func Fatal() flog.IEntry {
	return _log.Fatal(1)
}
func Fatalf(format string, args ...any) {
	_log.Fatal(1).Msgf(format, args...)
}

func SetOutput(w io.Writer) {
	flog.SetOutput(w)
}

func SetLevel(lvl string) error {
	if level, err := flog.ParseLevel(lvl); err != nil {
		return err
	} else {
		flog.SetLevel(level)
	}
	return nil
}
