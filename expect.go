package goutils

import (
	"testing"
)

func expect(t *testing.T, src interface{}, dst interface{}) {
	if src != dst {
		t.Errorf("Expected %v (type %[1]T) - Got %v (type %[2]T)", dst, src)
	}
}
