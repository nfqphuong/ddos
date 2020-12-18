package cmd

import "flag"

func getVars() {
	var (
		cFlag = flag.Int("c", 1, "total number of threads")
		nFlag = flag.Int("n", 1, "total number of requests")
	)
}
