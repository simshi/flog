package flog

type NopEntry struct {
}

func (e *NopEntry) Init(l Level, skip int) IEntry {
	return e
}
func (e *NopEntry) Int(k string, v int) IEntry {
	return e
}
func (e *NopEntry) Int8(k string, v int8) IEntry {
	return e
}
func (e *NopEntry) Int16(k string, v int16) IEntry {
	return e
}
func (e *NopEntry) Int32(k string, v int32) IEntry {
	return e
}
func (e *NopEntry) Int64(k string, v int64) IEntry {
	return e
}
func (e *NopEntry) Uint(k string, v uint) IEntry {
	return e
}
func (e *NopEntry) Uint8(k string, v uint8) IEntry {
	return e
}
func (e *NopEntry) Uint16(k string, v uint16) IEntry {
	return e
}
func (e *NopEntry) Uint32(k string, v uint32) IEntry {
	return e
}
func (e *NopEntry) Uint64(k string, v uint64) IEntry {
	return e
}
func (e *NopEntry) Float32(k string, v float32) IEntry {
	return e
}
func (e *NopEntry) Float64(k string, v float64) IEntry {
	return e
}

func (e *NopEntry) IntPad0(k string, v int, pad int) IEntry {
	return e
}
func (e *NopEntry) Int8Pad0(k string, v int8, pad int) IEntry {
	return e
}
func (e *NopEntry) Int16Pad0(k string, v int16, pad int) IEntry {
	return e
}
func (e *NopEntry) Int32Pad0(k string, v int32, pad int) IEntry {
	return e
}
func (e *NopEntry) Int64Pad0(k string, v int64, pad int) IEntry {
	return e
}

func (e *NopEntry) UintPad0(k string, v uint, pad int) IEntry {
	return e
}
func (e *NopEntry) Uint8Pad0(k string, v uint8, pad int) IEntry {
	return e
}
func (e *NopEntry) Uint16Pad0(k string, v uint16, pad int) IEntry {
	return e
}
func (e *NopEntry) Uint32Pad0(k string, v uint32, pad int) IEntry {
	return e
}
func (e *NopEntry) Uint64Pad0(k string, v uint64, pad int) IEntry {
	return e
}

func (e *NopEntry) Hex(k string, v int) IEntry {
	return e
}
func (e *NopEntry) Bool(k string, v bool) IEntry {
	return e
}
func (e *NopEntry) Str(k string, v string) IEntry {
	return e
}
func (e *NopEntry) Any(k string, v any) IEntry {
	return e
}
func (e *NopEntry) Err(v error) IEntry {
	return e
}

func (e *NopEntry) Msg(m string) {
}
func (e *NopEntry) Msgf(format string, v ...any) {
}
