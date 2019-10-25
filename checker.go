package main

import (
	"fmt"
	"log"
	"reflect"
	"sort"
)

func smooth(puzz *Puzzle, env *Env) {
	mapx := make([]int8, env.longSize)
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
	ic := invCount(puzz.puzMap, env)
	fmt.Printf("%f\n", ic)
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

func checkMap(puzz *Puzzle, env *Env) {
	smooth(puzz, env)
	if !checkIsOk(puzz, env) {
		puzz.puzMap = nil
		fmt.Println("bah c'est pas bon lol")
		log.Fatalf("bouge plus")
	}
}

func checkPuzz(puzz Puzzle, goal MyMap) bool {
	return !reflect.DeepEqual(goal, puzz.puzMap)
}
