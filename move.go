package main

func move(puzz *Puzzle, env *Env, move int) *Puzzle {
	newPuzz := Puzzle{}
	newPuzz.newborn(puzz)
	newPuzz.zero = puzz.zero + move
	if newPuzz.zero < env.longSize &&
		newPuzz.zero >= 0 &&
		!(puzz.zero%env.size == env.size-1 && move == 1) &&
		!(puzz.zero%env.size == 0 && move == -1) {

		newPuzz.puzMap[puzz.zero] = newPuzz.puzMap[newPuzz.zero]
		newPuzz.puzMap[newPuzz.zero] = 0

		newPuzz.parent = puzz
		newPuzz.puzPrevCost(env)
		newPuzz.giveMeKey()
		newPuzz.puzRank()
		newPuzz.moved = move

		return
	}
	return &Puzzle{parent: puzz}
}

/*
	mettre en comme le parent pour test la memoir
	int > int8 (voir pour le []int8)
*/
