package flog

import "os"

// for Fatal log exit
type ExitEntry struct {
	*Entry
}

var _ IEntry = &ExitEntry{}

func (e *ExitEntry) Msg(m string) {
	e.Entry.Msg(m)
	os.Exit(-1)
}

func (e *ExitEntry) Msgf(format string, v ...any) {
	e.Entry.Msgf(format, v...)
	os.Exit(-1)
}

func (e *ExitEntry) Init(l Level, skip int) IEntry {
	e.Entry.Init(LEVEL_FATAL, skip+1)
	return e
}
func (e *ExitEntry) Int(k string, v int) IEntry {
	e.Entry.Int(k, v)
	return e
}
func (e *ExitEntry) Int8(k string, v int8) IEntry {
	e.Entry.Int8(k, v)
	return e
}
func (e *ExitEntry) Int16(k string, v int16) IEntry {
	e.Entry.Int16(k, v)
	return e
}
func (e *ExitEntry) Int32(k string, v int32) IEntry {
	e.Entry.Int32(k, v)
	return e
}
func (e *ExitEntry) Int64(k string, v int64) IEntry {
	e.Entry.Int64(k, v)
	return e
}
func (e *ExitEntry) Uint(k string, v uint) IEntry {
	e.Entry.Uint(k, v)
	return e
}
func (e *ExitEntry) Uint8(k string, v uint8) IEntry {
	e.Entry.Uint8(k, v)
	return e
}
func (e *ExitEntry) Uint16(k string, v uint16) IEntry {
	e.Entry.Uint16(k, v)
	return e
}
func (e *ExitEntry) Uint32(k string, v uint32) IEntry {
	e.Entry.Uint32(k, v)
	return e
}
func (e *ExitEntry) Uint64(k string, v uint64) IEntry {
	e.Entry.Uint64(k, v)
	return e
}
func (e *ExitEntry) Float32(k string, v float32) IEntry {
	e.Entry.Float32(k, v)
	return e
}
func (e *ExitEntry) Float64(k string, v float64) IEntry {
	e.Entry.Float64(k, v)
	return e
}

func (e *ExitEntry) IntPad0(k string, v int, pad int) IEntry {
	e.Entry.IntPad0(k, v, pad)
	return e
}
func (e *ExitEntry) Int8Pad0(k string, v int8, pad int) IEntry {
	e.Entry.Int8Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Int16Pad0(k string, v int16, pad int) IEntry {
	e.Entry.Int16Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Int32Pad0(k string, v int32, pad int) IEntry {
	e.Entry.Int32Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Int64Pad0(k string, v int64, pad int) IEntry {
	e.Entry.Int64Pad0(k, v, pad)
	return e
}

func (e *ExitEntry) UintPad0(k string, v uint, pad int) IEntry {
	e.Entry.UintPad0(k, v, pad)
	return e
}
func (e *ExitEntry) Uint8Pad0(k string, v uint8, pad int) IEntry {
	e.Entry.Uint8Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Uint16Pad0(k string, v uint16, pad int) IEntry {
	e.Entry.Uint16Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Uint32Pad0(k string, v uint32, pad int) IEntry {
	e.Entry.Uint32Pad0(k, v, pad)
	return e
}
func (e *ExitEntry) Uint64Pad0(k string, v uint64, pad int) IEntry {
	e.Entry.Uint64Pad0(k, v, pad)
	return e
}

func (e *ExitEntry) Hex(k string, v int) IEntry {
	e.Entry.Hex(k, v)
	return e
}
func (e *ExitEntry) Bool(k string, v bool) IEntry {
	e.Entry.Bool(k, v)
	return e
}
func (e *ExitEntry) Str(k string, v string) IEntry {
	e.Entry.Str(k, v)
	return e
}
func (e *ExitEntry) Any(k string, v any) IEntry {
	e.Entry.Any(k, v)
	return e
}
func (e *ExitEntry) Err(v error) IEntry {
	e.Entry.Err(v)
	return e
}
