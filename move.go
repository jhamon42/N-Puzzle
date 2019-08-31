package main

import "fmt"

func moveRight(puzz puzzle) puzzle {
	if puzz.zero.y != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y-1]
		puzz.array[puzz.zero.x][puzz.zero.y-1] = 0
		puzz.zero.y--
		return puzz
	}
	return puzzle{}
}

func moveLeft(puzz puzzle) puzzle {
	if puzz.zero.y < len(puzz.array)-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x][puzz.zero.y+1]
		puzz.array[puzz.zero.x][puzz.zero.y+1] = 0
		puzz.zero.y++
		return puzz
	}
	return puzzle{}
}

func moveUp(puzz puzzle) puzzle {
	fmt.Printf("%d %d\n", puzz.zero.x, len(puzz.array)-1)
	if puzz.zero.x < len(puzz.array)-1 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x+1][puzz.zero.y]
		puzz.array[puzz.zero.x+1][puzz.zero.y] = 0
		puzz.zero.x++
		return puzz
	}
	return puzzle{}
}

func moveDown(puzz puzzle) puzzle {
	if puzz.zero.x != 0 {
		puzz.array[puzz.zero.x][puzz.zero.y] = puzz.array[puzz.zero.x-1][puzz.zero.y]
		puzz.array[puzz.zero.x-1][puzz.zero.y] = 0
		puzz.zero.x--
		return puzz
	}
	fmt.Println("ici")
	return puzzle{}
}
