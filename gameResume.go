package main

import (
	"fmt"
)

func gameResume(puzz *puzzle, env *env) {
	fmt.Println(puzz.array)
	fmt.Println("")
	fmt.Println(env.goal)
}
