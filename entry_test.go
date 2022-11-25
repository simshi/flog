package flog

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"runtime"
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

	e = e.Int("key1", 32)
	checkStrContains(t, e, " key1=32")

	e = e.Int("key2", -233)
	checkStrContains(t, e, " key2=-233")

	e.Msg("int")
	checkStrContains(t, e, " int\n")
}
func TestEntryIntPad0(t *testing.T) {
	setup()

	e := NewEntry()
	e.Init(LEVEL_DEBUG, 0)

	e = e.IntPad0("k1", 32, 1)
	checkStrContains(t, e, " k1=32")

	e = e.IntPad0("k2", 32, 2)
	checkStrContains(t, e, " k1=32 k2=32")

	e = e.IntPad0("k3", 32, 3)
	checkStrContains(t, e, " k2=32 k3=032")

	e = e.IntPad0("neg", -32, 5)
	checkStrContains(t, e, " k3=032 neg=-0032")

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
func BenchmarkInt(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e := NewEntry()
		e.Int("key", i)
	}
}
func BenchmarkTimeAndLevel_fmt(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s \033[2;37m%s\033[0m",
			time.Now().Format(TIME_LAYOUT),
			"DBG",
		)
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
		e := ep.Get().(*Entry)
		e.Init(LEVEL_DEBUG, 0).Msg("")
	}
}
func BenchmarkLogFields(b *testing.B) {
	setup()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		e := ep.Get().(*Entry)
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
