package main

const (
	HALT = iota
	PUSH = iota
	POP  = iota
	ADD  = iota
)

type VM struct {
	stack Stack
}

func (vm *VM) Run(instructions []Instruction, trace func(string, []string)) {
	for _, instruction := range instructions {
		switch instruction.opcode {
		case PUSH:
			trace("PUSH", vm.stack.Dump())
			vm.stack.Push(instruction.operands[0])
			break
		case POP:
			trace("POP", vm.stack.Dump())
			vm.stack.Pop()
			break
		case ADD:
			trace("ADD", vm.stack.Dump())
			op2, _ := vm.stack.Pop()
			op1, _ := vm.stack.Pop()
			res, _ := op1.Call("add", op2)
			vm.stack.Push(res)
			break
		}
	}
	trace("NO INSTRUCTION", vm.stack.Dump())
}
