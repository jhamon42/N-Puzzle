package main

func aStarAlgo(puzz *puzzle, but [][]int) {
	var open []*puzzle // = puzz.array et ses fils
	var close []*puzzle

	succes := false
	for open != nil && !succes {
		// state = un des fils de puzz
		// open = append() les fils de state
		if checkPuzz(puzz, but) { // c'est plus state
			succes = true
		} else {
			open = nil //
			close = append(close, puzz)
		}
	}

}
