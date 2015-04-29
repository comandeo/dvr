package main

import (
	"github.com/nsf/termbox-go"
)

func printLine(str string, x int, y int) {
	for i := range str {
		termbox.SetCell(x+i, y, rune(str[i]), termbox.ColorDefault, termbox.ColorDefault)
	}
}

func Trace(op string, dump []string) {
	var i int
	var str string
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for i, str = range dump {
		printLine(str, 0, i)
	}
	printLine("== BOTTOM ==", 0, i+1)
	printLine("Press ENTER", 0, i+2)
	printLine("Next instruction: "+op, 40, 0)
	termbox.Flush()
loop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEnter:
				break loop
			}
		}
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	int1 := Integer{1}
	int2 := Integer{2}
	instructions := []Instruction{
		Instruction{PUSH, []Object{int1}},
		Instruction{PUSH, []Object{int2}},
		Instruction{ADD, []Object{}}}
	vm := new(VM)
	vm.Run(instructions, Trace)
}
