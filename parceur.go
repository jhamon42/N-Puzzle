package main

import (
	"bufio"
	"fmt"
    "strings"
	"os"
)

func parceur(fileName string) ([][]int) {
	var arg [][]int
	f, err := os.Open(fileName)
	checkerr(err)
	defer f.Close()
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))
	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "* "))
	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
		var line = strings.Split(scanner.Text(), " ")
		fmt.Println(line)
	}
	checkerr(scanner.Err())
	return arg
}