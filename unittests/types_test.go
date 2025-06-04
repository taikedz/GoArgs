package goargsunittest

import (
	"testing"
	"time"
	"github.com/taikedz/goargs/goargs"
	"github.com/taikedz/gocheck"
)

var functhings []string


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

	gocheck.Equal(t, "stringy", *my_string)

	gocheck.Equal(t, 1, *my_int)
	gocheck.Equal(t, 1, *my_int64)
	gocheck.Equal(t, 1, *my_uint)

	gocheck.Equal(t, 1.1, *my_float)
	gocheck.Equal(t, 1.1, *my_float64)

	gocheck.Equal(t, true, *my_bool)

	gocheck.Equal(t, dr, *my_duration)

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

	gocheck.Equal(t, "strung", *my_string)

	gocheck.Equal(t, -2, *my_int)
	gocheck.Equal(t, 10_000, *my_int64)
	gocheck.Equal(t, 5, *my_uint)

	gocheck.Equal(t, 6.7, *my_float)
	gocheck.Equal(t, 8008.9009, *my_float64)

	gocheck.Equal(t, false, *my_bool)

	dr5, _ := time.ParseDuration("5s")
	gocheck.Equal(t, dr5, *my_duration)

}

func Test_SpecialTypes(t *testing.T) {
	parser := goargs.NewParser("help")
	counter := parser.Count("guest", "help")
	choice := parser.Choices("dish", []string{"rice", "noodles"}, "help")
	appended := parser.Appender("toppings", "help")
	parser.SetShortFlag('t', "toppings")
	parser.Func("extra", func(value string) error {functhings = append(functhings, value); return nil}, "help")
	mode := parser.Mode("style", "chinese", map[rune]string{'c':"chinese", 'j':"japanese", 'v':"vietnamese"}, "Help")

	parser.Parse([]string{"--guest", "--dish", "noodles", "thingy", "--style", "vietnamese", "-c", "-t", "egg", "-t", "bamboo", "--extra", "spice", "--extra", "onions", "--guest","-j"})

	gocheck.EqualArr(t, []string{"thingy"}, parser.Args())
	gocheck.Equal(t, 2, *counter)
	gocheck.Equal(t, "noodles", *choice)
	gocheck.EqualArr(t, []string{"egg", "bamboo"}, *appended)
	gocheck.EqualArr(t, []string{"spice", "onions"}, functhings)
	gocheck.Equal(t, "japanese", *mode)
}

func Test_ShortFlagCombo(t *testing.T) {
	parser := goargs.NewParser("help")

	yes := parser.Bool("yes", false, "help")
	parser.SetShortFlag('y', "yes")
	slow := parser.Bool("slow", false, "help")
	parser.SetShortFlag('s', "slow")
	walk := parser.Mode("walk", "forest", map[rune]string{'f':"forest", 'b':"beach", 'h':"hill"}, "help")

	if err := parser.Parse([]string{"noop", "-yhb"}); err != nil {
		t.Errorf("Failed parse: %v", err)
	}

	gocheck.EqualArr(t, []string{"noop"}, parser.Args())
	gocheck.Equal(t, true, *yes)
	gocheck.Equal(t, false, *slow)
	gocheck.Equal(t, "beach", *walk)
}
