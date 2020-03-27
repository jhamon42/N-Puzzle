package main

func parceHeuristic(flags *Flags, env *Env) {
	switch flags.her {
	case "hamming":
		env.heuri = hamming
	case "manhattan":
		env.heuri = manhattan
	case "IC":
		env.heuri = invCount
	}
}

func parceAlgo(flags *Flags, env *Env) {
	switch flags.algo {
	case "a*":
		env.algo = aStarAlgo
	// case "ida*":
	// env.algo = idaStarAlgo
	case "aStar":
		env.algo = aStarAlgo
		// case "idaStar":
		// env.algo = idaStarAlgo
	}
	// case "BDF":
	// 	env.algo =
	// case "RF":
	// 	env.algo =
	// case "BF":
	// 	env.algo =
}

func parceEnv(flags *Flags) *Env {
	tmp := &Env{}

	parceAlgo(flags, tmp)
	parceHeuristic(flags, tmp)

	return tmp
}
