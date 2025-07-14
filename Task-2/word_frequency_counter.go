package main

import (
	"fmt"
	"strings"
)

func main() {
	result := wordCounter("I want to be a pro Go developer!")

	fmt.Print(result)
}

func wordCounter(s string) map[string]int {

	count := map[string]int{}

	words := strings.Fields(s)

	for _, word := range words {
		word = strings.ToLower(word)
		word = strings.TrimSpace(word)
		count[word] += 1
	}

	return count
}
