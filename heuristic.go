package main

func hamming(puzMap Int8Slice, env *Env) (h float32) {
	goalMap := env.goal.puzMap
	for i := 0; i < env.longSize; i++ {
		if puzMap[i] != goalMap[i] {
			h++
		}
	}
	return
}

func manhattan(puzMap Int8Slice, env *Env) (h float32) {
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

func invCount(puzMap Int8Slice, env *Env) float32 {
	h := float32(0)
	tmpMap := puzMap.Trim()

	for indexMI, keyMI := range tmpMap {
		for _, keyM := range tmpMap[indexMI:] {
			if keyMI > keyM {
				h++
			}
		}
	}
	return h
}
