package main

import (
	"fmt"
)

func visu(puzz *puzzle) {
	fmt.Printf("\u001b[2J")
	for i := 0; puzz.size > i; i++ {
		fmt.Printf("\t")
		for j := 0; puzz.size > j; j++ {
			if puzz.array[i][j] < 10 && puzz.array[i][j] != 0 {
				fmt.Printf("   %d", puzz.array[i][j])
			} else if puzz.array[i][j] == 0 {
				fmt.Printf("   -")
			} else {
				fmt.Printf("  %d", puzz.array[i][j])
			}
		}
		fmt.Printf("\n\n")
	}
	fmt.Printf("\n\n")
}