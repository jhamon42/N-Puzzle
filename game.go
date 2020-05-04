package main

import (
	term "github.com/nsf/termbox-go"
)

func gameInteractive(env *Env) {
	initTerm()
	visu(env)
	state := &env.state
game:
	for checkPuzz(state.puz.label, env.goal.puz.label) {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break game
			case term.KeyArrowUp:
				state.puz = move(state, env, -(env.size))
			case term.KeyArrowDown:
				state.puz = move(state, env, env.size)
			case term.KeyArrowLeft:
				state.puz = move(state, env, -1)
			case term.KeyArrowRight:
				state.puz = move(state, env, 1)
			case term.KeyDelete:
				if state.puz.parent != nil {
					state.backUp(state.puz.parent)
				}
			}
			if state.puz.label == "" {
				if state.puz.parent != nil {
					state.backUp(state.puz.parent)
				}
			}
			visu(env)
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
	// fmt.Println("gg algo ~_~")
}
