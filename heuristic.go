package main

func hamming(puzz *puzzle, env *env) {
	for i := 0; i < env.longSize; i++ {
		if puzz.array[i] != env.goal[i] {
			puzz.h++
		}
	}
}

func manhattan(puzz *puzzle, env *env) {
	for i := 0; i < env.longSize; i++ {
		if puzz.array[i] != env.goal[i] {
			tmp := 0
			for j := i; j < env.longSize; j++ {
				tmp++
			}
			puzz.h += (float64)(tmp/env.size) + (float64)(tmp%env.size)
		}
	}
}
