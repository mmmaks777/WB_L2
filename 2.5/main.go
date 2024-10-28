package main

import (
	"fmt"
	"sort"
	"strings"
)

func createKey(word string) string {
	runes := []rune(word)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func findAnagrams(words []string) map[string][]string {
	anagrams := make(map[string][]string)
	wordOrder := make(map[string]int)

	for idx, word := range words {
		lowerWord := strings.ToLower(word)
		key := createKey(word)

		anagrams[key] = append(anagrams[key], lowerWord)

		if _, exists := wordOrder[key]; !exists {
			wordOrder[key] = idx
		}
	}

	result := make(map[string][]string)

	for key, group := range anagrams {
		if len(group) > 1 {
			uniqueGroup := uniqueAndSort(group)
			if len(uniqueGroup) > 1 {
				firstWordIndex := wordOrder[key]
				firstWord := strings.ToLower(words[firstWordIndex])

				if _, exists := result[firstWord]; !exists {
					result[firstWord] = uniqueGroup
				}
			}
		}
	}
	return result
}

func uniqueAndSort(words []string) []string {
	uniqueWords := make(map[string]bool)
	result := []string{}
	for _, word := range words {
		if !uniqueWords[word] {
			uniqueWords[word] = true
			result = append(result, word)
		}
	}
	sort.Strings(result)
	return result
}

func main() {
	words := []string{"пятак", "тяпка", "пятка", "листок", "слиток", "столик", "кот", "ток", "окт", "кот"}
	anagrams := findAnagrams(words)

	for key, group := range anagrams {
		fmt.Printf("%s: %v\n", key, group)
	}
}
