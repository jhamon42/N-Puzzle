package main

import (
	"strconv"
)

func (puz *Puzzle) giveMeKey() string {
	tmpStr := []byte("")
	for _, key := range puz.puzMap {
		tmpStr = strconv.AppendInt(tmpStr, int64(key), 16)
	}
	puz.label = string(tmpStr)
	return puz.label
}

func (puz *Puzzle) puzRank() float32 {
	puz.f = puz.g + puz.h
	return puz.f
}

func (puz *Puzzle) puzPrevCost(env *Env) float32 {
	puz.h = env.heuri(puz.puzMap, env)
	return puz.h
}

func (puz *Puzzle) newborn(parent *Puzzle) {
	puz.parent = parent
	puz.puzMap = make(MyMap, len(parent.puzMap))
	copy(puz.puzMap, parent.puzMap)
	puz.g = parent.g + 1
}

func (puz *Puzzle) findNeighbor(env *Env) []*Puzzle {
	truc := make([]*Puzzle, 4)

	truc[0] = move(puz, env, -(env.size))
	truc[1] = move(puz, env, env.size)
	truc[2] = move(puz, env, 1)
	truc[3] = move(puz, env, -1)

	return truc
}
