package main

import (
	"fmt"
	"sort"
)

func smooth(puzz *puzzle, env *env) {
	mapi := make(map[int]int)
	ii := 1
	var mapx []int
	for i := 0; env.size > i; i++ {
		for j := 0; env.size > j; j++ {
			mapi[puzz.array[i][j]] = 0
			mapx = append(mapx, puzz.array[i][j])
		}
	}
	sort.Ints(mapx)
	for i := range mapx {
		mapi[mapx[i]] = ii
		ii++
	}
	mapi[mapx[ii-2]] = 0
	for i := 0; env.size > i; i++ {
		for j := 0; env.size > j; j++ {
			puzz.array[i][j] = mapi[puzz.array[i][j]]
			if puzz.array[i][j] == 0 {
				puzz.zero.x = i
				puzz.zero.y = j
			}
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
