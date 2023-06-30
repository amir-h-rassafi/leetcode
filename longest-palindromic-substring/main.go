package main

import "fmt"

// calculate palindrom wth o(n)/2
func isPalindrom(s string) bool {

	if len(s) == 0 {
		return true
	}

	var leftPtr, RightPtr int

	leftPtr, RightPtr = 0, len(s)-1

	if leftPtr == RightPtr {
		return true
	}

	for {
		if s[leftPtr] != s[RightPtr] {
			return false
		}

		if leftPtr >= RightPtr {
			break
		}

		leftPtr++
		RightPtr--
	}

	return true
}

func findLongestPalindromR(s string, cache *map[string]bool) string {

	if _, ok := (*cache)[s]; ok {
		return s
	}

	if isPalindrom(s) {
		(*cache)[s] = true
		return s
	}

	left := findLongestPalindromR(s[1:], cache)

	right := findLongestPalindromR(s[0:len(s)-1], cache)

	if len(left) > len(right) {
		return left
	} else {
		return right
	}
}

// sliding window O(n^3)
func findLongestPalindrom(s string) string {

	for windowSize := len(s); windowSize > 0; windowSize-- {
		for startPtr := 0; startPtr <= len(s)-windowSize; startPtr++ {
			if isPalindrom(s[startPtr : startPtr+windowSize]) {
				return s[startPtr : startPtr+windowSize]
			}
		}
	}

	return ""
}

func longestPalindrome(s string) string {
	return findLongestPalindrom(s)
}

func main() {
	fmt.Println(longestPalindrome("abaab"))
}
