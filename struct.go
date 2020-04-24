package main

// PriorityQueue :
type PriorityQueue []*Puzzle

// PuzzleMap :
type PuzzleMap map[string]*Puzzle

// Puzzle :
type Puzzle struct {
	label string // "1 2 5 6 4 8 . 3 7 "
	index int

	parent *Puzzle

	f    float32 // * rank h+g
	open bool
}

// PuzzleEnv :
type PuzzleEnv struct {
	puz *Puzzle

	h    float32
	g    int
	zero int
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
	initial  *Puzzle
	goal     *Puzzle
	state    PuzzleEnv
	size     int
	longSize int
	heuri    func(puzMap []int, env *Env) (h float32)
	algo     func(env *Env) error
}
