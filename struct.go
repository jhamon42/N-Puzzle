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
	algo  *string //aStar / idaStar* / ...
	her   *string //P* / CL / L
	file  *string //any / ""*
	rand  *int    //3-10 / 4*
	quiet *bool   //true / false*
	visu  *bool   //true / false*
	but   *string //snail / normal*
}
