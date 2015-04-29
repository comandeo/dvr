package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value Object
	next  *Element
}

func (stack *Stack) Pop() (value Object, err error) {
	if stack.size == 0 {
		return nil, errors.New("Stack is empty")
	}
	value, stack.top = stack.top.value, stack.top.next
	stack.size--
	return
}

func (stack *Stack) Peek() (value Object, err error) {
	if stack.size == 0 {
		return nil, errors.New("Stack is empty")
	}
	value = stack.top.value
	return
}

func (stack *Stack) Push(value Object) {
	stack.top = &Element{value, stack.top}
	stack.size++
}

func (stack *Stack) Dump() []string {
	dump := make([]string, 0)
	elem := stack.top
	for elem != nil {
		switch elem.value.Type() {
		case "Integer":
			dump = append(dump, fmt.Sprintf("Integer (%d)", elem.value.(Integer).value))
		}
		elem = elem.next
	}
	return dump
}
