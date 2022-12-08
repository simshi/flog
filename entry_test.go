package flog

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

// `Z` or `+08:00` depends on whether we're in UTC
var TIME_STRLEN = len(time.Now().Format(TIME_LAYOUT))

const FILENAME_LEN = 20
const fakeMessage = "Test logging, but use a somewhat realistic message length."

func setup() {
	SetOutput(io.Discard)
	SetLevel(LEVEL_DEBUG)
}

func TestEntryInit(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)
	checkPos(t, e, TIME_STRLEN+1+14+FILENAME_LEN+2)
	checkTimeStr(t, e)
}
func TestEntryMsg(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)
	pos := TIME_STRLEN + 1 + 14 + FILENAME_LEN + 2
	checkPos(t, e, pos)
	checkTimeStr(t, e)

	e.Msg("abc")
	pos += 1 + 3 + 1
	checkPos(t, e, pos)
	checkTimeStr(t, e)
	checkStrContains(t, e, " abc\n")
}

func TestEntryInt(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)

	e = e.Int("key1", 32).(*Entry)
	checkStrContains(t, e, " key1=32")

	e = e.Int("key2", -233).(*Entry)
	checkStrContains(t, e, " key2=-233")

	e.Msg("int")
	checkStrContains(t, e, " int\n")
}
func TestEntryUint(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)

	e = e.Uint("key1", 32).(*Entry)
	checkStrContains(t, e, " key1=32")

	e = e.Uint8("k8", 233).(*Entry)
	checkStrContains(t, e, " k8=233")

	e = e.Uint32("k32", 12345).(*Entry)
	checkStrContains(t, e, " k32=12345")

	e = e.Uint64("k64", 123456).(*Entry)
	checkStrContains(t, e, " k64=123456")

	e.Msg("int")
	checkStrContains(t, e, " int\n")
}
func TestEntryIntPad0(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)

	e = e.IntPad0("k1", 32, 1).(*Entry)
	checkStrContains(t, e, " k1=32")

	e = e.IntPad0("k2", 32, 2).(*Entry)
	checkStrContains(t, e, " k1=32 k2=32")

	e = e.IntPad0("k3", 32, 3).(*Entry)
	checkStrContains(t, e, " k2=32 k3=032")

	e = e.IntPad0("neg", -32, 5).(*Entry)
	checkStrContains(t, e, " k3=032 neg=-0032")

	e.Msg("pad0")
	checkTimeStr(t, e)
	checkStrContains(t, e, " pad0\n")
}
func TestEntryUintPad0(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)

	e = e.UintPad0("k1", 32, 1).(*Entry)
	checkStrContains(t, e, " k1=32")

	e = e.UintPad0("k2", 32, 2).(*Entry)
	checkStrContains(t, e, " k1=32 k2=32")

	e = e.UintPad0("k3", 32, 3).(*Entry)
	checkStrContains(t, e, " k2=32 k3=032")

	e = e.Uint64Pad0("k64", 123456, 8).(*Entry)
	checkStrContains(t, e, " k64=00123456")

	e.Msg("pad0")
	checkTimeStr(t, e)
	checkStrContains(t, e, " pad0\n")
}

func checkPos(t *testing.T, e *Entry, pos int) {
	if e.pos != pos {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("%d: expected pos=%d, got %d", line, pos, e.pos)
	}
}
func checkTimeStr(t *testing.T, e *Entry) {
	if bytes.Index(e.a[:], []byte("T")) != 10 {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("%d: expected 'T' in time string, got '%s'", line, string(e.a[:]))
	}
	if e.a[23] != 'Z' && e.a[23] != '+' && e.a[23] != '-' {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("%d: expected '+/-/Z' in time string, got '%s'", line, string(e.a[:]))
	}
}
func checkStrContains(t *testing.T, e *Entry, dst string) {
	if !strings.Contains(string(e.a[:]), dst) {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("%d: expected string contains '%s', but got '%s'", line, dst, string(e.a[:]))
	}
}

