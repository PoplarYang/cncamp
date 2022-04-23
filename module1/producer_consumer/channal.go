package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// type Worker struct {
// 	Name string
// }

type Producer struct {
	Name            string
	ProduceInterval int
	ProduceCounts   int
}

type Consumer struct {
	Name            string
	ConsumeInterval int
	ConsumeCounts   int
}

func (p Producer) produce(ch chan<- int, wg *sync.WaitGroup) {
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= p.ProduceCounts; i++ {
		defer wg.Done()
		data := rand.Intn(1000)
		ch <- data
		fmt.Printf("%s produce goods: %d, this is No. %d, total is %d\n", p.Name, data, i, p.ProduceCounts)
		time.Sleep(time.Second * time.Duration(p.ProduceInterval))
	}
}

func (c Consumer) consume(ch <-chan int, wg *sync.WaitGroup) {
	for i := 1; i <= c.ConsumeCounts; i++ {
		defer wg.Done()
		recv, ok := <-ch
		if ok {
			fmt.Printf("%s consume goods: %d, this is No. %d, total need is %d\n", c.Name, recv, i, c.ConsumeCounts)
		} else {
			fmt.Printf("%s find channal closed\n", c.Name)
		}

		time.Sleep(time.Second * time.Duration(c.ConsumeInterval))
	}
}

func main() {
	goodsCounts := 30
	messages := make(chan int, 10)
	ProcudeWG := sync.WaitGroup{}
	ConsumeWG := sync.WaitGroup{}
	ProcudeWG.Add(goodsCounts)
	ConsumeWG.Add(goodsCounts)

	produce1 := Producer{
		Name:            "produce1",
		ProduceInterval: 1,
		ProduceCounts:   10,
	}

	go produce1.produce(messages, &ProcudeWG)

	produce2 := Producer{
		Name:            "produce2",
		ProduceInterval: 2,
		ProduceCounts:   12,
	}

	go produce2.produce(messages, &ProcudeWG)

	produce3 := Producer{
		Name:            "produce3",
		ProduceInterval: 4,
		ProduceCounts:   8,
	}

	go produce3.produce(messages, &ProcudeWG)

	consume1 := Consumer{
		Name:            "Consume1",
		ConsumeInterval: 1,
		ConsumeCounts:   10,
	}

	go consume1.consume(messages, &ConsumeWG)

	consume2 := Consumer{
		Name:            "Consume2",
		ConsumeInterval: 2,
		ConsumeCounts:   20,
	}
	go consume2.consume(messages, &ConsumeWG)

	ProcudeWG.Wait()
	ConsumeWG.Wait()
}
