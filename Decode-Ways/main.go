package main

import (
	"fmt"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// Initialize dp[0] and dp[1]
	prev, current := 1, 1

	// Check for leading zeros
	if s[0] == '0' {
		return 0
	}

	for i := 1; i < n; i++ {
		// Single digit
		oneDigit := int(s[i] - '0')
		if oneDigit == 0 {
			current = 0
		}

		// Two digits
		twoDigits := int((s[i-1]-'0')*10 + (s[i] - '0'))
		if twoDigits >= 10 && twoDigits <= 26 {
			current += prev
		}

		prev, current = current, prev+current
	}

	return current
}

//does not account for possible decodings with multiple twoDigit pairings 
// func numDecodings(s string) int {
// 	n := len(s)
// 	//catch empty string exception
// 	if n == 0 {
// 		return 0
// 	}

// 	// Check for leading zeros
// 	if s[0] == '0' {
// 		return 0
// 	}

// 	count := 1

// 	for i := 2; i <= n; i++ {
// 		// Two digits
// 		twoDigits := int((s[i-2]-'0')*10 + (s[i-1] - '0'))
// 		//10 and 20 do not add to possible decodings
// 		if twoDigits >= 11 && twoDigits <= 26 && twoDigits != 20{
// 			//if the 2 digit number would leave a 0 next no possible decodings have been added 
// 			if i<n && int(s[i]-'0') == 0{
// 				continue
// 			}
// 			count++
// 		}
// 	}

// 	return count
// }

func main() {
	fmt.Println(numDecodings("12"))  // Output: 2
	fmt.Println(numDecodings("226")) // Output: 3
	fmt.Println(numDecodings("06"))  // Output: 0
}
