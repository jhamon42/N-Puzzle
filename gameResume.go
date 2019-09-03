package main

import (
	"fmt"
)

func gameResume(env *env) {
	fmt.Println(env.actuel.array)
	fmt.Println("")
	fmt.Println(env.goal)
}
