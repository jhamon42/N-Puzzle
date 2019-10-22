package main

func hamming(puzMap MyMap, env *Env) (h float64) {
	goalMap := env.goal.puzMap
	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			h++
		}
	}
	return
}

func manhattan(puzMap MyMap, env *Env) (h float64) {
	goalMap := env.goal.puzMap
	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			tmp := 0
			for j := i; j < env.longSize; j++ {
				tmp++
			}
			h += (float64)(tmp/env.size) + (float64)(tmp%env.size)
		}
	}
	return
}
