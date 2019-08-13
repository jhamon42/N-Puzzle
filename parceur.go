package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type puzzle struct {
	size  int
	array [][]int
}

func numPuzzle(scanner *bufio.Scanner) (int, error) {
	for scanner.Scan() {
		var line = strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if strings.HasPrefix(line[0], "#") || line[0] == "" {
			continue
		}
		a, err := strconv.Atoi(line[0])
		if j := len(line); j > 1 {
			for i := 1; j > i; i++ {
				if line[i] == " " || line[i] == "" {
					continue
				} else if strings.HasPrefix(line[i], "#") {
					break
				} else {
					return 0, nil
				}
			}
		}
		return a, err
	}
	return 0, nil
}

func parceur(fileName string) [][]int {
	f, err := os.Open(fileName)
	checkerr(err)
	defer f.Close()
	puzz := &puzzle{0, nil}
	scanner := bufio.NewScanner(f)
	puzz.size, err = numPuzzle(scanner)
	fmt.Println(puzz.size)
	checkerr(scanner.Err())
	return puzz.array
}
