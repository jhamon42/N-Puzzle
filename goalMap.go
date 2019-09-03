package main

import "strings"

func goalMapBasic(env *env) myMap {
	goal := make([]int, env.longSize)
	for i := 1; env.longSize > i; i++ {
		goal[i-1] = i
	}
	goal[env.longSize-1] = 0
	return goal
}

func goalMapSnail(env *env) myMap {
	tmp := make([][]int, env.size)
	for i := 0; i < env.size; i++ {
		tmp[i] = make([]int, env.size)
	}
	i, j, ii := -1, 0, 1
	dimMax := env.size - 1
	dimMin := 0
	for dimMax > dimMin {
		for i < dimMax {
			i++
			tmp[j][i] = ii
			ii++
		}
		for j < dimMax {
			j++
			tmp[j][i] = ii
			ii++
		}
		dimMax--
		for i > dimMin {
			i--
			tmp[j][i] = ii
			ii++
		}
		dimMin++
		for j > dimMin {
			j--
			tmp[j][i] = ii
			ii++
		}
	}
	tmp[j][i] = 0
	var goal []int
	for _, slice := range tmp {
		for _, tile := range slice {
			goal = append(goal, tile)
		}
	}
	return goal
}

func goalMap(flags *flags, env *env) myMap {
	if strings.EqualFold(flags.goal, "basic") {
		return goalMapBasic(env)
	}
	return goalMapSnail(env)
}
