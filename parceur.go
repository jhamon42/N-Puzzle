package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func parsPuzzle(scanner *bufio.Scanner, puzz *puzzle, env *env) {
	regex := `^(\d{1,9} +){` + strconv.Itoa(env.size-1) + `}(\d{1,9} *)(#.*)*$`
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
	if len(puzz.array) != env.size {
		puzz.array = nil
	}
}

func validPuzzle(puzz *puzzle, env *env) {
	for i := 0; env.size > i; i++ {
		for j := 0; env.size > j; j++ {
			for ii := i; env.size > ii; ii++ {
				for jj := 0; env.size > jj; jj++ {
					if puzz.array[i][j] == puzz.array[ii][jj] && i+j < ii+jj {
						puzz.array = nil
						return
					}
				}
			}
		}
	}
}

func parceFile(fileName string, env *env) *puzzle {
	f, err := os.Open(fileName)
	checkerr(err)
	defer f.Close()
	puzz := &puzzle{}
	scanner := bufio.NewScanner(f)
	env.size = numPuzzle(scanner)
	parsPuzzle(scanner, puzz, env)
	checkerr(scanner.Err())
	if puzz.array == nil {
		log.Fatalf("file parsing %s: error", fileName)
	}
	validPuzzle(puzz, env)
	if puzz.array == nil {
		log.Fatalf("file parsing %s: error (duplicat)", fileName)
	}
	return puzz
}

func parce(flags *flags, env *env) *puzzle {
	puzz := &puzzle{}
	if flags.file != "" {
		puzz = parceFile(flags.file, env)
	} else {
		puzz = generator(flags.rand, env)
	}
	puzz.zero = &coord{}
	checkMap(puzz, env)
	return puzz
}
