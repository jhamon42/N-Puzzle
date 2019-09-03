package main

func parceHeuristic(flags *flags, env *env) {
	switch flags.her {
	case "hamming":
		env.heuri = hamming
	case "manhattan":
		env.heuri = manhattan
	}
}

func parceAlgo(flags *flags, env *env) {
	switch flags.algo {
	case "a*":
		env.algo = aStarAlgo
	case "ida*":
		env.algo = idaStarAlgo
	case "aStar":
		env.algo = aStarAlgo
	case "idaStar":
		env.algo = idaStarAlgo
	}
	// case "BDF":
	// 	env.algo =
	// case "RF":
	// 	env.algo =
	// case "BF":
	// 	env.algo =
}

func parceEnv(flags *flags) *env {
	tmp := &env{}
	tmp.actuel = &puzzle{}

	parceAlgo(flags, tmp)
	parceHeuristic(flags, tmp)

	return tmp
}
