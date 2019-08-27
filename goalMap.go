package main

func goalMapNormal(puzz *puzzle) [][]int {
	ii := 1
	var but [][]int
	for i := 0; puzz.size > i; i++ {
		bac := []int{}
		for j := 0; puzz.size > j; j++ {
			bac = append(bac, ii)
			ii++
		}
		but = append(but, bac)
	}
	but[puzz.size-1][puzz.size-1] = 0
	return but
}

func goalMapSnail(puzz *puzzle) [][]int {
	but := make([][]int, puzz.size)
	for i := 0; i < puzz.size; i++ {
		but[i] = make([]int, puzz.size)
	}

	i, j, ii := -1, 0, 1
	dimMax := puzz.size - 1
	dimMin := 0

	for dimMax > dimMin {
		for i < dimMax {
			i++
			but[j][i] = ii
			ii++
		}
		for j < dimMax {
			j++
			but[j][i] = ii
			ii++
		}
		dimMax--
		for i > dimMin {
			i--
			but[j][i] = ii
			ii++
		}
		dimMin++
		for j > dimMin {
			j--
			but[j][i] = ii
			ii++
		}
	}
	but[j][i] = 0
	return but
}
