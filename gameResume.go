package main

import (
	"fmt"
)

func gameResume(env *Env) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println(env.initial.label)
	fmt.Println("")
	fmt.Println(env.state.puzArray)
	fmt.Println("")
	fmt.Println(env.goal.puzArray)
}
