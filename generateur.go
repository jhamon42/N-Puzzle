package main

import (
	"math/rand"
	"time"
)

func generator(val int, env *env) *puzzle {
	env.size = val
	env.longSize = val * val
	var tab = make([]int, env.longSize)
	puzz := &puzzle{}

	for i := 0; i < env.longSize; i++ {
		tab[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tab), func(i, j int) { tab[i], tab[j] = tab[j], tab[i] })

	puzz.array = tab
	return puzz
}
