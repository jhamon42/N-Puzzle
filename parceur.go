package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type puzzle struct {
	size  int
	array [][]int
}

func numPuzzle(scanner *bufio.Scanner) int {
	for scanner.Scan() {
		var line = strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if strings.HasPrefix(line[0], "#") || line[0] == "" {
			continue
		}
		a, _ := strconv.Atoi(line[0])
		if j := len(line); j > 1 {
			for i := 1; j > i; i++ {
				if line[i] == " " || line[i] == "" {
					continue
				} else if strings.HasPrefix(line[i], "#") {
					break
				} else {
					return 0
				}
			}
		}
		return a
	}
	return 0
}

func parsPuzzle(scanner *bufio.Scanner, puzz *puzzle) {
	regex := `^(\d{1,9} +){` + strconv.Itoa(puzz.size-1) + `}(\d{1,9} *)(#.*)*$`
	for scanner.Scan() {
		var bac []int
		var tab = strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if strings.HasPrefix(tab[0], "#") || tab[0] == "" {
			continue
		}
		matched, _ := regexp.MatchString(regex, strings.TrimSpace(scanner.Text()))
		if matched == false {
			puzz.array = nil
			return
		}
		var j = len(tab)
		for i := 0; j > i; i++ {
			a, err := strconv.Atoi(tab[i])
			if err != nil {
				continue
			}
			bac = append(bac, a)
		}
		puzz.array = append(puzz.array, bac)
	}
	if len(puzz.array) != puzz.size {
		puzz.array = nil
	}
}

func validPuzzle(puzz *puzzle) {
	for i := 0; puzz.size > i; i++ {
		for j := 0; puzz.size > j; j++ {
			for ii := i + 1; puzz.size > ii; ii++ {
				for jj := j + 1; puzz.size > jj; jj++ {
					if puzz.array[i][j] == puzz.array[ii][jj] {
						puzz.array = nil
						return
					}
				}
			}
		}
	}
}

func parceur(fileName string) *puzzle {
	f, err := os.Open(fileName)
	checkerr(err)
	defer f.Close()
	puzz := &puzzle{0, nil}
	scanner := bufio.NewScanner(f)
	puzz.size = numPuzzle(scanner)
	parsPuzzle(scanner, puzz)
	checkerr(scanner.Err())
	validPuzzle(puzz)
	return puzz
}
