package main

import (
	"math/rand"
	"time"
)

func generator(val int, env *Env) Puzzle {
	env.size = val
	env.longSize = val * val
	var tab = make([]int8, env.longSize)
	puzz := Puzzle{}

	for i := 0; i < env.longSize; i++ {
		tab[i] = int8(i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tab), func(i, j int) { tab[i], tab[j] = tab[j], tab[i] })

	puzz.puzMap = tab
	return puzz
}
