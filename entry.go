package flog

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const FIXED_ENTRY_SIZE = 1024 - 16
const TIME_LAYOUT = "2006-01-02T15:04:05.000Z07:00"

type Entry struct {
	a          [FIXED_ENTRY_SIZE]byte
	pos        int
	cachedTime time.Time
}

func NewEntry() *Entry {
	return &Entry{}
}

func (e *Entry) Init(l Level, skip int) *Entry {
	e.pos = 0
	e.writeTime(time.Now())
	e.writeSep()
	e.writeStr(levelConsoleMap[l])
	e.writeSep()
	e.writeCaller(skip + 1)
	return e
}

// finishing move
func (e *Entry) Msg(m string) {
	if e == nil {
		return
	}

	e.writeSep()
	e.pos += copy(e.a[e.pos:], m)
	e.writeByte(byte('\n'))
	gWriter.Write(e.a[:e.pos])
	gEntryPool.Put(e)
}
func (e *Entry) Msgf(format string, v ...any) {
	if e == nil {
		return
	}

	e.writeSep()
	e.pos += copy(e.a[e.pos:], fmt.Sprintf(format, v...))
	e.writeByte(byte('\n'))
	gWriter.Write(e.a[:e.pos])
	gEntryPool.Put(e)
}

// for Fatal log exit
type ExitEntry struct {
	*Entry
}

func (e *ExitEntry) Msg(m string) {
	e.Entry.Msg(m)
	os.Exit(-1)
}
func (e *ExitEntry) Msgf(format string, v ...any) {
	e.Entry.Msgf(format, v...)
	os.Exit(-1)
}

// chainable methods
type IntSet interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}
type UintSet interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Go has not support type-parameterized methods yet
// TODO: change this to a method `Entry.Int[T IntSet](k, v)` after Go became smart
func anyInt[T IntSet](e *Entry, k string, v T) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	writeAnyInt(e, v)
	return e
}
func anyUint[T UintSet](e *Entry, k string, v T) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	writeAnyUint(e, v)
	return e
}

func anyIntPad0[T IntSet](e *Entry, k string, v T, pad int) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	writeAnyIntPad0(e, v, pad)
	return e
}
func anyUintPad0[T UintSet](e *Entry, k string, v T, pad int) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	writeAnyUintPad0(e, v, pad)
	return e
}

func (e *Entry) Int(k string, v int) *Entry {
	return anyInt(e, k, v)
}
func (e *Entry) Int8(k string, v int8) *Entry {
	return anyInt(e, k, v)
}
func (e *Entry) Int16(k string, v int16) *Entry {
	return anyInt(e, k, v)
}
func (e *Entry) Int32(k string, v int32) *Entry {
	return anyInt(e, k, v)
}
func (e *Entry) Int64(k string, v int64) *Entry {
	return anyInt(e, k, v)
}

func (e *Entry) Uint(k string, v uint) *Entry {
	return anyUint(e, k, v)
}
func (e *Entry) Uint8(k string, v uint8) *Entry {
	return anyUint(e, k, v)
}
func (e *Entry) Uint16(k string, v uint16) *Entry {
	return anyUint(e, k, v)
}
func (e *Entry) Uint32(k string, v uint32) *Entry {
	return anyUint(e, k, v)
}
func (e *Entry) Uint64(k string, v uint64) *Entry {
	return anyUint(e, k, v)
}

func (e *Entry) IntPad0(k string, v int, pad int) *Entry {
	return anyIntPad0(e, k, v, pad)
}
func (e *Entry) Int8Pad0(k string, v int8, pad int) *Entry {
	return anyIntPad0(e, k, v, pad)
}
func (e *Entry) Int16Pad0(k string, v int16, pad int) *Entry {
	return anyIntPad0(e, k, v, pad)
}
func (e *Entry) Int32Pad0(k string, v int32, pad int) *Entry {
	return anyIntPad0(e, k, v, pad)
}
func (e *Entry) Int64Pad0(k string, v int64, pad int) *Entry {
	return anyIntPad0(e, k, v, pad)
}

func (e *Entry) UintPad0(k string, v uint, pad int) *Entry {
	return anyUintPad0(e, k, v, pad)
}
func (e *Entry) Uint8Pad0(k string, v uint8, pad int) *Entry {
	return anyUintPad0(e, k, v, pad)
}
func (e *Entry) Uint16Pad0(k string, v uint16, pad int) *Entry {
	return anyUintPad0(e, k, v, pad)
}
func (e *Entry) Uint32Pad0(k string, v uint32, pad int) *Entry {
	return anyUintPad0(e, k, v, pad)
}
func (e *Entry) Uint64Pad0(k string, v uint64, pad int) *Entry {
	return anyUintPad0(e, k, v, pad)
}

func (e *Entry) Hex(k string, v int) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	e.writeStr("0x")
	e.writeHex64(uint64(v))
	return e
}

func (e *Entry) Bool(k string, v bool) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	if v {
		e.writeStr("true")
	} else {
		e.writeStr("false")
	}
	return e
}

