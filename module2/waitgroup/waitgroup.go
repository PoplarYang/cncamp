package main

import (
	"fmt"
	"sync"
	"time"
)

// 通过sleep实现等待子线程结束
func waitBySleep() {
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
	time.Sleep(time.Second)
}

// 通过channel等待子线程结束
func waitByChannel() {
	ch := make(chan int, 3)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		defer close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}

// 通过waitgroup等待子线程结束
func waitByWaitGroup() {
	wg := sync.WaitGroup{}
	// 如果此处放入子线程，可能会导致没执行add主线程就结束了
	wg.Add(3)
	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Println(i)
			defer wg.Done()
		}
	}()
	wg.Wait()
}

func main() {
	waitBySleep()
	waitByChannel()
	waitByWaitGroup()
}
