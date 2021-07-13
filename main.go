package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var result map[string]int = make(map[string]int, 10)
	words := strings.Split(s, " ")
	for _, w := range words {
		_, ok := result[w]
		if ok {
			result[w] += 1
		} else {
			result[w] = 1
		}
	}

	return result
}

func main() {
	wc.Test(WordCount)
}
