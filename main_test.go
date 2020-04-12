package main

import "testing"

func TestSubtract(t *testing.T) {
	res := Subtract(2, 4)

	if res != 0 {
		t.Error()
	}
}
