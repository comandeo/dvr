package main

import "testing"

type SimpleObject struct {
	value int
}

func (so SimpleObject) RespondTo(method string) bool {
	return true
}

func (so SimpleObject) Call(method string, args ...Object) (result Object, err error) {
	return
}

func (so SimpleObject) Type() string {
	return "SimpleObject"
}

func TestStack(t *testing.T) {
	stack := new(Stack)
	val, err := stack.Pop()
	if val != nil && err == nil {
		t.Error("Pop from empty stack must return nil and error")
	}
	stack.Push(SimpleObject{1})
	stack.Push(SimpleObject{2})
	stack.Push(SimpleObject{3})
	for i := 3; i > 0; i-- {
		obj, err := stack.Pop()
		if err == nil && obj.(SimpleObject).value != i {
			t.Errorf("Pop: %d expected, but got %d", i, obj.(SimpleObject).value)
		}
	}
}
