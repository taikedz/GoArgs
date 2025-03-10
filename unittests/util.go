package goargsunittest

import (
	"testing"
	"time"
)

func CheckEqual[V string|int|int64|uint|float32|float64|bool|time.Duration](t *testing.T, exp_value V, got_value V) {
	if exp_value != got_value {
		t.Errorf("Got %v // Exp %v", exp_value, got_value)
	}
}

func CheckEqualArr[V string|int|float32|bool](t *testing.T, exp_value []V, got_value []V) {
	if len(exp_value) != len(got_value) {
		t.Errorf("Got %v // Exp %v", exp_value, got_value)
	}

	for i:=0; i<len(exp_value) && i<len(got_value); i++ {
		if exp_value[i] != got_value[i] {
			t.Errorf("Got %v // Exp %v", exp_value, got_value)
			return
		}
	}
}
