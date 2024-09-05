package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func makeRequest(url string, client *http.Client, wg *sync.WaitGroup, requestChan chan<- int, index int) {
	defer wg.Done()

	for {
			// start := time.Now()
			resp, err := client.Get(url)
			if err != nil {
					fmt.Println("Error:", err)
					continue
			}

			defer resp.Body.Close()

			_, err = io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading response:", err)
				continue
			}
			// elapsed := time.Since(start)
			// fmt.Printf("Request took %s\n", elapsed)
			requestChan <- index
	}
}

func loadTest() {
	url := "https://ca-wrkshp-ary-go.wittywave-5d3299aa.westus.azurecontainerapps.io" // Replace with your target URL
	numConcurrentRequests := 2000
	client := &http.Client{
		Transport: &http.Transport{
				MaxIdleConns:       2500,
				IdleConnTimeout:    30 * time.Second,
				// DisableKeepalive:   false,
				MaxIdleConnsPerHost: 2500,
				TLSHandshakeTimeout: 10 * time.Second,
		},
}
	var wg sync.WaitGroup
	requestChan := make(chan int)

	go func() {
			tick := time.NewTicker(1 * time.Second)
			defer tick.Stop()

			var totalRequests int
			for range tick.C {
					totalRequests += <-requestChan
					fmt.Printf("Requests per second: %.2f\n", float64(totalRequests))
					totalRequests = 0
			}
	}()

	for i := 0; i < numConcurrentRequests; i++ {
			wg.Add(1)
			go makeRequest(url, client, &wg, requestChan, i)
	}
	wg.Wait()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! From Aryan's Computer")
	})

	http.HandleFunc("/loadTest", func(w http.ResponseWriter, r *http.Request) {
		loadTest()
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
