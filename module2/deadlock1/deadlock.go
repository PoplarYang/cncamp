package main

// 可以利用sync.WaitGroup解决，在所有的 data channel 的输入处理之前，wg.Wait()这个goroutine会处于等待状态（wg.Wait()源码就是for循环）。当执行方法处理完后（wg.Done），wg.Wait()就会放开执行，执行后面的close(ch)。
import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func produce(index int, ch chan<- string) {
	time.Sleep(time.Duration(index) * time.Second)
	s := fmt.Sprintf("No.%d compeleted", index)
	ch <- s
	defer wg.Done()
}

func main() {
	ch := make(chan string, 10)

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go produce(i, ch)
	}

	for ret := range ch {
		fmt.Println(len(ch))
		fmt.Println(ret)
	}
}
