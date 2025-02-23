package goargs

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
)

func CheckEqual[V string|int|float32|bool](t *testing.T, exp_value V, got_value V) {
	if exp_value != got_value {
		t.Errorf("Got %v // Exp %v", exp_value, got_value)
	}
}

func CheckEqualArr[V string|int|float32|bool](t *testing.T, exp_value []V, got_value []V) {
	if len(exp_value) != len(got_value) {
		t.Errorf("Got %v // Exp %v", exp_value, got_value)
	}

	for i:=0; i<len(exp_value); i++ {
		if exp_value[i] != got_value[i] {
			t.Errorf("Got %v // Exp %v", exp_value, got_value)
			return
		}
	}
}

// ===

func Test_ParseArgs_MakeVar(t *testing.T) {
	parser := goargs.NewParser("")

	name := parser.String("name", "nobody", "Their name")
	age := parser.Int("age", -1, "Their age")
	height := parser.Float("height", 0.0, "Their height")
	admit := parser.Bool("admit", false, "Whether to admit")

	args := []string{"one", "--name", "Alex", "two", "--age", "20", "--height", "1.8", "--admit", "--unknown", "--", "alpha", "beta"}
	if err := parser.Parse(args, true); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, *name, "Alex")
	CheckEqual(t, *age, 20)
	CheckEqual(t, *height, 1.8)
	CheckEqual(t, *admit, true)
}

func Test_ParseArgs_Specials(t *testing.T) {
	parser := goargs.NewParser("")

	verbose_lvl := parser.Count("verbose", "How verbose to be")
	parser.SetShortFlag('v', "verbose")
	choice := parser.Choices("occupation", []string{"studying", "employed", "free", "Current occupation"}, "job")

	args := []string{"--occupation", "employed", "--verbose", "-v", "--verbose", "--", "alpha", "beta"}
	if err := parser.Parse(args, false); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, *verbose_lvl, 3)
	CheckEqual(t, *choice, "employed")

	if err := parser.Parse([]string{"--occupation", "unknown"}, false); err == nil {
		t.Errorf("'occupation unknown' Should have failed, but var assigned: '%s'", *choice)
	}
}

func Test_ParseArgs_Shortflags(t *testing.T) {
	parser := goargs.NewParser("")

	verbose := parser.Count("verbose", "help")
	parser.SetShortFlag('v', "verbose")
	admit := parser.Bool("admit", false, "help")
	parser.SetShortFlag('a', "admit")

	var queue []string
	parser.Func("queue", "help", func(s string) error {
		queue = append(queue, s)
		return nil
	})
	parser.SetShortFlag('Q', "queue")

	name := parser.String("name", "who", "help")
	parser.SetShortFlag('N', "name")

	CheckEqual(t, *verbose, 0)
	CheckEqual(t, *admit, false)
	CheckEqual(t, *name, "who")

	if err := parser.Parse([]string{"-Q", "one", "-vav", "-N", "Rae", "--queue", "two"}, false); err != nil {
		t.Errorf("Failed shortflags parse: %v", err)
	}
	CheckEqual(t, *verbose, 2)
	CheckEqual(t, *admit, true)
	CheckEqual(t, *name, "Rae")
	CheckEqualArr(t, queue, []string{"one", "two"})

	if err := parser.Parse([]string{"-vavN", "Roo"}, false); err == nil {
		t.Errorf("Shortflags parse with combined vavN should have failed, value of name is '%s'", *name)
	}
	CheckEqual(t, *verbose, 4)
	CheckEqual(t, *admit, true)
	CheckEqual(t, *name, "Rae")

	if err := parser.SetShortFlag('u', "unknown"); err == nil {
		t.Errorf("Setting unknown short flag should have failed")
	}
}

func Test_ParseArgs_Good(t *testing.T) {
	var name string
	var age int
	var height float32
	var admit bool

	parser := goargs.NewParser("")

	parser.StringVar(&name, "name", "nobody", "Their name")
	parser.IntVar(&age, "age", -1, "Their age")
	parser.FloatVar(&height, "height", 0.0, "Their height")
	parser.BoolVar(&admit, "admit", false, "Whether to admit")

	args := []string{"one", "--name", "Alex", "two", "--age", "20", "--height", "1.8", "--admit", "--unknown", "--", "alpha", "beta"}
	if err := parser.Parse(args, true); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, name, "Alex")
	CheckEqual(t, age, 20)
	CheckEqual(t, height, 1.8)
	CheckEqual(t, admit, true)

	CheckEqualArr(t, parser.Args(), []string{"one", "two", "--unknown"})
	CheckEqualArr(t, parser.PassdownArgs, []string{"alpha", "beta"})
}

func Test_ParseArgs_Fail(t *testing.T) {
	parser := goargs.NewParser("")
	var value string
	var number int
	parser.StringVar(&value, "val", "nothing", "Some value")
	parser.IntVar(&number, "num", -1, "Some numeral")

	if err := parser.Parse([]string{"front", "--val"}, false); err == nil {
		t.Errorf("Should have failed for --val ! Got instead: %s", value)
	}
	parser.ClearParsedData()
	
	if err := parser.Parse([]string{"--num", "NaN"}, false); err == nil {
		t.Errorf("Should have failed for --num ! Got instead: %d", number)
	}
	parser.ClearParsedData()
	
	if err := parser.Parse([]string{"--unknown", "what"}, false); err == nil {
		t.Errorf("Should have failed for --unknown ! Parser content is: %v", parser)
	}
	parser.ClearParsedData()
}

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

func Test_Unpack(t *testing.T) {
	var name string
	var count int
	var ratio float32
	args := []string{"Jay Smith", "15", "37.8", "other", "stuff"}

	if remains, err := goargs.Unpack(args, &name, &count, &ratio); err == nil {
		CheckEqual(t, name, "Jay Smith")
		CheckEqual(t, count, 15)
		CheckEqual(t, ratio, 37.8)
		CheckEqualArr(t, remains, []string{"other", "stuff"})
	} else {
		t.Errorf("Should have parsed OK, got error: %v", err)
	}

	if err := goargs.UnpackExactly(args[:3], &name, &count, &ratio); err != nil {
		t.Errorf("Should have succeeded!")
	}

	if err := goargs.UnpackExactly(args, &name, &count, &ratio); err == nil {
		t.Errorf("Should have failed due to excess tokens!")
	}

	if err := goargs.UnpackExactly(args[:2], &name, &count, &ratio); err == nil {
		t.Errorf("Should have failed due to insufficient tokens!")
	}
}