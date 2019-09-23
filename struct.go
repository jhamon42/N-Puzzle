package main

type myMap = []int

type PriorityQueue []*Item

type Item struct {
	value 	*puzzle
	priority 	float64 // f

	index 		int
}

type puzzle struct {
	array myMap		
	moved string
	zero  int
	h     float64	// heuristique value
	g     float64	// deep value
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
