package main

import "testing"

func TestAdd(t *testing.T) {
	int1 := Integer{1}
	int2 := Integer{3}
	if res := int1.Add(int2); res.value != 4 {
		t.Errorf("Add: %d expected but got %d", 4, res.value)
	}
}

func TestCall(t *testing.T) {
	int1 := Integer{1}
	int2 := Integer{3}
	if res, _ := int1.Call("add", int2); res.(Integer).value != 4 {
		t.Errorf("Call: %d was expected but got %d", 4, res.(Integer).value)
	}
	if _, err := int1.Call("some", int2); err == nil {
		t.Errorf("Call: error was expected but didn't get one")
	}
}
