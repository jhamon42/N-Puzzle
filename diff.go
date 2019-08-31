package main

func diff1(puzz *puzzle, env *env) int {
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
			if i != 0 {
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

func diff3(puzz *puzzle, env *env) int {
	diff := 0
	for x := 0; x < env.size; x++ {
		for y := 0; y < env.size; y++ {
			i := 0
		test:
			for xx := x; xx < env.size; xx++ {
				for yy := y; yy < env.size; yy++ {
					if puzz.array[xx][yy] == env.goal[x][y] {
						break test
					} else if puzz.array[xx][yy] > env.goal[x][y] {
						i++
					}
				}
			}
			diff += i
		}
	}
	return diff
}
