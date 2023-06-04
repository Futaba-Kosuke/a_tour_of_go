package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	results := make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		_, exists := results[word]
		if exists {
			results[word] = 1
		} else {
			results[word] += 1
		}
	}
	return results
}

func main() {
	wc.Test(WordCount)
}
