package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	repeat := func(done <-chan interface{}, values ...interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
				time.Sleep(1e9)
			}
		}()
		return valueStream
	}

	repeatFn := func(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	toString := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- v.(string):
				}
			}
		}()
		return stringStream
	}

	toInt := func(done <-chan interface{}, valueStream <-chan interface{}) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			select {
			case <-done:
				return
			case v, _ := <-valueStream:
				intStream <- v.(int)
			}
		}()
		return intStream
	}

	take := func(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i <= num; i++ {
				select {
				case <-done:
					return
				case value, _ := <-valueStream:
					//fmt.Printf("%v", v)
					switch value.(type) {
					case int:
						takeStream <- value.(int)
					case string:
						takeStream <- value.(string)
					}
				}
			}
		}()
		return takeStream
	}

	//扇入:将多个数据流复用或者合并成一个流.
	fanIn := func(done <-chan interface{}, channels ...<-chan interface{}) <-chan interface{} {
		var wg sync.WaitGroup
		multiplexedStream := make(chan interface{})
		multiplex := func(c <-chan interface{}) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}
		//从所有的channle里取值
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		//等待所有的读操作结束
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	done := make(chan interface{})
	defer close(done)

	randFunc := func() interface{} {
		return rand.Int()
	}

	for ele := range take(done, repeat(done, 1), 10) {
		fmt.Printf("%v", ele)
	}

	fmt.Println()
	for num := range take(done, repeatFn(done, randFunc), 10) {
		fmt.Println(num)
	}

	var message string
	for token := range toString(done, take(done, repeat(done, "I", "Am."), 5)) {
		message += token
	}

	fmt.Println(message)
}
