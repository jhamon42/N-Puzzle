package main

/*
   loi:
   - Total number of states ever selected in the "opened"
   set (complexity in time)
   - Maximum number of states ever
   represented in memory at the same timeduring the
   search (complexity in size)
   - Number of moves required
   to transition from the initial state to the final
   state,according to the searchs
   - The ordered sequence
   of states that make up the solution, according to
   thesearch
   - The puzzle may be unsolvable, in which
   case you have to inform the user andexit
*/

func main() {
	env := &env{}
	flags := parceFlags()
	puzz := parce(flags, env)
	env.goal = goalMap(puzz, flags, env)
	if flags.inter {
		game(puzz, env)
		gameResume(puzz, env)
		endTerm()
	} else {
		if flags.visu {
			initTerm()
			defer endTerm()
		}
		gameAlgo(puzz, env)
		gameResume(puzz, env)
	}
}
