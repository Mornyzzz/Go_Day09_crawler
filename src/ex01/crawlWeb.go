package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func crawlWeb(ctx context.Context, input <-chan string) <-chan string {
	result := make(chan string)
	go func() {
		defer close(result)
		var wg sync.WaitGroup
		semaphore := make(chan struct{}, 8)
		for url := range input {
			select {
			case <-ctx.Done():
				return
			default:
			}
			wg.Add(1)
			semaphore <- struct{}{}
			go func(url string) {
				defer func() {
					<-semaphore // Удаляем структуру из канала после завершения работы горутины
					wg.Done()
				}()
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Error fetching URL %s: %v\n", url, err)
					return
				}
				defer resp.Body.Close()

				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Printf("Error reading body for URL %s: %v\n", url, err)
					return
				}
				result <- url
				bodyString := string(body)
				result <- bodyString
				time.Sleep(1 * time.Second)
			}(url)
		}
		wg.Wait()
	}()
	return result
}

func main() {
	start := time.Now()
	urls := make(chan string, 100)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for i := 0; i < 23; i++ {
		urls <- "https://example.com"
	}
	close(urls)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGINT)
		<-c
		fmt.Println("\nCrawling stopped")
		cancel()
	}()

	page := crawlWeb(ctx, urls)

	count := 0
	for x := range page {
		fmt.Println(x)
		fmt.Printf("-- %d --\n", count)
		count++

	}
	duration := time.Since(start).Milliseconds()
	fmt.Printf("Время выполнения: %d микросекунд\n", duration)

}
