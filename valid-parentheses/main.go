package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := make([]rune, 0)

	for _, char := range s {
		switch char {
		case '(', '{', '[':
			stack = append(stack, char)
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}

			top := stack[len(stack)-1]
			if (char == ')' && top != '(') || (char == '}' && top != '{') || (char == ']' && top != '[') {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({})"))
}
