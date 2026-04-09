//go:build !solution

package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	urls := os.Args[1:]

	messageChannel := make(chan string)
	start := time.Now()

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				messageChannel <- err.Error()
				return
			}
			resp.Body.Close()
			messageChannel <- url
		}(url)
	}

	for range len(urls) {
		msg := <-messageChannel
		end := time.Since(start)

		fmt.Println(end, msg)
	}

	fmt.Println("elapsed", time.Since(start))
}
