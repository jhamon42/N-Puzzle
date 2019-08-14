package main

import "fmt"

// import "syscall"
// import "os"
// import "os/exec"

func main() {
	puzz := parceur("./test.txt")
	fmt.Println(puzz.array)
}
