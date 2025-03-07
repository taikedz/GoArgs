package goargsunittest

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
)

func Test_tokenize(t *testing.T) {
    var fore []string
    var aft []string

    fore, aft = goargs.SplitTokensBefore("--", []string{"a", "b c", "--", "d", "e f"})
    CheckEqualArr(t, []string{"a", "b c"}, fore)
    CheckEqualArr(t, []string{"d", "e f"}, aft)

    fore, aft = goargs.SplitTokensBefore("--", []string{"--", "a" , "--", "e f"})
    CheckEqualArr(t, []string{}, fore)
    CheckEqualArr(t, []string{"a", "--", "e f"}, aft)

    fore, aft = goargs.SplitTokensBefore("--", []string{"n", "p x", "--"})
    CheckEqualArr(t, []string{"n", "p x"}, fore)
    CheckEqualArr(t, []string{}, aft)
}
