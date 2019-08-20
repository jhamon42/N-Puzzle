package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func gameResume(puzz *puzzle) {
	fmt.Println(puzz.array)
	fmt.Println("Enter or Esc for exit")
	switch ev := term.PollEvent(); ev.Type {
	case term.EventKey:
		switch ev.Key {
		case term.KeyEsc:
			return
		case term.KeyEnter:
			return
		}
		visu(puzz)
	case term.EventError:
		panic(ev.Err)
	}
}
