package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Value struct {
	Productor int
	Value     int
}

func main() {

	/**** 原型
	spawn := func() chan int {
		ch := make(chan int)
		go func() {
			ch <- rand.Intn(100)
		}()

		return ch
	}

	ch1 := spawn()
	ch2 := spawn()

	for i := 0; i < 2; i++ {
		select {
		case n := <-ch1:
			fmt.Printf("chi: %d\n", n)
		case n := <-ch2:
			fmt.Printf("ch2: %d\n", n)
		}
	}
	*/

	spawnProducter := func(i int) chan int {
		ch := make(chan int)
		go func() {
			for {
				ch <- rand.Intn(100)
			}
		}()
		return ch
	}

	nProducter := 3
	inputs := make([]chan int, nProducter)
	for i := 0; i < nProducter; i++ {
		inputs[i] = spawnProducter(i)
	}

	output := make(chan Value)

	for i := 0; i < nProducter; i++ {
		go func(i int) {
			for {
				value := <-inputs[i]
				output <- Value{Productor: i, Value: value}
			}
		}(i)
	}

	for value := range output {
		fmt.Printf("%d: %d\n", value.Productor, value.Value)
		time.Sleep(250 * time.Millisecond)
	}
}
