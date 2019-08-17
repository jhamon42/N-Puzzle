package main

func moveRight(puzz *puzzle) {
	if puzz.zero.y != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y-1]
		puzz.array[puzz.zero.x][puzz.zero.y-1] = 0
		puzz.zero.y--
	}
}

func moveLeft(puzz *puzzle) {
	if puzz.zero.y != puzz.size-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y+1]
		puzz.array[puzz.zero.x][puzz.zero.y+1] = 0
		puzz.zero.y++
	}
}

func moveUp(puzz *puzzle) {
	if puzz.zero.x != puzz.size-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x+1][puzz.zero.y]
		puzz.array[puzz.zero.x+1][puzz.zero.y] = 0
		puzz.zero.x++
	}
}

func moveDown(puzz *puzzle) {
	if puzz.zero.x != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x-1][puzz.zero.y]
		puzz.array[puzz.zero.x-1][puzz.zero.y] = 0
		puzz.zero.x--
	}
}
