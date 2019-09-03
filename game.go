package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func gameInteractive(env *env) {
	initTerm()
	visu(env.actuel, env)
	tmp := *env.actuel
game:
	for checkPuzz(env.actuel, env.goal) {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break game
			case term.KeyArrowUp:
				tmp = moveUp(env.actuel, env)
			case term.KeyArrowDown:
				tmp = moveDown(env.actuel, env)
			case term.KeyArrowLeft:
				tmp = moveLeft(env.actuel, env)
			case term.KeyArrowRight:
				tmp = moveRight(env.actuel, env)
			}
			if tmp.array != nil {
				*env.actuel = tmp
			}
			visu(env.actuel, env)
		case term.EventError:
			panic(ev.Err)
		}
	}
	gameResume(env)
	endTerm()
}

func gameAlgo(env *env) {
	var err error
	env.actuel, err = env.algo(env)
	checkerr(err)
	fmt.Println("gg algo ~_~")
}
