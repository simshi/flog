package flog

type IEntry interface {
	Init(l Level, skip int) IEntry

	// finishing move
	Msg(m string)
	Msgf(f string, args ...any)

	Any(k string, v any) IEntry

	// Int[T AnyInt](k string, v T) IEntry
	Int(k string, v int) IEntry
	Int8(k string, v int8) IEntry
	Int16(k string, v int16) IEntry
	Int32(k string, v int32) IEntry
	Int64(k string, v int64) IEntry

	// Uint[T AnyUint](k string, v T) IEntry
	Uint(k string, v uint) IEntry
	Uint8(k string, v uint8) IEntry
	Uint16(k string, v uint16) IEntry
	Uint32(k string, v uint32) IEntry
	Uint64(k string, v uint64) IEntry

	IntPad0(k string, v int, pad int) IEntry
	Int8Pad0(k string, v int8, pad int) IEntry
	Int16Pad0(k string, v int16, pad int) IEntry
	Int32Pad0(k string, v int32, pad int) IEntry
	Int64Pad0(k string, v int64, pad int) IEntry

	UintPad0(k string, v uint, pad int) IEntry
	Uint8Pad0(k string, v uint8, pad int) IEntry
	Uint16Pad0(k string, v uint16, pad int) IEntry
	Uint32Pad0(k string, v uint32, pad int) IEntry
	Uint64Pad0(k string, v uint64, pad int) IEntry

	Float32(k string, v float32) IEntry
	Float64(k string, v float64) IEntry

	Hex(k string, v int) IEntry

	Bool(k string, v bool) IEntry
	Str(k string, v string) IEntry
	Err(v error) IEntry
}
