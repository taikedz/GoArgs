package goargsunittest

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
)


func Test_ParseArgs_MakeVar(t *testing.T) {
	parser := goargs.NewParser("")

	name := parser.String("name", "nobody", "Their name")
	age := parser.Int("age", -1, "Their age")
	height := parser.Float("height", 0.0, "Their height")
	admit := parser.Bool("admit", false, "Whether to admit")

	args := []string{"one", "--name", "Alex", "two", "--age", "20", "--height", "1.8", "--admit", "--unknown", "--", "alpha", "beta"}
    parser.RequireFlagDefs(false)
	if err := parser.Parse(args); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, "Alex", *name)
	CheckEqual(t, 20,     *age)
	CheckEqual(t, 1.8,    *height)
	CheckEqual(t, true,   *admit)
}

func Test_ParseArgs_Specials(t *testing.T) {
	parser := goargs.NewParser("")

	verbose_lvl := parser.Count("verbose", "How verbose to be")
	parser.SetShortFlag('v', "verbose")
	choice := parser.Choices("occupation", []string{"studying", "employed", "free", "Current occupation"}, "job")

	args := []string{"--occupation", "employed", "--verbose", "-v", "--verbose", "--", "alpha", "beta"}
	if err := parser.Parse(args); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, 3, *verbose_lvl)
	CheckEqual(t, "employed", *choice)

	if err := parser.Parse([]string{"--occupation", "unknown"}); err == nil {
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
	parser.Func("queue", func(s string) error {
		queue = append(queue, s)
		return nil
	}, "help")
	parser.SetShortFlag('Q', "queue")

	name := parser.String("name", "who", "help")
	parser.SetShortFlag('N', "name")

	CheckEqual(t, 0, *verbose)
	CheckEqual(t, false, *admit)
	CheckEqual(t, "who", *name)

	if err := parser.Parse([]string{"-Q", "one", "-vav", "-N", "Rae", "--queue", "two"}); err != nil {
		t.Errorf("Failed shortflags parse: %v", err)
	}
	CheckEqual(t, 2, *verbose)
	CheckEqual(t, true, *admit)
	CheckEqual(t, "Rae", *name)
	CheckEqualArr(t, queue, []string{"one", "two"})

	if err := parser.Parse([]string{"-vavN", "Roo"}); err == nil {
		t.Errorf("Shortflags parse with combined vavN should have failed, value of name is '%s'", *name)
	}
	CheckEqual(t, 4, *verbose)
	CheckEqual(t, true, *admit)
	CheckEqual(t, "Rae", *name)
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
    parser.RequireFlagDefs(false)
	if err := parser.Parse(args); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, "Alex", name)
	CheckEqual(t, 20,     age)
	CheckEqual(t, 1.8,    height)
	CheckEqual(t, true,   admit)

	CheckEqualArr(t, []string{"one", "two", "--unknown"}, parser.Args())
	CheckEqualArr(t, []string{"alpha", "beta"},           parser.PassdownArgs)
}

func Test_ParseArgs_Fail(t *testing.T) {
	parser := goargs.NewParser("")
	var value string
	var number int
	parser.StringVar(&value, "val", "nothing", "Some value")
	parser.IntVar(&number, "num", -1, "Some numeral")

	if err := parser.Parse([]string{"front", "--val"}); err == nil {
		t.Errorf("Should have failed for --val ! Got instead: %s", value)
	}
	parser.ClearParsedData()
	
	if err := parser.Parse([]string{"--num", "NaN"}); err == nil {
		t.Errorf("Should have failed for --num ! Got instead: %d", number)
	}
	parser.ClearParsedData()
	
	if err := parser.Parse([]string{"--unknown", "what"}); err == nil {
		t.Errorf("Should have failed for --unknown ! Parser content is: %v", parser)
	}
	parser.ClearParsedData()
}


func Test_Appender(t *testing.T) {
    parser := goargs.NewParser("")

    values := parser.Appender("file", "File name")
    parser.SetShortFlag('f', "file")

    if err := parser.Parse([]string{"--file", "one", "-f", "two"}); err != nil {
        t.Errorf("Error parsing appender: %v", err)
    } else {
        CheckEqualArr(t, *values, []string{"one", "two"})
    }
}

func Test_ParseArgs_Unknowns(t *testing.T) {
    tokens := []string{"a", "--unknown", "", "b", "-x", "c"}
    parser := goargs.NewParser("")

    if err := parser.Parse(tokens); err == nil {
        t.Errorf("Should have failed parsing tokens")
    }

    parser.ClearParsedData()

    parser.RequireFlagDefs(false)
    if err := parser.Parse(tokens); err != nil {
        t.Errorf("Failed to correctly parse unknown tokens: %v", err)
    } else {
        CheckEqualArr(t, tokens, parser.Args())
    }
}
