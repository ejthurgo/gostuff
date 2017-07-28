package main

import (
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	args := os.Args[1:]

	var wg sync.WaitGroup
	fs := time.Now()
	for _, v := range args {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			ts := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error:", err)
			}
			defer resp.Body.Close()
			fmt.Printf("URL: '%s', Code: %d, Time: %s\n", url, resp.StatusCode, time.Since(ts))
		}(v)
	}
	wg.Wait()
	fmt.Printf("Finished after %s\n", time.Since(fs))

}
