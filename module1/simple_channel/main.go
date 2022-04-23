package main

import "fmt"

func main() {
	// 同步channel
	ch := make(chan int, 1)
	defer close(ch)
	go func(ch chan int) {
		fmt.Println("child thread run start")
		ch <- 1
		fmt.Println("child thread run end")
	}(ch)
	fmt.Println("main thread run start")
	<-ch
	fmt.Println("main thread run end")
}
