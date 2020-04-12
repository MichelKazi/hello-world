package main

import "testing"

func TestSubtract(t *testing.T) {
	res := Subtract(2, 4)

	if res != -2 {
		t.Error()
	}
}
