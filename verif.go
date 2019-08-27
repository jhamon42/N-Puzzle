package main

import (
	"reflect"
)

func checkPuzz(puzz *puzzle, but [][]int) bool {
	return !reflect.DeepEqual(but, puzz.array)
}
