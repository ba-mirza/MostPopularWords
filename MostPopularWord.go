package main

import "fmt"

func main() {
	_slice := make([]string, 0)
	words := append(_slice, "dsf", "sdfr", "asd", "asd", "asd")
	fmt.Println(MostPopularWords(words))
}

func MostPopularWords(words []string) string {
	wordsCount := make(map[string]int, 0)
	mostPopWord := ""
	max := 0

	for _, word := range words {
		wordsCount[word]++
		if wordsCount[word] > max {
			max = wordsCount[word]
			mostPopWord = word
		}
	}

	return mostPopWord
}
