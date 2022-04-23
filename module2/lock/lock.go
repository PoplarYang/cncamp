package main

import (
	"fmt"
	"sync"
)

// 会产生死锁
func lock() {
	lock := sync.Mutex{}

	for i := 1; i <= 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println(i)
	}
}

// 会产生死锁
func rwlock() {
	lock := sync.RWMutex{}

	for i := 1; i <= 3; i++ {
		lock.Lock()
		defer lock.Unlock()
		fmt.Println(i)
	}
}

func rlock() {
	lock := sync.RWMutex{}

	for i := 1; i <= 3; i++ {
		lock.RLock()
		defer lock.RUnlock()
		fmt.Println(i)
	}
}

func main() {
	rlock()
}
