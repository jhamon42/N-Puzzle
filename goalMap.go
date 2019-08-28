package main

import "strings"

func goalMapBasic(puzz *puzzle) [][]int {
	ii := 1
	var goal [][]int
	for i := 0; puzz.size > i; i++ {
		pool := []int{}
		for j := 0; puzz.size > j; j++ {
			pool = append(pool, ii)
			ii++
		}
		goal = append(goal, pool)
	}
	goal[puzz.size-1][puzz.size-1] = 0
	return goal
}

func goalMapSnail(puzz *puzzle) [][]int {
	goal := make([][]int, puzz.size)
	for i := 0; i < puzz.size; i++ {
		goal[i] = make([]int, puzz.size)
	}

	i, j, ii := -1, 0, 1
	dimMax := puzz.size - 1
	dimMin := 0

	for dimMax > dimMin {
		for i < dimMax {
			i++
			goal[j][i] = ii
			ii++
		}
		for j < dimMax {
			j++
			goal[j][i] = ii
			ii++
		}
		dimMax--
		for i > dimMin {
			i--
			goal[j][i] = ii
			ii++
		}
		dimMin++
		for j > dimMin {
			j--
			goal[j][i] = ii
			ii++
		}
	}
	goal[j][i] = 0
	return goal
}

func goalMap(puzz *puzzle, typ string) [][]int {
	goal := make([][]int, puzz.size)
	if strings.EqualFold(typ, "basic") {
		goal = goalMapBasic(puzz)
	} else {
		goal = goalMapSnail(puzz)
	}
	return goal
}
