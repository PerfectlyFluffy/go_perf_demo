package demo

import (
	"math"
	"runtime"
)

type batch struct {
	start uint64
	stop  uint64
}

type context struct {
	batchCount  int
	batches     []batch
	loopCount   uint64
	threadCount int
}

func newContext(loopCount uint64) context {
	ctx := context{
		loopCount:   loopCount,
		threadCount: runtime.NumCPU(),
	}
	ctx.batchCount = math.MaxUint16 - (math.MaxUint16 % ctx.threadCount)
	ctx.setBatches()
	return ctx
}

func (c *context) setBatches() {
	c.batches = make([]batch, c.batchCount)
	batchSize := c.loopCount / uint64(c.batchCount)
	diffCount := c.loopCount - batchSize*uint64(c.batchCount)

	start := uint64(1)
	for i := 0; i < c.batchCount; i++ {
		stop := start + batchSize
		if diffCount > 0 {
			stop++
			diffCount--
		}

		b := batch{
			start: start,
			stop:  stop,
		}
		c.batches[i] = b

		start = stop
	}
}
