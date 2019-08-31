package main

func parceHeuristic(flags *flags, env *env) {
	switch flags.her {
	case "manhattan":
		env.heuri = diff1
	case "hamming":
		env.heuri = diff3
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
func parceEnv(flags *flags, env *env) {
	parceAlgo(flags, env)
	parceHeuristic(flags, env)
}
