package main

import (
	"fmt"
	"reflect"
	"sort"
)

func smooth(puzz *puzzle, env *env) {
	mapx := make([]int, env.longSize)
	copy(mapx, puzz.array)
	sort.Ints(mapx)
	if mapx[0] != 0 {
		tmp := mapx[env.longSize-1]
		for i, val := range mapx[:env.longSize] {
			mapx[i] = tmp
			tmp = val
		}
	}
	for val, valux := range mapx {
		for i, valu := range puzz.array {
			if valux == valu {
				puzz.array[i] = val
				if val == 0 {
					puzz.zero = i
				}
			}
			i++
		}
	}
}

func checkIsOk(puzz *puzzle, env *env) bool {
	return true
}

func checkMap(puzz *puzzle, env *env) {
	smooth(puzz, env)
	if env.size < 5 && !checkIsOk(puzz, env) {
		puzz.array = nil
		fmt.Println("bah c'est pas bon lol")
	}
}

func checkPuzz(puzz *puzzle, goal myMap) bool {
	return !reflect.DeepEqual(goal, puzz.array)
}
