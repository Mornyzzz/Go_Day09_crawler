package main

import (
	"fmt"
	"sync"
	"time"
)

func sleepSort(arr []int) chan int {
	c := make(chan int)
	var wg sync.WaitGroup

	for i := range arr {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			time.Sleep(time.Duration(val) * time.Millisecond)
			c <- val
		}(arr[i])
	}

	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}

func main() {
	arr := []int{3, 123, 432, 12, 76, 32, 87, 8, 9, 10, 11, 12, 27, 28, 29}
	intChan := sleepSort(arr)
	for val := range intChan {
		fmt.Println(val)
	}
}
