package mymath

import (
	"testing"
)

func TestPower(t *testing.T) {
	expect := 256
	actual := Power(16)
	if expect != actual {
		t.Error("Power(16) ! = 256")
	}
}

func TestPower2(t *testing.T) {
	t.Log("pass")
}
