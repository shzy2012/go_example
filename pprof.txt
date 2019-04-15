package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var list sync.Map

func Counter(wg *sync.WaitGroup) {
	time.Sleep(1e9)
	var counter int
	for i := 0; i < 1000000; i++ {
		time.Sleep(time.Millisecond * 200)
		counter++
	}

	wg.Done()
}

var lock sync.Mutex

func UnUse(id int) {
	log.Printf("go %v", id)
	key := fmt.Sprintf("%v", id)
	list.Delete(key)
	time.Sleep(time.Second * time.Duration(rand.Int63n(60)))
	log.Printf("die %v", id)
}

func block(ch <-chan int) {
	for {
		log.Println("block here")
		<-ch
		log.Println("block release")
	}
}

func main() {
	flag.Parse()
	//远程获取pprof数据
	go func() {
		log.Println(http.ListenAndServe(":10000", nil))
	}()

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Counter(&wg)
	}

	ch := make(chan int)
	go block(ch)

	time.Sleep(time.Second * 5)
	for i := 0; i < 1000000; i++ {
		//time.Sleep(time.Second * time.Duration(rand.Int63n(2)))
		key := fmt.Sprintf("%v", i)
		list.Store(key, key)
		go UnUse(i)
	}

	wg.Wait()

	//sleep 10min,在程序退出前可以查看性能参数
	time.Sleep(time.Second * 60)
}
