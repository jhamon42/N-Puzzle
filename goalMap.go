package main

import (
	"strings"
)

func goalMapBasic(env *Env) []int {
	goal := make([]int, env.longSize)
	for i := 1; env.longSize > i; i++ {
		goal[i-1] = int(i)
	}
	goal[env.longSize-1] = 0
	return goal
}

func goalMapSnail(env *Env) []int {
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
			goal = append(goal, int(tile))
		}
	}
	return goal
}

func goalMap(flags *Flags, env *Env) {
	var gMap []int
	if strings.EqualFold(flags.goal, "basic") {
		gMap = goalMapBasic(env)
	} else {
		gMap = goalMapSnail(env)
	}
	puz := Puzzle{}
	puz.label = arrayToString(gMap, " ")
	env.goal.puzArray = gMap
	env.goal.puz = &puz
	return
}
