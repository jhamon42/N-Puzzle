package main

import (
	"math/rand"
	"time"
)

func generator(val int, env *Env) {
	env.size = val
	env.longSize = val * val
	var tab = make([]int, env.longSize)

	for i := 0; i < env.longSize; i++ {
		tab[i] = int(i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(tab), func(i, j int) { tab[i], tab[j] = tab[j], tab[i] })

	env.state.puzArray = tab
	return
}
