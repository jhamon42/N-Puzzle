package main

import (
	"fmt"
	"log"
	"strconv"

	term "github.com/nsf/termbox-go"
)

func initTerm() {
	err := term.Init()
	checkerr(err)
}

func endTerm() {
	fmt.Println("\nEnter or Esc for exit")
	switch ev := term.PollEvent(); ev.Type {
	case term.EventKey:
		switch ev.Key {
		case term.KeyEsc:
			term.Close()
			return
		case term.KeyEnter:
			term.Close()
			return
		}
	case term.EventError:
		log.Fatal(ev.Err)
	}
}

func visu(puzz Puzzle, env *Env) {
	term.Sync()
	fmt.Println(puzz)
	fmt.Printf("\n\n")
	s := len(strconv.Itoa(env.size*env.size-1)) + 2
	iter := 0
	for i := 0; env.size > i; i++ {
		fmt.Printf("\t")
		for j := 0; env.size > j; j++ {
			if puzz.puzMap[iter] == 0 {
				fmt.Printf("%-*s", s, "-")
			} else {
				fmt.Printf("%-*d", s, puzz.puzMap[iter])
			}
			iter++
		}
		fmt.Printf("\n\n")
	}
	fmt.Printf("\n\n")
	fmt.Printf("rank : [%f]  deep : [%f]\n", puzz.f, puzz.g)
}
