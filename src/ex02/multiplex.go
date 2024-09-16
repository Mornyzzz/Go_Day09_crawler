package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func multiplex(channels ...<-chan interface{}) <-chan interface{} {
	fanIn := make(chan interface{})
	wg := sync.WaitGroup{}
	for _, ch := range channels {
		wg.Add(1)
		go func(ch <-chan interface{}) {
			defer wg.Done()
			for v := range ch {
				fanIn <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(fanIn)
	}()
	return fanIn
}

func main() {
	chan1 := make(chan interface{}, 100)
	chan2 := make(chan interface{}, 100)
	chan3 := make(chan interface{}, 100)

	res := multiplex(chan1, chan2, chan3)
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	for i := 0; i < 6; i++ {
		a := random.Intn(100)
		switch random.Intn(3) + 1 {
		case 1:
			chan1 <- a
			fmt.Printf("channel 1: %d\n", a)
		case 2:
			a += 100
			chan2 <- a
			fmt.Printf("channel 2: %d\n", a)
		case 3:
			a += 200
			chan3 <- a
			fmt.Printf("channel 3: %d\n", a)
		}
	}

	close(chan1)
	close(chan2)
	close(chan3)

	fmt.Println("--result--")
	for msg := range res {
		fmt.Printf("fanIn: %d\n", msg)
	}
}
