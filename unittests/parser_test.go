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
	if err := parser.Parse(args, true); err != nil {
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
	if err := parser.Parse(args, false); err != nil {
		t.Errorf("Failed parse: %v", err)
		return
	}

	CheckEqual(t, 3, *verbose_lvl)
	CheckEqual(t, "employed", *choice)

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

	CheckEqual(t, 0, *verbose)
	CheckEqual(t, false, *admit)
	CheckEqual(t, "who", *name)

	if err := parser.Parse([]string{"-Q", "one", "-vav", "-N", "Rae", "--queue", "two"}, false); err != nil {
		t.Errorf("Failed shortflags parse: %v", err)
	}
	CheckEqual(t, 2, *verbose)
	CheckEqual(t, true, *admit)
	CheckEqual(t, "Rae", *name)
	CheckEqualArr(t, queue, []string{"one", "two"})

	if err := parser.Parse([]string{"-vavN", "Roo"}, false); err == nil {
		t.Errorf("Shortflags parse with combined vavN should have failed, value of name is '%s'", *name)
	}
	CheckEqual(t, 4, *verbose)
	CheckEqual(t, true, *admit)
	CheckEqual(t, "Rae", *name)

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