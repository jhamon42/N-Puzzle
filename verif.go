package main

import (
	"reflect"
)

func checkPuzz(puzz *puzzle) bool {
	return !reflect.DeepEqual(puzz.but, puzz.array)
}
