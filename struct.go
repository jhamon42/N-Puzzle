package main

type puzzle struct {
	size  int
	array [][]int
	but   [][]int
	zero  *coord
}

type coord struct {
	x int
	y int
}
