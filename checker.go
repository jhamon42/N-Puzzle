package main

import (
	"fmt"
	"reflect"
	"sort"
)

func smooth(puzz *Puzzle, env *Env) {
	mapx := make([]int, env.longSize)
	copy(mapx, puzz.puzMap)
	sort.Ints(mapx)
	if mapx[0] != 0 {
		tmp := mapx[env.longSize-1]
		for i, val := range mapx[:env.longSize] {
			mapx[i] = tmp
			tmp = val
		}
	}
	for val, valux := range mapx {
		for i, valu := range puzz.puzMap {
			if valux == valu {
				puzz.puzMap[i] = val
				if val == 0 {
					puzz.zero = i
				}
			}
			i++
		}
	}
}

func checkIsOk(puzz *Puzzle, env *Env) bool {
	return true
}

func checkMap(puzz *Puzzle, env *Env) {
	smooth(puzz, env)
	if env.size < 5 && !checkIsOk(puzz, env) {
		puzz.puzMap = nil
		fmt.Println("bah c'est pas bon lol")
	}
}

func checkPuzz(puzz Puzzle, goal MyMap) bool {
	return !reflect.DeepEqual(goal, puzz.puzMap)
}
