package log

import (
	"errors"

	"github.com/simshi/flog"
)

func ExampleDebug() {
	flog.SetLevel(flog.LEVEL_DEBUG)
	Debug().Int("key", -42).IntPad0("code", 0, 8).UintPad0("value", 321, 4).Msg("Hello World!")
	Info().Str("host", "example.com").Uint("port", 9527).Msg("Hello World!")
	Warn().Hex("raw", 0x42ab).Hex("neg", -2).Msgf("Greeting %s!", "World")
	const pi = 3.141592653589793238462643383279
	Warn().Float32("pi", pi).Float64("pi64", pi).Msgf("%d,%.2f Hello %s!", 42, 42.0, "World")
	Error().Bool("msg sent", true).Err(errors.New("timeout")).Msg("Hello World!")

	// Output:
}

func ExampleFatal() {
	Fatal().Msg("Critical error, Byebye World!")
	// Output:
}
