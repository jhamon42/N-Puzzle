package main

func moveRight(puzz *puzzle, env *env) puzzle {
	new := puzzle{}
	new.array = make(myMap, env.longSize)
	copy(new.array, puzz.array)
	new.zero = puzz.zero
	if new.zero%env.size > 0 {
		new.array[new.zero] = new.array[new.zero-1]
		new.array[new.zero-1] = 0
		new.zero--
		new.moved = "right"
		return new
	}
	return puzzle{}
}

func moveLeft(puzz *puzzle, env *env) puzzle {
	new := puzzle{}
	new.array = make(myMap, env.longSize)
	copy(new.array, puzz.array)
	new.zero = puzz.zero
	if new.zero%env.size < env.size-1 {
		new.array[new.zero] = new.array[new.zero+1]
		new.array[new.zero+1] = 0
		new.zero++
		new.moved = "left"
		return new
	}
	return puzzle{}
}

func moveUp(puzz *puzzle, env *env) puzzle {
	new := puzzle{}
	new.array = make(myMap, env.longSize)
	copy(new.array, puzz.array)
	new.zero = puzz.zero
	if new.zero < env.longSize-env.size {
		new.array[new.zero] = new.array[new.zero+env.size]
		new.array[new.zero+env.size] = 0
		new.zero += env.size
		new.moved = "up"
		return new
	}
	return puzzle{}
}

func moveDown(puzz *puzzle, env *env) puzzle {
	new := puzzle{}
	new.array = make(myMap, env.longSize)
	copy(new.array, puzz.array)
	new.zero = puzz.zero
	if new.zero > env.size-1 {
		new.array[new.zero] = new.array[new.zero-env.size]
		new.array[new.zero-env.size] = 0
		new.zero -= env.size
		new.moved = "down"
		return new
	}
	return puzzle{}
}
