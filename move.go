package main

func moveRight(puzz puzzle) puzzle {
	new := puzzle(puzz)
	new.array = make([][]int, len(puzz.array))
	for j := 0; j < len(puzz.array); j++ {
		new.array[j] = append(new.array[j], puzz.array[j]...)
	}
	if new.zero.y != 0 {
		new.array[new.zero.x][new.zero.y] = new.array[new.zero.x][new.zero.y-1]
		new.array[new.zero.x][new.zero.y-1] = 0
		new.zero.y--
		// fmt.Print("Moved right : ")
		// fmt.Println(new.array)
		return new
	}
	return puzzle{}
}

func moveLeft(puzz puzzle) puzzle {
	new := puzzle(puzz)
	new.array = make([][]int, len(puzz.array))
	for j := 0; j < len(puzz.array); j++ {
		new.array[j] = append(new.array[j], puzz.array[j]...)
	}
	if new.zero.y < len(new.array)-1 {
		new.array[new.zero.x][new.zero.y] = new.array[new.zero.x][new.zero.y+1]
		new.array[new.zero.x][new.zero.y+1] = 0
		new.zero.y++
		// fmt.Print("Moved left : ")
		// fmt.Println(new.array)
		return new
	}
	return puzzle{}
}

func moveUp(puzz puzzle) puzzle {
	new := puzzle(puzz)
	new.array = make([][]int, len(puzz.array))
	for j := 0; j < len(puzz.array); j++ {
		new.array[j] = append(new.array[j], puzz.array[j]...)
	}
	if new.zero.x < len(new.array)-1 {
		new.array[new.zero.x][new.zero.y] = new.array[new.zero.x+1][new.zero.y]
		new.array[new.zero.x+1][new.zero.y] = 0
		new.zero.x++
		new.moved = "up"
		// fmt.Print("Moved up : ")
		// fmt.Println(new.array)
		// fmt.Println(new.moved)
		return new
	}
	return puzzle{}
}

func moveDown(puzz puzzle) puzzle {
	new := puzzle(puzz)
	new.array = make([][]int, len(puzz.array))
	for j := 0; j < len(puzz.array); j++ {
		new.array[j] = append(new.array[j], puzz.array[j]...)
	}
	if new.zero.x != 0 {
		new.array[new.zero.x][new.zero.y] = new.array[new.zero.x-1][new.zero.y]
		new.array[new.zero.x-1][new.zero.y] = 0
		new.zero.x--
		// fmt.Print("Moved down : ")
		// fmt.Println(new.array)
		return new
	}
	return puzzle{}
}
