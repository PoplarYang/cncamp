package main

import (
	"fmt"
	"sync"
	"time"
)

type Queue struct {
	queue []string
	cond  *sync.Cond
}

func (q *Queue) Enqueue(mesg string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, mesg)
	fmt.Printf("putting %s to queue, notify all\n", mesg)
	q.cond.Broadcast()
}

func (q *Queue) Dequeue() (mesg string) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()
	if len(q.queue) == 0 {
		fmt.Println("no data available, wait")
		q.cond.Wait()
	}
	mesg = q.queue[0]
	q.queue = q.queue[1:]
	return
}

func main() {
	queue := Queue{
		queue: []string{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}

	// 生产30条数据并消费
	dataCounts := 30
	go func() {
		for i := 0; i < dataCounts; i++ {
			queue.Enqueue("a")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < dataCounts; i++ {
		mesg := queue.Dequeue()
		time.Sleep(time.Second)
		fmt.Printf("getting %s from queue\n", mesg)
	}
}
