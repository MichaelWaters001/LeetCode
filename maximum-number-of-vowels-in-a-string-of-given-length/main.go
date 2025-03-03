package main

func main() {

	tests := []string{"foo", "x", "a", "aeiou"}
	for _, test := range tests {
		println(test, ": ", maxVowels(test, len(test)))
	}
	println("abciiidef: ", maxVowels("abciiidef", 3))

}

func maxVowels(s string, k int) int {
	max := 0   //track the max count of vowels
	cur := 0   // track vowels in window
	left := 0  // left pointer
	right := 0 // right pointer

	for right < len(s) { //keep in bounds
		// ++ if right is vowel
		if isVowel(s[right]) {
			cur++
		}

		// move right
		right++

		// move left if right - left > k
		if k < (right - left) {
			// if left = vowel --
			if isVowel(s[left]) {
				cur--
			}
			left++
		}
		// track max vowels in window
		if max < cur{
			max = cur
		}

	}
	return max
}

func isVowel(r byte) bool {
	return (r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u')
}
