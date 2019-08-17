package main

import "fmt"

/*
	[1 2 3]
	[4 5 6]
	[7 8 0]

	[1 2 3]
	[8 0 4]
	[7 6 5]
*/

func chkNormal(puzz *puzzle) bool {
	return true
}

func chkCargo(puzz *puzzle) bool {
	return true
}

func checkPuzz(puzz *puzzle) bool {
	tst := chkNormal(puzz)
	fmt.Println(chkNormal(puzz))
	fmt.Println(chkCargo(puzz))
	return tst
}
