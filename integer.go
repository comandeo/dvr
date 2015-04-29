package main

import "errors"

type Integer struct {
	value int
}

func (object Integer) Add(another Integer) Integer {
	return Integer{object.value + another.value}
}

func (object Integer) RespondTo(method string) bool {
	return method == "add"
}

func (object Integer) Call(method string, args ...Object) (result Object, err error) {
	switch method {
	case "add":
		result = object.Add(args[0].(Integer))
	default:
		err = errors.New("No method error")
	}
	return
}

func (object Integer) Type() string {
	return "Integer"
}
