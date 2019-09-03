package main

import (
	"errors"
	"fmt"
	"reflect"
)

func remove(s []puzzle, i int) (news []puzzle) {
	news = append(news, s[:i]...)
	news = append(news, s[i+1:]...)
	return
}

func insert(p puzzle, s []puzzle, i int) (news []puzzle) {
	news = append(news, s[:i]...)
	news = append(news, p)
	news = append(news, s[i:]...)
	return
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
	lenOpen := 1

	fmt.Println(env.actuel.array)
	open = append(open, *env.actuel)
	for lenOpen != 0 {
		var actual = open[lenOpen-1]

		open = open[:lenOpen-1]
		lenOpen--
		closed = append(closed, actual)
		up := moveUp(&actual, env)
		down := moveDown(&actual, env)
		left := moveLeft(&actual, env)
		right := moveRight(&actual, env)
		successors := [4]puzzle{up, down, left, right}

		for _, successor := range successors {
			successor.prev = &actual
			if successor.array == nil {
				continue
			}

			if reflect.DeepEqual(successor.array, env.goal) {
				return &successor, nil
			}

			in, i := puzzInList(successor, closed)

			env.heuri(&successor, env)
			successor.g = actual.g + 0.5
			if in && closed[i].g+closed[i].h > successor.g+successor.h {
				for j, puzz := range open {
					if puzz.g+puzz.h > successor.g+successor.h {
						i = j
						break
					}
				}
				open = insert(successor, open, i)
				lenOpen++
				continue
			}
			i = lenOpen - 1
			if i == -1 {
				open = insert(successor, open, 0)
				lenOpen++
			}

		test:
			for i > 0 {
				if successor.g+successor.h <= open[i].h+open[i].g {
					open = insert(successor, open, i)
					lenOpen++
					i--
					for i >= 0 {

						if reflect.DeepEqual(successor.array, open[i].array) {
							remove(open, i)
							break test
						}
						i--
					}
				} else if reflect.DeepEqual(successor.array, open[i].array) {
					break
				}
				i--
			}
			if i == 0 {
				open = insert(successor, open, i)
				lenOpen++
			}
			fmt.Printf("--  %f  --  %d  -- %d\n", successor.g+successor.h, lenOpen, lenOpen+len(closed))
		}
	}
	fmt.Println(closed[len(closed)-1].array)
	return nil, errors.New("Couldn't solve")
}
