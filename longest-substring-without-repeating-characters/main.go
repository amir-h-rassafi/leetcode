package main

import (
	"fmt"
)

func getMax(i int, j int) int {
	if i < j {
		return j
	}
	return i
}

func lengthOfLongestSubstring(ss string) int {
	ans, r, l := 0, 0, 0
	p := make(map[byte]int, 255)

	for {

		if r >= len(ss) {
			break
		}

		if value, ok := p[ss[r]]; ok {
			l = getMax(value+1, l)
		}

		p[ss[r]] = r
		ans = getMax(ans, r-l+1)
		r++
	}

	return ans
}

func main() {
	fmt.Println("Let's go")
	fmt.Println(lengthOfLongestSubstring("abacde"))
}
