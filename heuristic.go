package main

func parceEnv(flags *flags, env *env) {
	switch flags.her {
	case "manhattan":
		env.heuri = diff1
	case "hamming":
		env.heuri = diff3
	}
}
