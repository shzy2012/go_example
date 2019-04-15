package main
import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	//定义channel，用于goroutine之间数据传递
	messags := make(chan string)

	//定义生产者
	producter := func(ch chan string) {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			ch <- fmt.Sprintf("%v", i)
			fmt.Printf("producter %v\n", i)
		}
	}

	//定义消费者
	consumer := func(ch chan string) {
		for {
			select {
			case k, ok := <-ch:
				if !ok {
					return
				}
				fmt.Printf("consumer %v\n", k)
				wg.Done()
			}
		}
	}
	//多个消费者模式
	for i := 0; i < 3; i++ {
		go consumer(messags)
	}

	//开始生产
	producter(messags)
	//等待任务结束
	wg.Wait()
}
