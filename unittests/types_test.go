package goargsunittest

import (
	"testing"
	"time"
	"github.com/taikedz/goargs/goargs"
)


func Test_BasicTypes(t *testing.T) {
	p := goargs.NewParser("help")

	my_string  := p.String("astring", "stringy", "help")

	my_int  := p.Int("aint", 1, "help")
	my_int64  := p.Int64("aint64", 1, "help")
	my_uint  := p.Uint("auint", 1, "help")

	my_float  := p.Float("afloat", 1.1, "help")
	my_float64  := p.Float64("afloat64", 1.1, "help")

	my_bool  := p.Bool("abool", true, "help")

	dr, _ := time.ParseDuration("20s")
	my_duration  := p.Duration("aduration", dr, "help")

	CheckEqual(t, "stringy", *my_string)

	CheckEqual(t, 1, *my_int)
	CheckEqual(t, 1, *my_int64)
	CheckEqual(t, 1, *my_uint)

	CheckEqual(t, 1.1, *my_float)
	CheckEqual(t, 1.1, *my_float64)

	CheckEqual(t, true, *my_bool)

	CheckEqual(t, dr, *my_duration)

	args := []string {
		"--astring", "strung",

		"--aint", "-2",
		"--aint64", "10000",
		"--auint", "5",

		"--afloat", "6.7",
		"--afloat64", "8008.9009",

		"--abool",
		"--aduration", "5s",
	}
	if err := p.Parse(args); err != nil {
		t.Errorf("Types parse fail: %v", err)
		return
	}

	CheckEqual(t, "strung", *my_string)

	CheckEqual(t, -2, *my_int)
	CheckEqual(t, 10_000, *my_int64)
	CheckEqual(t, 5, *my_uint)

	CheckEqual(t, 6.7, *my_float)
	CheckEqual(t, 8008.9009, *my_float64)

	CheckEqual(t, false, *my_bool)

	dr5, _ := time.ParseDuration("5s")
	CheckEqual(t, dr5, *my_duration)

}
