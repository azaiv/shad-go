//go:build !solution

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	files := os.Args[1:]

	counts := map[string]int64{}

	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		lines := strings.Split(string(data), "\n")

		for _, line := range lines {
			counts[line]++
		}
	}

	for line, count := range counts {
		if count >= 2 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
