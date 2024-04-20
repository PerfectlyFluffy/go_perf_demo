package demo

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	Run(LOOP_COUNT)
}
