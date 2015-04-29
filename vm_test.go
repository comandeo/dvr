package main

import "testing"

func TestRun(t *testing.T) {
	int1 := Integer{1}
	int2 := Integer{2}
	instructions := []Instruction{
		Instruction{PUSH, []Object{int1}},
		Instruction{PUSH, []Object{int2}},
		Instruction{ADD, []Object{}}}
	vm := new(VM)
	vm.Run(instructions, func(string, []string) {})
	if res, err := vm.stack.Pop(); err != nil || res.(Integer).value != 3 {
		t.Errorf("Pop: 3 expected, but got %d", res.(Integer).value)
	}
}
