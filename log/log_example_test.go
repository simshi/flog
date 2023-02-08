package log

import (
	"errors"
)

func ExampleDebug() {
	SetLevel("debug")
	Debug().Int("key", 42).Int("neg", -23).IntPad0("code", -1, 4).UintPad0("value", 321, 4).Msg("Hello World!")
	Info().Str("host", "example.com").Uint("port", 9527).Any("any_array", []string{"1", "2"}).Msg("Hello World!")
	Warn().Hex("hex", 0x42ab).Hex("neg", -2).Msgf("Hello %s!", "World")
	const pi = 3.141592653589793238462643383279
	Warn().Float32("pi", pi).Float64("pi64", pi).Msgf("%d,%.2f Hello %s!", 42, 42.0, "World")
	Error().Bool("msg sent", true).Err(errors.New("timeout")).Msg("Hello World!")
	// Fatal().IntPad0("msg", 64, 4).Err(errors.New("too many msg")).Msg("Hello World!")

	// shorthand
	Debugf("Hello %s!", "World")
	Infof("Hello %s!", "World")
	Warnf("Hello %s!", "World")
	Errorf("Hello %s!", "World")
	// Fatalf("Hello %s!", "World")

	// Output:
}
