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
		tab := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if strings.HasPrefix(tab[0], "#") || tab[0] == "" {
			continue
		}
		matched, _ := regexp.MatchString(regex, strings.TrimSpace(scanner.Text()))
		if matched == false {
			puzz.array = nil
			return
		}
		j := len(tab)
		for i := 0; j > i; i++ {
			a, err := strconv.Atoi(tab[i])
			if err != nil {
				continue
			}
			puzz.array = append(puzz.array, a)
		}
	}
	if len(puzz.array) != env.longSize {
		puzz.array = nil
	}
}

func validPuzzle(puzz *puzzle, env *env) {
	for i := 0; env.size > i; i++ {
		for j := i + 1; env.size > j; j++ {
			if puzz.array[i] == puzz.array[j] {
				puzz.array = nil
				return
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
	env.longSize = env.size * env.size
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

func parcePuzz(flags *flags, env *env) {
	if flags.file != "" {
		env.actuel = parceFile(flags.file, env)
	} else {
		env.actuel = generator(flags.rand, env)
	}
	env.goal = goalMap(flags, env)
	env.heuri(env.actuel, env)
	checkMap(env.actuel, env)
}
