package main

import (
	"runtime"
	"testing"
)

func Benchmark_Channel(b *testing.B) {
	runtime.GOMAXPROCS(1)
	queue := make(chan int)

	go func() {
		for {
			<-queue
		}
	}()

	for i := 0; i < b.N; i++ {
		queue <- i
	}

	close(queue)
}
