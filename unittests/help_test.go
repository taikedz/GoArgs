package goargsunittest

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
)


func Test_helpstr(t *testing.T) {
	parser := goargs.NewParser("Do lots")

	parser.String("gopher", "gaffer", "Wee rat")
	parser.Bool("whack", false, "Slam it?")

	helptext := parser.SPrintHelp()
	expect := "Do lots\n\n  --gopher STRING\n    default: gaffer\n    Wee rat\n  --whack\n    default: false\n    Slam it?"
	if helptext != expect {
		t.Errorf("Mismatched help strings. Got:\n<<%s>>\nInstead of:\n<<%s>>", helptext, expect)
	}
}

func Test_findhelp(t *testing.T) {
	CheckEqual(t, 1, goargs.FindHelpFlag([]string{"a", "-h", "--help"}))
	CheckEqual(t, 0, goargs.FindHelpFlag([]string{"-h", "next", "--help"}))
	CheckEqual(t, 2, goargs.FindHelpFlag([]string{"-he", "next", "--help"}))
	CheckEqual(t, -1, goargs.FindHelpFlag([]string{"a", "--", "-h", "--help"}))
}