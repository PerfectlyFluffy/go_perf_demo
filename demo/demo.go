package demo

import (
	"fmt"
	"runtime"
	"sync"
)

var WG_TEST sync.WaitGroup

func Run(loopCount uint64) {
	//fmt.Print() // => Impact performance! 1/2
	//fmt.Print("") // => Impact performance! 1/3
	//Func0Param() // => Impact performance! 1/2
	//Func1Param("") // => Impact performance! 1/3
	//Func3Param("", "", "") // => Impact performance! 1/3
	//WG_TEST.Add(0) // => Impact performance! 1/3

	EmptyFunction()    // => No impact on performance!
	Func2Param("", "") // => No impact on performance!

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

func EmptyFunction() {}

func RunOutOfScope() {
	fmt.Print()
}

func Func0Param() {
	fmt.Print()
}

func Func1Param(empty string) {
	fmt.Print(empty)
}

func Func2Param(empty string, empty2 string) {
	fmt.Print(empty, empty2) // => Will **not** impact performance if present only one time.
	// fmt.Print(empty, empty2) // => Will impact performance if uncommented.
}

func Func3Param(empty string, empty2 string, empty3 string) {
	fmt.Print(empty, empty2, empty3)
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
