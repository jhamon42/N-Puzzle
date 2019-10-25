package main

import "fmt"

func hamming(puzMap MyMap, env *Env) (h float32) {
	goalMap := env.goal.puzMap
	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			h++
		}
	}
	return
}

func manhattan(puzMap MyMap, env *Env) (h float32) {
	goalMap := env.goal.puzMap

	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			tmp := 0
			for j := i; j < env.longSize && puzMap[i] != goalMap[j]; j++ {
				tmp++
				if puzMap[i] == goalMap[j] {
					h += ((float32)(tmp/env.size) * (float32)(tmp/env.size)) + ((float32)(tmp%env.size) * (float32)(tmp%env.size))
				}
			}
		}
	}
	return
}

func invCount(puzMap MyMap, env *Env) (h float32) {
	fmt.Println(puzMap)
	for i := 0; i < env.longSize-1; i++ {
		for j := i + 1; j < env.longSize && (puzMap[i] == env.goal.puzMap[j]); j++ {
			if puzMap[j] != 0 && puzMap[i] != 0 && puzMap[i] > puzMap[j] {
				h++
			}
		}
	}
	return
}
