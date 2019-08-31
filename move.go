package main

func moveRight(puzz puzzle) puzzle {
	if puzz.zero.y != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y-1]
		puzz.array[puzz.zero.x][puzz.zero.y-1] = 0
		puzz.zero.y--
		puzz.moved = "right"
		return puzz
	}
	return puzzle{}
}

func moveLeft(puzz puzzle) puzzle {
	if puzz.zero.y != len(puzz.array)-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y+1]
		puzz.array[puzz.zero.x][puzz.zero.y+1] = 0
		puzz.zero.y++
		puzz.moved = "left"
		return puzz
	}
	return puzzle{}
}

func moveUp(puzz puzzle) puzzle {
	if puzz.zero.x != len(puzz.array)-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x+1][puzz.zero.y]
		puzz.array[puzz.zero.x+1][puzz.zero.y] = 0
		puzz.zero.x++
		puzz.moved = "up"
		return puzz
	}
	return puzzle{}
}

func moveDown(puzz puzzle) puzzle {
	if puzz.zero.x != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x-1][puzz.zero.y]
		puzz.array[puzz.zero.x-1][puzz.zero.y] = 0
		puzz.zero.x--
		puzz.moved = "down"
		return puzz
	}
	return puzzle{}
}
