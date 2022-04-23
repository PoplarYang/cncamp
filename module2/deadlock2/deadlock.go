package main

// 适用于通道不关闭，需要时刻循环执行数据并且处理的情境下

import (
	"fmt"
	"time"
)

func produce(index int, ch chan<- string) {
	time.Sleep(time.Duration(index) * time.Second)
	s := fmt.Sprintf("No.%d complete", index)
	ch <- s
}

func main() {
	ch := make(chan string, 10)

	for i := 0; i < 4; i++ {
		go produce(i, ch)
	}

	for {
		select {
		case i := <-ch: // select会一直等待，直到某个case的通信操作完成时，就会执行case分支对应的语句
			println(i)
		default:
			fmt.Println("Channel has no data")
			time.Sleep(time.Second)
		}
	}
}
