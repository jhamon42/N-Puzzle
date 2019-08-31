package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func game(puzz *puzzle, env *env) {
	initTerm()
	visu(puzz, env)
	for checkPuzz(puzz, env.goal) {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				return
			case term.KeyArrowUp:
				*puzz = moveUp(*puzz)
			case term.KeyArrowDown:
				*puzz = moveDown(*puzz)
			case term.KeyArrowLeft:
				*puzz = moveLeft(*puzz)
			case term.KeyArrowRight:
				*puzz = moveRight(*puzz)
			}
			visu(puzz, env)
		case term.EventError:
			panic(ev.Err)
		}
	}
}

func gameAlgo(puzz *puzzle, env *env) {
	// for checkPuzz(puzz, env.goal) {
	// }
	fmt.Println(env.heuri(puzz, env))
	final, err := env.algo(env)
	checkerr(err)
	fmt.Println("Hurray")
	fmt.Println(final.array)
	fmt.Println(final.g)
}
