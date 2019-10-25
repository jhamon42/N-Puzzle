package main

// MyMap :
type MyMap = []int

// PriorityQueue :
type PriorityQueue []*Puzzle

// PuzzleMap :
type PuzzleMap map[string]*Puzzle

// Puzzle :
type Puzzle struct {
	puzMap MyMap
	label  string

	h      float32 // * heuristique value
	f      float32 // * rank h+g
	g      float32 // * deep value
	parent *Puzzle

	open  bool
	close bool

	moved int
	zero  int

	index int
}

// Flags :
type Flags struct {
	algo  string
	goal  string
	file  string
	her   string
	quiet bool
	inter bool
	visu  bool
	rand  int
}

// Env :
type Env struct {
	initial  Puzzle
	goal     Puzzle
	size     int
	longSize int
	heuri    func(puzMap MyMap, env *Env) (h float32)
	algo     func(env *Env) error
}
