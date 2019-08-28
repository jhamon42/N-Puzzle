package main

import (
	"fmt"
	"log"

	term "github.com/nsf/termbox-go"
)

func initTrem() {
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
