package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(ch chan<- int) {
	rand.Seed(time.Now().UnixNano())
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		data := rand.Intn(1000)
		ch <- data
		fmt.Println("producer send:", data)
	}
}

func Consumer1(ch <-chan int) {
	ticker := time.NewTicker(time.Second)
	count := 0
	for range ticker.C {
		recv, ok := <-ch
		if ok {
			count += 1
			fmt.Printf("consumer1 recv: %d\n", recv)
		} else {
			fmt.Println("consumer1 channal closed")
		}
	}
}

func Consumer2(ch <-chan int) {
	for recv := range ch {
		fmt.Printf("consumer2 recv: %d\n", recv)
	}
	defer fmt.Println("consumer2 channal closed")
}

func main() {
	messages := make(chan int, 10)
	go Producer(messages)
	// go Consumer1(messages)
	go Consumer2(messages)
	time.Sleep(time.Second * 5)
}
