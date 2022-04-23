package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 先入后出顺序执行
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	loopFunc()
	time.Sleep(time.Second)
}

func loopFunc() {
	lock := sync.Mutex{}
	for i := 0; i < 3; i++ {
		go func(i int) {
			lock.Lock()
			// 及时释放锁
			defer lock.Unlock()
			fmt.Println("loopFunc:", i)
		}(i)
	}
}
