package main

func hamming(puzMap []int, env *Env) (h float32) {
	goalMap := env.goal.puzArray
	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			h++
		}
	}
	return h
}

func manhattan(puzMap []int, env *Env) (h float32) {
	goalMap := env.goal.puzArray
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
	return h
}

func invCount(puzMap []int, env *Env) float32 {
	h := float32(0)
	tmpMap := trimInt(puzMap)
	tmpGMap := trimInt(env.goal.puzArray)
	for indexMI, keyMI := range tmpMap {
		for _, keyM := range tmpMap[indexMI:] {
			if indexInt(tmpGMap, keyMI) > indexInt(tmpGMap, keyM) {
				h++
			}
		}
	}
	return h
}
