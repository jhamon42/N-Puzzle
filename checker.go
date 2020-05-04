package main

import (
	"fmt"
	"log"
	"sort"
)

func smooth(puzz []int, env *Env) {
	mapx := make([]int, env.longSize)
	copy(mapx, puzz)
	sort.Ints(mapx)
	if mapx[0] != 0 {
		tmp := mapx[env.longSize-1]
		for i, val := range mapx[:env.longSize] {
			mapx[i] = tmp
			tmp = val
		}
	}
	for val, valux := range mapx {
		for i, valu := range puzz {
			if valux == valu {
				puzz[i] = int(val)
				if val == 0 {
					env.state.zero = i
				}
			}
			i++
		}
	}
}

func checkIsOk(puzz PuzzleEnv, env *Env) bool {
	ic := invCount(puzz.puzArray, env)
	if env.size%2 == 1 {
		if int(ic)%2 == 0 {
			return true
		}
		return false
	}
	if !(int(ic)%2 == 0 && (env.size-(puzz.zero%env.size))%2 == 0) &&
		(int(ic)%2 == 0 || (env.size-(puzz.zero%env.size))%2 == 0) {
		return true
	}
	return false
}

func checkMap(env *Env) {
	smooth(env.state.puzArray, env)
	if !checkIsOk(env.state, env) {
		fmt.Println("bouge plus, tu n'est pas bon")
		log.Fatalln("la flem")
	}
}

func checkPuzz(puzz string, goal string) bool {
	return puzz != goal
}
