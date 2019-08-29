package main

func diff1(puzz *puzzle, but [][]int) int {
	diff := 0
	for x := 0; x < puzz.size; x++ {
		for y := 0; y < puzz.size; y++ {
			i := 0
		test:
			for xx := x; xx < puzz.size; xx++ {
				for yy := y; yy < puzz.size; yy++ {
					if puzz.array[xx][yy] == but[x][y] {
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

func diff2(puzz *puzzle, but [][]int) int {
	diff := 0
	for x := 0; x < puzz.size; x++ {
		for y := 0; y < puzz.size; y++ {
			i := 0
		test:
			for xx := x; xx < puzz.size; xx++ {
				for yy := y; yy < puzz.size; yy++ {
					if puzz.array[xx][yy] == but[x][y] {
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

func diff3(puzz *puzzle, but [][]int) int {
	diff := 0
	for x := 0; x < puzz.size; x++ {
		for y := 0; y < puzz.size; y++ {
			i := 0
		test:
			for xx := x; xx < puzz.size; xx++ {
				for yy := y; yy < puzz.size; yy++ {
					if puzz.array[xx][yy] == but[x][y] {
						break test
					} else if puzz.array[xx][yy] > but[x][y] {
						i++
					}
				}
			}
			diff += i
		}
	}
	return diff
}
