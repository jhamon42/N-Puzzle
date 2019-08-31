package main

import (
	"errors"
	"fmt"
	"reflect"
)

func remove(s []puzzle, i int) []puzzle {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func puzzInList(a puzzle, list []puzzle) (bool, int) {
	for i, b := range list {
		if reflect.DeepEqual(b.array, a.array) {
			return true, i
		}
	}
	return false, 0
}

func getSmallest(open []puzzle) (ret int) {
	for i, p := range open {
		if p.g+p.h < open[ret].g+open[ret].h {
			ret = i
		}
	}
	return
}

func aStarAlgo(env *env) (*puzzle, error) {
	var open []puzzle // = puzz.array et ses fils
	var closed []puzzle

	fmt.Println(env.actuel.array)
	open = append(open, env.actuel)

	for open != nil {

		i := getSmallest(open)
		var actual = open[i]
		remove(open, i)

		fmt.Println(actual.array)

		up := moveUp(actual)
		down := moveDown(actual)
		left := moveLeft(actual)
		right := moveRight(actual)
		fmt.Println(up.array)
		fmt.Println(down.array)
		fmt.Println(left.array)
		fmt.Println(right.array)

		return nil, errors.New("blq")

		successors := [4]puzzle{up, down, left, right}

		for _, successor := range successors {
			successor.prev = &actual
			if successor.array == nil {
				continue
			}
			if reflect.DeepEqual(successor.array, env.goal) {
				return &successor, nil
			}

			successor.g = actual.g + 1
			successor.h = env.heuri(&successor, env)

			in, i := puzzInList(successor, open)
			if in && (open[i].g+open[i].h < successor.g+successor.h) {
				continue
			}

			in, i = puzzInList(successor, closed)
			if in && (closed[i].g+closed[i].h < successor.g+successor.h) {
				continue
			}

			open = append(open, successor)
		}

		closed = append(closed, actual)
	}
	return nil, errors.New("Couldn't solve")
}
