package main

func hamming(puzz *puzzle, env *env) int {
	diff := 0
	for x := 0; x < env.size; x++ {
		for y := 0; y < env.size; y++ {
			if puzz.array[x][y] != env.goal[x][y] {
				diff++
			}
		}
	}
	return diff
}

func diff2(puzz *puzzle, env *env) int {
	diff := 0
	for x := 0; x < env.size; x++ {
		for y := 0; y < env.size; y++ {
			i := 0
		test:
			for xx := x; xx < env.size; xx++ {
				for yy := y; yy < env.size; yy++ {
					if puzz.array[xx][yy] == env.goal[x][y] {
						break test
					}
					i++
				}
			}
			diff += i
		}
	}
	return diff
}

func manhattan(puzz *puzzle, env *env) int {
	diff := 0
	for x := 0; x < env.size; x++ {
		for y := 0; y < env.size; y++ {
			i := 0
		test:
			for xx := x; xx < env.size; xx++ {
				for yy := y; yy < env.size; yy++ {
					if puzz.array[xx][yy] == env.goal[x][y] {
						break test
					} else {
						i++
					}
				}
			}
			diff += (i / env.size) + (i % env.size)
		}
	}
	return diff
}