func BenchmarkInt_fmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("key=%d", i)
	}
}
func BenchmarkInt_append(b *testing.B) {
	b.ReportAllocs()
	buf := make([]byte, 64)
	for i := 0; i < b.N; i++ {
		n := copy(buf, "key=")
		buf = buf[:n]
		buf = strconv.AppendInt(buf, int64(i), 10)
		_ = buf
	}
}
func BenchmarkInt(b *testing.B) {
	setup()
	b.ReportAllocs()
	e := NewEntry()
	for i := 0; i < b.N; i++ {
		e.pos = 0
		e.Int("key", i)
	}
}
func (e *Entry) writeUint64(v uint64) {
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
func (e *Entry) writeInt64(v int64) {
	if v < 0 {
		e.a[e.pos] = byte('-')
		e.pos += 1
		e.writeUint64(uint64(-v))
	} else {
		e.writeUint64(uint64(v))
	}
}
func xInt64(e *Entry, k string, v int64) *Entry {
	if e == nil {
		return e
	}

	e.writeSep()
	e.writeStr(k)
	e.writeDelimar()
	e.writeInt64(v)
	return e
}
func xInt(e *Entry, k string, v int) *Entry {
	return xInt64(e, k, int64(v))
}
func xInt8(e *Entry, k string, v int8) *Entry {
	return xInt64(e, k, int64(v))
}
func xInt16(e *Entry, k string, v int16) *Entry {
	return xInt64(e, k, int64(v))
}
func xInt32(e *Entry, k string, v int32) *Entry {
	return xInt64(e, k, int64(v))
}
func BenchmarkSpecificInt(b *testing.B) {
	setup()
	b.ReportAllocs()
	e := NewEntry()
	for i := 0; i < b.N; i++ {
		e.pos = 0
		xInt(e, "int", 42)
		xInt8(e, "i8", 8)
		xInt16(e, "i16", 16)
		xInt32(e, "i32", 32)
		xInt64(e, "i64", 64)
	}
}
func BenchmarkAnyInt(b *testing.B) {
	setup()
	b.ReportAllocs()
	e := NewEntry()
	for i := 0; i < b.N; i++ {
		e.pos = 0
		anyInt(e, "int", 42)
		anyInt(e, "i8", int8(8))
		anyInt(e, "i16", int16(16))
		anyInt(e, "i32", int32(32))
		anyInt(e, "i64", int64(64))
	}
}

func BenchmarkTimeAndLevel_fmt(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(1)
		_ = fmt.Sprintf("%s %s %s:%d",
			time.Now().Format(TIME_LAYOUT),
			"\033[2;37mDBGs\033[0m",
			file, line,
		)
	}
}
func BenchmarkTimeAndLevel_onebyone(b *testing.B) {
	setup()
	b.ReportAllocs()
	e := &Entry{}
	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(1)
		t := time.Now()
		e.pos = 0
		writeAnyInt(e, t.Year())
		e.writeByte(byte('-'))
		writeAnyIntPad0(e, t.Month(), 2)
		e.writeByte(byte('-'))
		writeAnyIntPad0(e, t.Day(), 2)
		e.writeByte(byte('T'))
		writeAnyIntPad0(e, t.Hour(), 2)
		e.writeByte(byte(':'))
		writeAnyIntPad0(e, t.Minute(), 2)
		e.writeByte(byte(':'))
		writeAnyIntPad0(e, t.Second(), 2)
		e.writeByte(byte('.'))
		writeAnyIntPad0(e, t.Nanosecond()/1000000, 3)
		b := t.AppendFormat(e.a[e.pos:e.pos], "Z07:00")
		e.pos += len(b)

		e.writeSep()
		e.writeStr("\033[2;37mDBGs\033[0m")
		e.writeSep()
		e.writeStr(file)
		e.writeByte(byte(':'))
		writeAnyInt(e, line)
	}
}
func BenchmarkTimeAndLevel(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e := NewEntry()
		e.Init(LEVEL_DEBUG, 0)
	}
}

func BenchmarkTimeAndLevelWithPool(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e := gEntryPool.Get().(*Entry)
		e.Init(LEVEL_DEBUG, 0).Msg("")
	}
}
func BenchmarkLogFields_fmt(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, file, line, _ := runtime.Caller(1)
		_ = fmt.Sprintf("%s \033[2;37m%s\033[0m %s:%d %s=%s %s=%d %s=%f %s",
			time.Now().Format(TIME_LAYOUT),
			"DBG",
			file, line,
			"string", "four!",
			"int", 123,
			"float", -3.141592653589793,
			fakeMessage,
		)
	}
}
func BenchmarkLogFields(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e := gEntryPool.Get().(*Entry)
		e.Init(LEVEL_INFO, 0).
			Str("string", "four!").
			Int("int", 123).
			Float32("float", -3.141592653589793).
			Msg(fakeMessage)
	}
}

var files = []string{
	"entry.go",
	"flog/entry.go",
	"/home/user/simshi/flog/entry.go",
}

func BenchmarkFileLastIndex(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			index := strings.LastIndex(f, "/")
			if index == -1 {
				_ = f
				continue
			}
			if first := strings.LastIndex(f[:index], "/"); first != -1 {
				_ = f[first+1:]
			} else {
				_ = f
			}
		}
	}
}

// extract last module/filename from caller info, clear but slower than Index
var reFile = regexp.MustCompile(`([^/]+/)?[^/]+$`)

func BenchmarkFileRegex(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for _, f := range files {
			_ = reFile.Find([]byte(f))
		}
	}
}
