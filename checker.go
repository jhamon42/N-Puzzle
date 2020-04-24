package main

// import (
// 	"fmt"
// 	"log"
// 	"reflect"
// )

// func smooth(puzz *Puzzle, env *Env) {
// 	mapx := make(Int8Slice, env.longSize)
// 	copy(mapx, puzz.puzMap)
// 	mapx.Sort()
// 	if mapx[0] != 0 {
// 		tmp := mapx[env.longSize-1]
// 		for i, val := range mapx[:env.longSize] {
// 			mapx[i] = tmp
// 			tmp = val
// 		}
// 	}
// 	for val, valux := range mapx {
// 		for i, valu := range puzz.puzMap {
// 			if valux == valu {
// 				puzz.puzMap[i] = int8(val)
// 				// if val == 0 {
// 				// 	puzz.zero = i
// 				// }
// 			}
// 			i++
// 		}
// 	}
// }

// func checkIsOk(puzz *Puzzle, env *Env) bool {
// 	ic := invCount(puzz.puzMap, env)
// 	if env.size%2 == 1 {
// 		if int(ic)%2 == 0 {
// 			return true
// 		}
// 		return false
// 	}
// 	if !(int(ic)%2 == 0 && (env.size-(puzz.zero%env.size))%2 == 0) &&
// 		(int(ic)%2 == 0 || (env.size-(puzz.zero%env.size))%2 == 0) {
// 		return true
// 	}
// 	return false
// }

// func checkMap(puzz *Puzzle, env *Env) {
// 	smooth(puzz, env)
// 	if !checkIsOk(puzz, env) {
// 		// puzz.puzMap = nil
// 		fmt.Println("bouge plus, tu n'est pas bon")
// 		// time.Sleep(2 * time.Second)
// 		log.Fatalln("la flem")
// 	}
// }

// func checkPuzz(puzz Puzzle, goal Int8Slice) bool {
// 	return !reflect.DeepEqual(goal, puzz.puzMap)
// }
