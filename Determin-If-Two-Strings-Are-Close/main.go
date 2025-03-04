package main

import (
	"reflect"
	"slices"
)

func main() {
	word1 := "abbzzca"
	word2 := "babzzcz"
	println(closeStrings(word1, word2))
}

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	//check the same runes are used
	runes1 := wordToRuneMap(word1)
	runes2 := wordToRuneMap(word2)

	if !reflect.DeepEqual(runes1, runes2) {
		return false
	}

	//check counts are equal
	counts1 := getRuneCounts(word1)
	counts2 := getRuneCounts(word2)

	return reflect.DeepEqual(counts1, counts2)
}

func wordToRuneMap(word string) (runes map[rune]struct{}) {
	runes = make(map[rune]struct{})
	for _, r := range word {
		runes[r] = struct{}{}
	}
	return
}

func getRuneCounts(word string) []int {
	runes := make(map[rune]int)
	for _, r := range word {
		runes[r]++
	}
	var counts []int
	for _, c := range runes {
		counts = append(counts, c)
	}
	slices.Sort(counts)
	return counts
}