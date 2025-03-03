package main

import "strings"

func main() {
	tests := []string{"the sky is blue", "double  space", " word", "end ", "single", "  ", " "}
	for _, test := range tests {
		reversed := reverseWords(test)
		println(reversed)
	}
}

func reverseWords(s string) string {

	words := strings.Split(s, " ")

	i := len(words)

	reversed := ""

	for i > 0 {
		i--
		if len(words[i]) == 0 {
			continue
		}
		reversed = reversed + " " + words[i]
	}

	reversed, _ = strings.CutPrefix(reversed, " ")
	return reversed

}
