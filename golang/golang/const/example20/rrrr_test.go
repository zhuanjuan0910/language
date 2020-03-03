package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := add(2, 4)
	if r != 7 {
		t.Fatalf("add(2, 4) error, expect:%d, actual:%d", 7, r)
	}
	t.Logf("test add succ")
}