func (e *Entry) Float32(k string, v float32) *Entry {
	if e == nil {
		return e
	}

	return e.appendFloat64(k, float64(v), 32)
}
func (e *Entry) Float64(k string, v float64) *Entry {
	if e == nil {
		return e
	}

	return e.appendFloat64(k, v, 64)
}
func (e *Entry) appendFloat64(k string, v float64, bits int) *Entry {
	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	e.writeFloat64(v, bits)
	return e
}

func (e *Entry) Str(k, s string) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	e.writeStr(s)
	return e
}

func (e *Entry) Err(err error) *Entry {
	if e == nil {
		return e
	}

	return e.Str("err", err.Error())
}

func (e *Entry) Any(k string, v any) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	e.writeStr(fmt.Sprintf("%v", v))
	return e
}

// private helper functions
func (e *Entry) writeByte(b byte) {
	e.a[e.pos] = b
	e.pos += 1
}
func (e *Entry) writeTime(t time.Time) {
	if t.Unix()/60 == e.cachedTime.Unix()/60 {
		e.pos = 17
		writeAnyUintPad0(e, uint(t.Second()), 2)
		e.pos += 1
		writeAnyUintPad0(e, uint(t.Nanosecond()/1000000), 3)
		if e.a[e.pos] != byte('Z') {
			e.pos += 6 // +08:00
		} else {
			e.pos += 1 // 'Z'
		}
	} else {
		// e.pos += copy(e.a[:], t.Format(TIME_LAYOUT))
		b := t.AppendFormat(e.a[e.pos:e.pos], TIME_LAYOUT)
		e.pos += len(b)

		e.cachedTime = t
	}
}
func (e *Entry) writeSep() {
	e.a[e.pos] = byte(' ')
	e.pos += 1
}
func (e *Entry) writeDelimar() {
	e.a[e.pos] = byte('=')
	e.pos += 1
}

func (e *Entry) writeStr(s string) {
	e.pos += copy(e.a[e.pos:], s)
}

// int to string conversion
var DIGITS = []byte("0123456789ABCDEF")
var DIGITS_INT = []byte(".FEDCBA9876543210123456789ABCDEF")

// Efficient Integer to String Conversions, by Matthew Wilson.
func writeAnyUint[T UintSet](e *Entry, v T) {
	s := e.pos
	for {
		e.a[e.pos] = DIGITS[v%10]
		e.pos += 1
		v /= 10
		if v == 0 {
			break
		}
	}

	reverseBytes(e.a[s:e.pos])
}
func writeAnyInt[T IntSet](e *Entry, v T) {
	if v < 0 {
		e.a[e.pos] = byte('-')
		e.pos += 1
	}
	s := e.pos
	for {
		e.a[e.pos] = DIGITS_INT[16+v%10]
		e.pos += 1
		v /= 10
		if v == 0 {
			break
		}
	}

	reverseBytes(e.a[s:e.pos])
}

func writeAnyUintPad0[T UintSet](e *Entry, v T, pad int) {
	s := e.pos
	for {
		e.a[e.pos] = DIGITS[v%10]
		e.pos += 1
		v /= 10
		if v == 0 {
			break
		}
	}

	for pad -= (e.pos - s); pad > 0; pad -= 1 {
		e.a[e.pos] = byte('0')
		e.pos += 1
	}

	reverseBytes(e.a[s:e.pos])
}
func writeAnyIntPad0[T IntSet](e *Entry, v T, pad int) {
	if v < 0 {
		e.a[e.pos] = byte('-')
		e.pos += 1
		pad -= 1
	}

	s := e.pos
	for {
		e.a[e.pos] = DIGITS_INT[16+v%10]
		e.pos += 1
		v /= 10
		if v == 0 {
			break
		}
	}

	for pad -= (e.pos - s); pad > 0; pad -= 1 {
		e.a[e.pos] = byte('0')
		e.pos += 1
	}

	reverseBytes(e.a[s:e.pos])
}

func (e *Entry) writeHex64(v uint64) {
	s := e.pos
	for {
		e.a[e.pos] = DIGITS[v%16]
		e.pos += 1
		v /= 16
		if v == 0 {
			break
		}
	}

	reverseBytes(e.a[s:e.pos])
}
func reverseBytes(b []byte) {
	j := len(b) - 1
	for i := 0; i < j; i += 1 {
		b[i], b[j] = b[j], b[i]
		j -= 1
	}
}

func (e *Entry) writeFloat64(v float64, bits int) {
	b := strconv.AppendFloat(e.a[e.pos:e.pos], v, 'f', -1, bits)
	e.pos += len(b)
}

func (e *Entry) writeCaller(skip int) {
	_, file, line, _ := runtime.Caller(skip + 1)
	if index := strings.LastIndex(file, "/"); index != -1 {
		if begin := strings.LastIndex(file[:index], "/"); begin != -1 {
			e.pos += copy(e.a[e.pos:], file[begin+1:])
		} else {
			e.pos += copy(e.a[e.pos:], file)
		}
	} else {
		e.pos += copy(e.a[e.pos:], file)
	}
	e.writeByte(byte(':'))
	writeAnyInt(e, line)
}
