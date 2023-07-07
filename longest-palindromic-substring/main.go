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

func findLongestPalindromR(s string, cache *map[string]string) string {

	if _, ok := (*cache)[s]; ok {
		return (*cache)[s]
	}

	if isPalindrom(s) {
		(*cache)[s] = s
		return s
	}

	left := findLongestPalindromR(s[1:], cache)

	right := findLongestPalindromR(s[0:len(s)-1], cache)

	// (*cache)[s[1:]] = left

	// (*cache)[s[0:len(s)-1]] = right

	if len(left) > len(right) {
		(*cache)[s] = left
		return left
	} else {
		(*cache)[s] = right
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
	cache := make(map[string]string)
	return findLongestPalindromR(s, &cache)
}

func main() {
	fmt.Println(longestPalindrome("fdal;kfjasl;kjf;aljksdfa;ksjfj;dsajfdaaaa"))
}
