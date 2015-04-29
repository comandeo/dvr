package main

type Object interface {
	RespondTo(method string) bool
	Call(method string, args ...Object) (result Object, err error)
	Type() string
}
