package goargsunittest

import (
	"testing"
	"github.com/taikedz/goargs/goargs"
)

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