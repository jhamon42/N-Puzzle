package main

import "fmt"

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

func main() {
	puzz := parceur("./test.txt")
	checkPuzz(puzz)
	fmt.Println(puzz.array)
}
