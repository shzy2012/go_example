package main

import (
	"log"
	"runtime"
	"time"
)

const (
	workerNum    = 100
	producterNum = 1
)

var (
	done chan bool
	p    chan bool
)

/*
 * 1：没有初始化的channle会阻塞
 * 2: 如果没有使用close关闭channle,则channle的返回值不会返回false
 * 生产者与消费者最好的比例：M:N 1:100
 */

func consumer(ch <-chan int) {
	for {
		v, more := <-ch
		if !more {
			//log.Println("all job done", v)
			done <- true
			return
		}

		log.Println("get:", v)
		time.Sleep(1e9)
	}
}

func product(queue chan<- int, p chan<- bool) {
	for i := 0; i <= 100; i++ {
		queue <- i
		log.Println("put:", i)
	}

	p <- true
}

func closeChannle(queue chan int, p chan bool) {
	i := 0
	for {
		<-p
		i++
		if i == producterNum {
			break
		}
	}

	close(p)
	close(queue)
}

func main() {

	//runtime.GOMAXPROCS(1)

	log.Println("current cpu num:", runtime.NumCPU())
	queue := make(chan int, 100)
	done = make(chan bool)
	p = make(chan bool)

	for i := 0; i < producterNum; i++ {
		go product(queue, p)
	}

	for i := 0; i < workerNum; i++ {
		go consumer(queue)
	}

	go closeChannle(queue, p)

	//control exit
	count := 0
	for range done {
		count++
		if count == workerNum {
			break
		}
	}
}
