package main

type myMap = []int

type puzzle struct {
	prev  *puzzle
	array myMap
	moved string
	zero  int
	h     float64
	g     float64
}

type flags struct {
	algo  string
	goal  string
	file  string
	her   string
	quiet bool
	inter bool
	visu  bool
	rand  int
}

type env struct {
	actuel   *puzzle
	goal     myMap
	size     int
	longSize int
	heuri    func(puzz *puzzle, env *env)
	algo     func(env *env) (*puzzle, error)
}
