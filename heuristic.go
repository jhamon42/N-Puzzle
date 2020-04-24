package main

func hamming(puzMap []int, env *Env) (h float32) {
	// goalMap := env.goal.label
	// for i := 0; i < env.longSize; i++ {
	// 	if puzMap[i] != strconv.Atoi(goalMap[i]) {
	// 		h++
	// 	}
	// }
	return h
}

func manhattan(puzMap []int, env *Env) (h float32) {
	// goalMap := env.goal.label

	// for i := 0; i < env.longSize; i++ {
	// 	if puzMap[i] != goalMap[i] {
	// 		tmp := 0
	// 		for j := i; j < env.longSize && puzMap[i] != goalMap[j]; j++ {
	// 			tmp++
	// 			if puzMap[i] == goalMap[j] {
	// 				h += ((float32)(tmp/env.size) * (float32)(tmp/env.size)) + ((float32)(tmp%env.size) * (float32)(tmp%env.size))
	// 			}
	// 		}
	// 	}
	// }
	return h
}

func invCount(puzMap []int, env *Env) float32 {
	h := float32(0)
	// tmpMap := puzMap.Trim()
	// tmpGMap := env.goal.puzMap.Trim()

	// for indexMI, keyMI := range tmpMap {
	// 	for _, keyM := range tmpMap[indexMI:] {
	// 		if tmpGMap.Index(keyMI) < tmpGMap.Index(keyM) {
	// 			h++
	// 		}
	// 	}
	// }

	return h
}
