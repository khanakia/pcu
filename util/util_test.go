package util

import (
	"testing"
)

func Test1(t *testing.T) {
	result := TagNameConvert("v6.22.0")
	if result != "^6.22" {
		t.Errorf("returned = %s; want ^6.22", result)
	}
}

func Test2(t *testing.T) {
	result := TagNameConvert("1.6.x")
	if result != "^1.6" {
		t.Errorf("returned = %s; want ^1.6", result)
	}
}

func Test3(t *testing.T) {
	result := TagNameConvert("3.7.1")
	if result != "^3.7" {
		t.Errorf("returned = %s; want ^3.7", result)
	}
}
