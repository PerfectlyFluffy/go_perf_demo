package main

import (
	"fmt"

	"github.com/PerfectlyFluffy/go_perf_demo/demo"
)

func main() {
	// Find a number that will make the benchmark
	// last between 3 to 5 secondes.
	loop_count := uint64(52500000000)

	fmt.Print()          // => No impact on performance!
	demo.EmptyFunction() // => No impact on performance!
	demo.RunOutOfScope() // => No impact on performance!

	demo.Run(loop_count)

	demo.RunOutOfScope() // => No impact on performance!
	demo.EmptyFunction() // => No impact on performance!
	fmt.Print()          // => No impact on performance!
}
