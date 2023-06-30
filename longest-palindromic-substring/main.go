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

func longestPalindrome(s string) string {

	if isPalindrom(s) {
		return s
	}

	left := longestPalindrome(s[1:])
	right := longestPalindrome(s[0 : len(s)-1])

	if len(left) > len(right) {
		return left
	} else {
		return right
	}

}

func main() {
	fmt.Println(longestPalindrome("babaddtattarrattatddetartrateedredividerb"))
}
