package main

type puzzle struct {
	size  int
	array [][]int
	zero  *coord
}

type coord struct {
	x int
	y int
}
