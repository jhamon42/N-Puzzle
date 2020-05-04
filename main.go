package main

import (
	"fmt"
	"strings"
	"time"
)

func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

func trimInt(a []int) []int {
	var trim []int
	for _, key := range a {
		if key != 0 {
			trim = append(trim, key)
		}
	}
	return trim
}

func indexInt(a []int, comp int) int {
	for index, key := range a {
		if key == comp {
			return index
		}
	}
	return -1
}

func main() {
	init := time.Now()
	flags := parceFlags()
	env := parceEnv(flags)
	parcePuzz(flags, env)
	fmt.Println(env)
	if flags.inter {
		gameInteractive(env)
	} else {
		// if flags.visu {
		// 	initTerm()
		// 	defer endTerm()
		// }
		// gameAlgo(env)
		// gameResume(env)
	}

	fmt.Println(time.Since(init))
}
