package goargs

import (
	"testing"
	"github.com/taikedz/gocheck"
)

func Test_Unpack(t *testing.T) {
	var name string
	var count int
	var ratio float32
    var truth bool
    var lies bool
	args := []string{"Jay Smith", "15", "37.8", "true", "0", "other", "stuff"}

	if remains, err := Unpack(args, &name, &count, &ratio, &truth, &lies); err == nil {
		gocheck.Equal(t, name, "Jay Smith")
		gocheck.Equal(t, count, 15)
		gocheck.Equal(t, ratio, 37.8)
        gocheck.Equal(t, truth, true)
        gocheck.Equal(t, lies, false)
		gocheck.EqualArr(t, remains, []string{"other", "stuff"})
	} else {
		t.Errorf("Should have parsed OK, got error: %v", err)
	}

	if err := UnpackExactly(args[:3], &name, &count, &ratio); err != nil {
		t.Errorf("Should have succeeded!")
	}

	if err := UnpackExactly(args, &name, &count, &ratio); err == nil {
		t.Errorf("Should have failed due to excess tokens!")
	}

	if err := UnpackExactly(args[:2], &name, &count, &ratio); err == nil {
		t.Errorf("Should have failed due to insufficient tokens!")
	}
}
