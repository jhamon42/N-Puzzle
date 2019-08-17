package main

import (
	"fmt"

	term "github.com/nsf/termbox-go"
)

/*
   pour les flag :
   - hersestique
   - algo (a*)
   - file ou generateur // si pas file, generateur
   - visu
   - quiet

   loi:
   - Total number of states ever selected in the "opened"
   set (complexity in time)
   - Maximum number of states ever
   represented in memory at the same timeduring the
   search (complexity in size)
   - Number of moves required
   to transition from the initial state to the final
   state,according to the search
   - The ordered sequence
   of states that make up the solution, according to
   thesearch
   - The puzzle may be unsolvable, in which
   case you have to inform the user andexit

*/
func reset() {
	term.Sync() // cosmestic purpose
}

func main() {

	puzz := parceur("./test.txt")
	checkMap(puzz)
	err := term.Init()
	if err != nil {
		panic(err)
	}

	defer term.Close()
	visu(puzz)
	fmt.Println(generator())
keyPressListenerLoop:
	for {
		switch ev := term.PollEvent(); ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyEsc:
				break keyPressListenerLoop
			case term.KeyArrowUp:
				reset()
				fmt.Println("Arrow Up pressed")
				moveUp(puzz)
			case term.KeyArrowDown:
				reset()
				fmt.Println("Arrow Down pressed")
				moveDown(puzz)
			case term.KeyArrowLeft:
				reset()
				fmt.Println("Arrow Left pressed")
				moveLeft(puzz)
			case term.KeyArrowRight:
				reset()
				fmt.Println("Arrow Right pressed")
				moveRight(puzz)

			default:
				// we only want to read a single character or one key pressed event
				reset()
				fmt.Println("ASCII : ", ev.Ch)

			}
			visu(puzz)
		case term.EventError:
			panic(ev.Err)
		}
	}
}
