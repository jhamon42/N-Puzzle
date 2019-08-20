package main

import term "github.com/nsf/termbox-go"

func game(puzz *puzzle) {
	visu(puzz)
	for checkPuzz(puzz) {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				return
			case term.KeyArrowUp:
				moveUp(puzz)
			case term.KeyArrowDown:
				moveDown(puzz)
			case term.KeyArrowLeft:
				moveLeft(puzz)
			case term.KeyArrowRight:
				moveRight(puzz)
			}
			visu(puzz)
		case term.EventError:
			panic(ev.Err)
		}
	}
}
