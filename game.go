package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

func gameInteractive(env *Env) {
	initTerm()
	visu(env.initial, env)
	tmp := &env.initial
game:
	for checkPuzz(*tmp, env.goal.puzMap) {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break game
			case term.KeyArrowUp:
				tmp = move(tmp, env, -(env.size))
			case term.KeyArrowDown:
				tmp = move(tmp, env, env.size)
			case term.KeyArrowLeft:
				tmp = move(tmp, env, -1)
			case term.KeyArrowRight:
				tmp = move(tmp, env, 1)
			case term.KeyDelete:
				if tmp.parent != nil {
					tmp = tmp.parent
				}
			}
			if tmp.puzMap == nil {
				tmp = tmp.parent
			}
			visu(*tmp, env)
		case term.EventError:
			panic(ev.Err)
		}
	}
	gameResume(env)
	endTerm()
}

func gameAlgo(env *Env) {
	err := env.algo(env)
	checkerr(err)
	fmt.Println("gg algo ~_~")
}
