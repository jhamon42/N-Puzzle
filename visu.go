package main

import (
	"fmt"
	"strconv"

	term "github.com/nsf/termbox-go"
)

func visu(puzz *puzzle, env *env) {
	term.Sync()
	fmt.Printf("\n\n")
	s := len(strconv.Itoa(env.size*env.size-1)) + 2

	for i := 0; env.size > i; i++ {
		fmt.Printf("\t")
		for j := 0; env.size > j; j++ {
			if puzz.array[i][j] == 0 {
				fmt.Printf("%-*s", s, "-")
			} else {
				fmt.Printf("%-*d", s, puzz.array[i][j])
			}
		}
		fmt.Printf("\n\n")
	}
	fmt.Printf("\n\n")
}
