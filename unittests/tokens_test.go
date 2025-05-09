package goargsunittest

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
	"github.com/taikedz/gocheck"
)

func Test_tokenize(t *testing.T) {
    var fore []string
    var aft []string

    fore, aft = goargs.SplitTokensBefore("--", []string{"a", "b c", "--", "d", "e f"})
    gocheck.EqualArr(t, []string{"a", "b c"}, fore)
    gocheck.EqualArr(t, []string{"d", "e f"}, aft)

    fore, aft = goargs.SplitTokensBefore("--", []string{"--", "a" , "--", "e f"})
    gocheck.EqualArr(t, []string{}, fore)
    gocheck.EqualArr(t, []string{"a", "--", "e f"}, aft)

    fore, aft = goargs.SplitTokensBefore("--", []string{"n", "p x", "--"})
    gocheck.EqualArr(t, []string{"n", "p x"}, fore)
    gocheck.EqualArr(t, []string{}, aft)
}
