package main

import "testing"

func TestDoAnotherThing(t *testing.T) {
	expected := "-32.340000 - is the 666c6f61743332"
	result := doAnotherThing(-32.34)

	if result != expected {
		t.Errorf("Test failed, expected %s but got %s", expected, result)
	}
}

func TestUseMultiplyAndReturn(t *testing.T) {
	result := UseMultiplyAndReturn()

	if result != 16 {
		t.Error()
	}
}
