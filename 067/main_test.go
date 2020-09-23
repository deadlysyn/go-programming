package main

import "testing"

func TestMain(t *testing.T) {
	x := subtract(20, 10)
	if x != 10 {
		t.Errorf("Expected 10, got %v", x)
	}
}
