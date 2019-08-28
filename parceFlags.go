package main

import (
	"flag"
	"os"
	"strings"
)

func validFlags(tmp *flags) {
	if !(strings.EqualFold(tmp.algo, "a*") ||
		strings.EqualFold(tmp.algo, "ida*") ||
		strings.EqualFold(tmp.algo, "aStar") ||
		strings.EqualFold(tmp.algo, "idaStar") ||
		strings.EqualFold(tmp.algo, "BDF") ||
		strings.EqualFold(tmp.algo, "RF") ||
		strings.EqualFold(tmp.algo, "BF")) {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !(strings.EqualFold(tmp.her, "manhattan") ||
		strings.EqualFold(tmp.her, "hamming")) {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if !(strings.EqualFold(tmp.goal, "basic") ||
		strings.EqualFold(tmp.goal, "snail")) {
		flag.PrintDefaults()
		os.Exit(1)
	}
	if tmp.file == "" && tmp.rand < 3 {
		flag.PrintDefaults()
		os.Exit(1)
	}
}

func parceFlags() *flags {
	tmp := &flags{}
	flag.StringVar(&tmp.algo, "Algo", "aStar", "Algo :\n a* - aStar\n ida* - idaStar\n BF - brutForce\n RF - rowFirst\n BDF - breadthDepthFirst\n")
	flag.StringVar(&tmp.her, "Heuristic", "manhattan", "Heuristic : manhattan / hamming")
	flag.StringVar(&tmp.file, "File", "", "any map file")
	flag.IntVar(&tmp.rand, "Random", 4, "size of the map randomly generate min: 3")
	flag.BoolVar(&tmp.quiet, "Quiet", false, "true to silence end stats")
	flag.BoolVar(&tmp.visu, "Visu", false, "true to see visu")
	flag.BoolVar(&tmp.inter, "Interactive", false, "true to play game")
	flag.StringVar(&tmp.goal, "Goal", "basic", "chose type of goal : basic or snail")
	flag.Parse()
	validFlags(tmp)
	return tmp
}
