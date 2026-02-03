package goargs

import (
	"testing"
	"github.com/taikedz/gocheck"
)

var functhings []string


func Test_SpecialTypes(t *testing.T) {
	parser := NewParser("help")
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
