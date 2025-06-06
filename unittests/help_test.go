package goargsunittest

import (
	"testing"
    "strings"
	"github.com/taikedz/goargs/goargs"
	"github.com/taikedz/gocheck"
)


func Test_helpstr(t *testing.T) {
	parser := goargs.NewParser("Whack-a-mole")

	parser.String("gopher", "gaffer", "Wee rat")
    parser.SetShortFlag('g', "gopher")
	parser.Bool("whack", false, "Slam it?")
	parser.Int("times", 1, "How many?")
    parser.Float("freq", 0.2, "WPS (whacks per second)")

	helptext := parser.SPrintHelp()
	expect := strings.Join([]string{
        "Whack-a-mole",
        "",
        "  --gopher STRING",
        "  -g STRING",
        "    default: gaffer",
        "    Wee rat",
        "  --whack",
        "    default: false",
        "    Slam it?",
        "  --times INT",
        "    default: 1",
        "    How many?",
        "  --freq FLOAT",
        // Float 32 will come out like this without knowing contextual precision
        "    default: 0.200000",
        "    WPS (whacks per second)",
        }, "\n")
	if helptext != expect {
		t.Errorf("Mismatched help strings. Got:\n<<%s>>\nInstead of:\n<<%s>>", helptext, expect)
	}
}

func noop(value string) error {
    return nil
}

func Test_helpst_special(t *testing.T) {
    parser := goargs.NewParser("Whack-a-mole")

    parser.Count("hard", "How hard?")
    parser.SetShortFlag('h', "hard")
    parser.Choices("tool", []string{"rolling-pin", "hammer"}, "What to use?")
    parser.Appender("noise", "Acceptable squeaks")
    parser.Func("custom", noop, "Something happens")
    parser.Mode("damage", "blunt", map[rune]string{'b':"blunt", 's':"sharp"}, "Damage type")

    helptext := parser.SPrintHelp()
    expect := strings.Join([]string{
        "Whack-a-mole",
        "",
        "  --hard",
        "  -h",
        "    (each appearance is counted)",
        "    How hard?",
        "  --tool STRING",
        "    default: rolling-pin",
        "    choices: rolling-pin, hammer",
        "    What to use?",
        "  --noise STRING",
        "    (can be specified multiple times)",
        "    Acceptable squeaks",
        "  --custom STRING",
        "    Something happens",
        "  --damage STRING",
        "    default: blunt",
        "    Damage type",
        "      -b : blunt",
        "      -s : sharp",
    }, "\n")
	if helptext != expect {
		t.Errorf("Mismatched help strings. Got:\n<<%s>>\nInstead of:\n<<%s>>", helptext, expect)
	}
}

func Test_findhelp(t *testing.T) {
	gocheck.Equal(t, 1, goargs.FindHelpFlag([]string{"a", "-h", "--help"}))
	gocheck.Equal(t, 0, goargs.FindHelpFlag([]string{"-h", "next", "--help"}))
	gocheck.Equal(t, 2, goargs.FindHelpFlag([]string{"-he", "next", "--help"}))
	gocheck.Equal(t, -1, goargs.FindHelpFlag([]string{"a", "--", "-h", "--help"}))
}
