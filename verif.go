package main

import (
	"reflect"
)

func checkPuzz(puzz *puzzle, goal myMap) bool {
	return !reflect.DeepEqual(goal, puzz.array)
}
