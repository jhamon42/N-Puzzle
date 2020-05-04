package main

func move(puzz *PuzzleEnv, env *Env, move int) *Puzzle {
	if puzz.zero+move < env.longSize &&
		puzz.zero+move >= 0 &&
		!(puzz.zero%env.size == env.size-1 && move == 1) &&
		!(puzz.zero%env.size == 0 && move == -1) {

		puzz.zero = puzz.zero + move
		puzz.puzArray[puzz.zero-move] = puzz.puzArray[puzz.zero]
		puzz.puzArray[puzz.zero] = 0

		newPuzz := Puzzle{}
		newPuzz.newborn(env)
		return &newPuzz
	}
	return &Puzzle{parent: puzz.puz}
}
