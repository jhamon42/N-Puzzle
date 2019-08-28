package main

import (
	"math/rand"
	"time"
)

func generator(val int) *puzzle {
	var a []int
	var tab [][]int
	puzz := &puzzle{}
	for i := 0; i < val*val; i++ {
		a = append(a, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	for i := 0; i < val; i++ {
		if i == 0 {
			tab = append(tab, a[i*val:val])
		} else {
			tab = append(tab, a[i*val:val*(i+1)])
		}
	}
	puzz.array = tab
	puzz.size = val
	return puzz
}
