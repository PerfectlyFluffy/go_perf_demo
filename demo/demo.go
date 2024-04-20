package demo

import (
	"fmt"
	"runtime"
	"sync"
)

// Find a number that will make the benchmark
// last between 3 to 5 secondes.
var LOOP_COUNT = uint64(52500000000)

func Run(loopCount uint64) {
	// Add "" to fmt.Print() to get lost performance back.
	fmt.Print() // ==> fmt.Print("")

	ctx := newContext(loopCount)
	runtime.GOMAXPROCS(ctx.threadCount)

	buffer := make(chan int, ctx.threadCount)
	var wg sync.WaitGroup
	wg.Add(ctx.batchCount)
	for i, b := range ctx.batches {
		buffer <- i
		go func(b batch) {
			defer wg.Done()
			findPerfectNumbers(b)
			<-buffer
		}(b)
	}
	wg.Wait()
	close(buffer)
}

func findPerfectNumbers(b batch) {
	for n := b.start; n < b.stop; n++ {
		if n%2 == 1 {
			continue
		}

		smallDivisor := uint64(1)
		aliquotSum := uint64(1)
		bigDivisor := n / 2
		for smallDivisor < bigDivisor && bigDivisor%2 == 0 {
			smallDivisor *= 2
			aliquotSum += smallDivisor + bigDivisor
			bigDivisor /= 2
		}

		if bigDivisor%2 == 1 {
			aliquotSum += smallDivisor*2 + bigDivisor
			if aliquotSum == n && isPerfect(n, smallDivisor) {
				fmt.Print()
			}
		}
	}
}

func isPerfect(n uint64, smallDivisor uint64) bool {
	for odd := uint64(3); odd < smallDivisor; odd += 2 {
		if n%odd == 0 {
			return false
		}
	}
	return true
}
