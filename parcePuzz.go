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

func filePuzzle(scanner *bufio.Scanner, env *Env) {
	regex := `^(\d{1,9} +){` + strconv.Itoa(env.size-1) + `}(\d{1,9} *)(#.*)*$`
	for scanner.Scan() {
		tab := strings.Split(strings.TrimSpace(scanner.Text()), " ")
		if strings.HasPrefix(tab[0], "#") || tab[0] == "" {
			continue
		}
		matched, _ := regexp.MatchString(regex, strings.TrimSpace(scanner.Text()))
		if matched == false {
			env.state.puzArray = nil
			return
		}
		j := len(tab)
		for i := 0; j > i; i++ {
			a, err := strconv.Atoi(tab[i])
			if err != nil {
				continue
			}
			env.state.puzArray = append(env.state.puzArray, a)
		}
	}
	if len(env.state.puzArray) != env.longSize {
		env.state.puzArray = nil
	}
}

func validPuzzle(env *Env) {
	for i := 0; env.size > i; i++ {
		for j := i + 1; env.size > j; j++ {
			if env.state.puzArray[i] == env.state.puzArray[j] {
				env.state.puzArray = nil
				return
			}
		}
	}
}

func parceFile(fileName string, env *Env) {
	f, err := os.Open(fileName)
	checkerr(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	env.size = numPuzzle(scanner)
	env.longSize = env.size * env.size
	filePuzzle(scanner, env)
	checkerr(scanner.Err())
	if env.state.puzArray == nil {
		log.Fatalf("file parsing %s: error", fileName)
	}
	validPuzzle(env)
	if env.state.puzArray == nil {
		log.Fatalf("file parsing %s: error (duplicat)", fileName)
	}
	return
}

func parcePuzz(flags *Flags, env *Env) {
	if flags.file != "" {
		parceFile(flags.file, env)
	} else {
		generator(flags.rand, env)
	}
	goalMap(flags, env)
	env.state.puzPrevCost(env)
	checkMap(env)
	puzz := Puzzle{}
	puzz.label = arrayToString(env.state.puzArray, " ")
	puzz.f = env.state.h
	env.state.puz = &puzz
	env.initial = env.state.puz
}
