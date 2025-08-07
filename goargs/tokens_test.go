package goargs

import (
	"testing"

	"github.com/taikedz/gocheck"
)

func Test_tokenize(t *testing.T) {
	var fore []string
	var aft []string

	fore, aft = splitTokensBefore("--", []string{"a", "b c", "--", "d", "e f"})
	gocheck.EqualArr(t, []string{"a", "b c"}, fore)
	gocheck.EqualArr(t, []string{"d", "e f"}, aft)

	fore, aft = splitTokensBefore("--", []string{"--", "a", "--", "e f"})
	gocheck.EqualArr(t, []string{}, fore)
	gocheck.EqualArr(t, []string{"a", "--", "e f"}, aft)

	fore, aft = splitTokensBefore("--", []string{"n", "p x", "--"})
	gocheck.EqualArr(t, []string{"n", "p x"}, fore)
	gocheck.EqualArr(t, []string{}, aft)
}
