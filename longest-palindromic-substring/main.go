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

func findLongestPalindrom(s string, cache *map[string]bool) string {

	if _, ok := (*cache)[s]; ok {
		return s
	}

	if isPalindrom(s) {
		(*cache)[s] = true
		return s
	}

	left := findLongestPalindrom(s[1:], cache)

	//left is priority
	if len(left) == len(s)-1 {
		return left
	}

	right := findLongestPalindrom(s[0:len(s)-1], cache)

	if len(left) > len(right) {
		return left
	} else {
		return right
	}
}

func longestPalindrome(s string) string {
	cache := make(map[string]bool)
	return findLongestPalindrom(s, &cache)
}

func main() {
	fmt.Println(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
}
