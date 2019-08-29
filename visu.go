package main

import (
	"fmt"
	"strconv"
	term "github.com/nsf/termbox-go"
)

func visu(puzz *puzzle) {
	term.Sync()
	fmt.Printf("\n\n")
	s :=  len(strconv.Itoa(puzz.size*puzz.size-1))+2

	for i := 0; puzz.size > i; i++ {
		fmt.Printf("\t")
		for j := 0; puzz.size > j; j++ {
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
