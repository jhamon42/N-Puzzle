package main

import (
	"fmt"
	"time"
)

func main() {
	init := time.Now()
	flags := parceFlags()
	env := parceEnv(flags)
	parcePuzz(flags, env)
	// var t Flags
	if flags.inter {
		gameInteractive(env)
	} else {
		if flags.visu {
			initTerm()
			defer endTerm()
		}
		gameAlgo(env)
		gameResume(env)
	}

	fmt.Println(time.Since(init))
}
