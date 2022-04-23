package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	var (
		mtx1 sync.Mutex
		mtx2 sync.Mutex
		a    int
	)
	a = 114
	wg.Add(1)
	go func() {
		//a=114
		mtx1.Lock() //加个锁

		fmt.Println("[1]-->: ", a) //输出一下设定的a的值
		//睡眠5s后再输出一下a的值，因为在2s后a的值已经在主线程中被更改了
		time.Sleep(5 * time.Second)
		fmt.Println("[2]-->: ", a)
		mtx1.Unlock() //解锁
		wg.Done()
	}()

	//睡眠2s后更改一下a的值
	time.Sleep(2 * time.Second)
	mtx2.Lock() //加个锁
	a = 514
	mtx2.Unlock() //解锁

	fmt.Println("[3]-->: ", a)
	wg.Wait()
}
