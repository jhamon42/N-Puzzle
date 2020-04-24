package main

import (
	"fmt"
	"strings"
	"time"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func main() {
	init := time.Now()
	flags := parceFlags()
	env := parceEnv(flags)
	parcePuzz(flags, env)
	// var t Flags
	// if flags.inter {
	// 	gameInteractive(env)
	// } else {
	// 	if flags.visu {
	// 		initTerm()
	// 		defer endTerm()
	// 	}
	// 	gameAlgo(env)
	// 	gameResume(env)
	// }

	fmt.Println(time.Since(init))
}
