package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type (
	subscribe chan interface{}         //订阅者为一个管道
	topicFunc func(v interface{}) bool //主题为一个过滤器
)

type Publisher struct {
	m          sync.RWMutex            //读写锁
	buffer     int                     //订阅队列的缓存大小
	timeout    time.Duration           //发布超时时间
	subscribes map[subscribe]topicFunc //订阅者信息
}

//发布一个主题
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range p.subscribes {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

//构建一个发布者对象
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    publishTimeout,
		subscribes: make(map[subscribe]topicFunc),
	}
}

//添加一个新的订阅者,订阅全部主题
func (p *Publisher) Subscribe() chan interface{} {
	return p.subscribeTopic(nil)
}

//添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) subscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribes[ch] = topic
	p.m.Unlock()
	return ch
}

//退出订阅
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()

	delete(p.subscribes, sub)
	close(sub)
}



//关闭发布者对象,同事关闭所有的订阅者管道
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	for sub := range p.subscribes {
		delete(p.subscribes, sub)
		close(sub)
	}

}

//发送主题,可以容忍一定的超时
func (p *Publisher) sendTopic(sub subscribe, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):
	}
}

/*
发布订阅（publish-and-subscribe）模型通常被简写为pub/sub模型。在这个模型中，消息生产者成为发布者（publisher），
而消息消费者则成为订阅者（subscriber），生产者和消费者是M:N的关系。在传统生产者和消费者模型中，是将消息发送到一个队列中，而发布订阅模型则是将消息发布给一个主题。
*/
func main() {
	p := NewPublisher(100*time.Microsecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.subscribeTopic(func(v interface{}) bool {
		//fmt.Printf("%s\n", v)
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,world")
	p.Publish("hello,golang")

	go func() {
		for msg := range all {
			fmt.Println("all:", msg)
		}
	}()

	go func() {
		for msg := range golang {
			fmt.Println("golang:", msg)
		}
	}()

	time.Sleep(time.Second * 30)
}
