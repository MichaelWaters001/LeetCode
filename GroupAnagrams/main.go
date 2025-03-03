package main

import "fmt"

func main() {
	anagrams := []string{"eat", "tea", "tan", "ate", "nat", "bat", "fart", "too", "oot", "to"}
	groups := groupAnagrams(anagrams)
	fmt.Println(groups)
}

type word struct {
	words []string     //original word
	m     map[rune]int //map of count of runes in the word
}

func groupAnagrams(strs []string) [][]string {
	mapsOfWords := make([]word, 0)
	// iterate over the list of words
	for _, s := range strs {
		found := false
		m := wordToMap(s)
		// check if the word is already in the list
		for i, w := range mapsOfWords {
			if compareMaps(w.m, m) {
				mapsOfWords[i].words = append(mapsOfWords[i].words, s)
				found = true
				break
			}
		}
		// if the word is not found, add it to the list as a new word
		if !found {
			mapsOfWords = append(mapsOfWords, word{[]string{s}, m})
		}
	}

	r := make([][]string, 0)
	for _, w := range mapsOfWords {
		r = append(r, w.words)
	}
	return r
}

// wordToMap converts a word to a map of runes
func wordToMap(word string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range word {
		m[r]++
	}
	return m
}

// compareMaps compares two maps of runes
func compareMaps(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v := range m1 {
		if m2[k] != v {
			return false
		}
	}
	return true
}
