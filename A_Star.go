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
func removetest(s []int, i int) []int {
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

	test := []int{1, 2, 3, 4}
	test = removetest(test, 0)
	test = removetest(test, 0)
	fmt.Println(test)

	fmt.Println(env.actuel.array)
	open = append(open, env.actuel)

	for len(open) != 0 {

		i := getSmallest(open)
		var actual = open[i]

		open = remove(open, i)
		closed = append(closed, actual)

		up := moveUp(actual)
		down := moveDown(actual)
		left := moveLeft(actual)
		right := moveRight(actual)

		successors := [4]puzzle{up, down, left, right}

		for _, successor := range successors {
			successor.prev = &actual
			if successor.array == nil {
				continue
			}
			if reflect.DeepEqual(successor.array, env.goal) {
				return &successor, nil
			}

			in, _ := puzzInList(successor, closed)
			if in {
				continue
			}

			successor.g = actual.g + 1
			successor.h = env.heuri(&successor, env)

			in, j := puzzInList(successor, open)
			if !in {
				open = append(open, successor)
			} else {
				if successor.h < open[j].h {
					open = remove(open, j)
					open = append(open, actual)
				}
			}

			// in, i := puzzInList(successor, open)
			// if in && (open[i].g+open[i].h < successor.g+successor.h) {
			// 	continue
			// }
			//
			// in, i = puzzInList(successor, closed)
			// if in && (closed[i].g+closed[i].h < successor.g+successor.h) {
			// 	continue
			// } else {
			// 	open = append(open, successor)
			// }

		}

	}
	return nil, errors.New("Couldn't solve")
}
