package main

type puzzle struct {
	array [][]int
	size  int
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
