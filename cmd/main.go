package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	requests := flag.Int("requests", 0, "Número total de requests")
	concurrency := flag.Int("concurrency", 0, "Número de chamadas simultâneas")

	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		flag.Usage()
		os.Exit(1)
	}

	start := time.Now()
	results := runLoadTest(*url, *requests, *concurrency)
	duration := time.Since(start)

	printReport(results, duration)
}

func printReport(results map[int]int, duration time.Duration) {
	totalRequests := 0
	successfulRequests := 0

	for status, count := range results {
		totalRequests += count
		if status == 200 {
			successfulRequests += count
		}
	}

	fmt.Printf("Tempo total gasto: %v\n", duration)
	fmt.Printf("Quantidade total de requests realizados: %d\n", totalRequests)
	fmt.Printf("Quantidade de requests com status HTTP 200: %d\n", successfulRequests)
	fmt.Println("Distribuição de outros códigos de status HTTP:")
	for status, count := range results {
		if status != 200 {
			fmt.Printf("Status %d: %d\n", status, count)
		}
	}
}

func runLoadTest(url string, totalRequests, concurrency int) map[int]int {
	results := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, concurrency)

	for i := 0; i < totalRequests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err != nil {
				log.Printf("Erro ao realizar request: %v", err)
				<-semaphore
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			results[resp.StatusCode]++
			mu.Unlock()

			<-semaphore
		}()
	}

	wg.Wait()
	return results
}
