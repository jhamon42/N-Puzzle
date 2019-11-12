package main

// Int8Slice : is just a int8 slice for implement sort
type Int8Slice []int8

// PriorityQueue :
type PriorityQueue []*Puzzle

// PuzzleMap :
type PuzzleMap map[string]*Puzzle

// Puzzle :
type Puzzle struct {
	puzMap Int8Slice
	label  string

	h      float32 // * heuristique value// pas utile
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
	heuri    func(puzMap Int8Slice, env *Env) (h float32)
	algo     func(env *Env) error
}
