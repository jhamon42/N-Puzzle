package main

type myMap = [][]int

type puzzle struct {
	array myMap
	h     int
	g     int
	zero  *coord
}

type coord struct {
	x int
	y int
}

type flags struct {
	algo  string
	her   string
	file  string
	rand  int
	quiet bool
	visu  bool
	goal  string
	inter bool
}

type env struct {
	actuel puzzle
	goal   myMap
	size   int
}
