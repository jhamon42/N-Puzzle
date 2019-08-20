package main

import (
	"fmt"
	"sort"
)

func smooth(puzz *puzzle) {
	mapi := make(map[int]int)
	ii := 1
	var mapx []int
	for i := 0; puzz.size > i; i++ {
		for j := 0; puzz.size > j; j++ {
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
	for i := 0; puzz.size > i; i++ {
		for j := 0; puzz.size > j; j++ {
			puzz.array[i][j] = mapi[puzz.array[i][j]]
			if puzz.array[i][j] == 0 {
				puzz.zero.x = i
				puzz.zero.y = j
			}
		}
	}
}

func butMap(puzz *puzzle) {
	ii := 1
	for i := 0; puzz.size > i; i++ {
		bac := []int{}
		for j := 0; puzz.size > j; j++ {
			bac = append(bac, ii)
			ii++
		}
		puzz.but = append(puzz.but, bac)
	}
	puzz.but[puzz.size-1][puzz.size-1] = 0
}

func checkIsOk(puzz *puzzle) bool {
	return true
}

func checkMap(puzz *puzzle) {
	smooth(puzz)
	butMap(puzz)
	if checkIsOk(puzz) {
		return
	}
	puzz.array = nil
	fmt.Println("bah c'est pas bon lol")
}
