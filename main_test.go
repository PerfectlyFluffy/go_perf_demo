package main

import (
	"testing"

	"github.com/PerfectlyFluffy/go_perf_demo/demo"
)

func BenchmarkMain(b *testing.B) {
	main()
}

func BenchmarkDemoRun(b *testing.B) {
	demo.Run(demo.LOOP_COUNT)
}
