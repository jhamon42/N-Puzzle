package main

func main() {
	flags := parceFlags()
	env := parceEnv(flags)
	parcePuzz(flags, env)

	if flags.inter {
		gameInteractive(env)
	} else {
		if flags.visu {
			initTerm()
			defer endTerm()
		}
		gameAlgo(env)
		gameResume(env)
	}
}
