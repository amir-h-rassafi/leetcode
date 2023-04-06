package main

import "fmt"

func getMinLen(s1 string, s2 string) int {
	
	if len(s1) < len(s2) {
		return len(s1)
	}

	return len(s2)
}

func longestCommonPrefix(strs []string) string {
	maxSub := strs[0]
	sub := ""
    for _, s := range strs {
		
		sub = ""
		
		for i:=0 ; i < getMinLen(maxSub, s); i++ {
			if s[i] != maxSub[i] {
				break
			}
			sub += string(s[i])
		}

		if len(sub) < len(maxSub) {
			maxSub = sub	
		}
	}

	return maxSub
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"amir", "am", "ambbas"}))
}