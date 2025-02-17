package goargs

import (
	"testing"
	"net.taikedz.goargs/goargs"
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

func Test_ParseArgs_GetVar(t *testing.T) {
	var parser goargs.Parser

	name := parser.String("name", "nobody")
	age := parser.Int("age", -1)
	height := parser.Float("height", 0.0)
	admit := parser.Bool("admit", false)

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

func Test_ParseArgs_Good(t *testing.T) {
	var name string
	var age int
	var height float32
	var admit bool

	var parser goargs.Parser

	parser.StringVar(&name, "name", "nobody")
	parser.IntVar(&age, "age", -1)
	parser.FloatVar(&height, "height", 0.0)
	parser.BoolVar(&admit, "admit", false)

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
	var parser goargs.Parser
	var value string
	var number int
	parser.StringVar(&value, "val", "nothing")
	parser.IntVar(&number, "num", -1)

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

	if remains, _, err := goargs.UnpackExactly(args[:3], &name, &count, &ratio); err != nil || remains != nil {
		t.Errorf("Should have succeeded!")
	}

	if remains, _, err := goargs.UnpackExactly(args, &name, &count, &ratio); err == nil || remains == nil {
		t.Errorf("Should have failed due to excess tokens!")
	}

	if remains, _, err := goargs.UnpackExactly(args[:2], &name, &count, &ratio); err == nil || remains != nil {
		t.Errorf("Should have failed due to insufficient tokens!")
	}
}