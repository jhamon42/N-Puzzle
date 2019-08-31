package main

import "strings"

func goalMapBasic(puzz *puzzle, env *env) myMap {
	ii := 1
	var goal [][]int
	for i := 0; env.size > i; i++ {
		pool := []int{}
		for j := 0; env.size > j; j++ {
			pool = append(pool, ii)
			ii++
		}
		goal = append(goal, pool)
	}
	goal[env.size-1][env.size-1] = 0
	return goal
}

func goalMapSnail(puzz *puzzle, env *env) myMap {
	goal := make([][]int, env.size)
	for i := 0; i < env.size; i++ {
		goal[i] = make([]int, env.size)
	}

	i, j, ii := -1, 0, 1
	dimMax := env.size - 1
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

func goalMap(puzz *puzzle, flags *flags, env *env) myMap {
	goal := make([][]int, env.size)
	if strings.EqualFold(flags.goal, "basic") {
		goal = goalMapBasic(puzz, env)
	} else {
		goal = goalMapSnail(puzz, env)
	}
	return goal
}
