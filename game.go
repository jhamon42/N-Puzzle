package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func game(puzz *puzzle, but [][]int) {
	initTerm()
	visu(puzz)
	for checkPuzz(puzz, but) {
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

func gameAlgo(puzz *puzzle, but [][]int) {
	// for checkPuzz(puzz, but) {
	// }
	fmt.Println(diff1(puzz, but))
	fmt.Println(diff2(puzz, but))
	fmt.Println(diff3(puzz, but))
}
