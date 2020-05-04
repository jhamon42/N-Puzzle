package main

import (
	"strconv"
	"strings"
)

// func (puz *Puzzle) giveMeKey() string {
// 	tmpStr := []byte("")
// 	for _, key := range puz.puzMap {
// 		tmpStr = strconv.AppendInt(tmpStr, int64(key), 16)
// 	}
// 	puz.label = string(tmpStr)
// 	return puz.label
// }

// func (puz *Puzzle) puzRank() float32 {
// 	puz.f = puz.g + puz.h
// 	return puz.f
// }

func (state *PuzzleEnv) heuri(env *Env) float32 {
	return env.heuri(state.puzArray, env)
}

func (state *PuzzleEnv) backUp(puz *Puzzle) {
	state.puz = puz
	state.puzArray = []int{}
	for i, j := range strings.Split(state.puz.label, " ") {
		tmp, _ := strconv.Atoi(j)
		state.puzArray = append(state.puzArray, tmp)
		if tmp == 0 {
			state.zero = i
		}
	}
}

func (puz *Puzzle) newborn(env *Env) {
	puz.parent = env.state.puz
	puz.g = env.state.puz.g + 1
	puz.f = env.state.heuri(env) + puz.g
	puz.label = arrayToString(env.state.puzArray, " ")
}

// func (puz *Puzzle) findNeighbor(env *Env) []*Puzzle {
// 	truc := make([]*Puzzle, 4)

// 	truc[0] = move(puz, env, -(env.size))
// 	truc[1] = move(puz, env, env.size)
// 	truc[2] = move(puz, env, 1)
// 	truc[3] = move(puz, env, -1)

// 	return truc
// }
