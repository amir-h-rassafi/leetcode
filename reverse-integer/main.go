package main

import (
	"errors"
	"fmt"
)

func reverse(x int) int {

	//2147483648 - 1, len => 10
	maxInt := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 7}
	//-2147483648
	minInt := []int{2, 1, 4, 7, 4, 8, 3, 6, 4, 8}

	revNumber := make([]int, 0, 10)

	realX := x

	var threshold *[]int
	if x > 0 {
		threshold = &maxInt
	} else {
		threshold = &minInt
		x = -x
	}

	var digit int

	for i := 0; i < 10; i++ {
		digit = x % 10
		revNumber = append(revNumber, digit)
		x = x / 10
		if x == 0 {
			break
		}
	}

	zeroSpace := make([]int, 10-len(revNumber))

	revNumber = append(zeroSpace, revNumber...)

	fmt.Println(revNumber, threshold)
	for i := 0; i < 10; i++ {
		if revNumber[i] < (*threshold)[i] {
			break
		}
		if revNumber[i] > (*threshold)[i] {
			return 0
		}
	}

	num := convert10DigitArrayToInt(&revNumber)

	if realX > 0 {
		return num
	} else {
		return -1 * num
	}
}

func convert10DigitArrayToInt(n *[]int) int {
	if len(*n) == 0 {
		errors.New("Wrong length")
	}
	k := 0
	for i := 0; i < 10; i++ {
		k += intPow(10, 9-i) * (*n)[i]
	}
	return k
}

func intPow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func main() {
	fmt.Println(reverse(-2147483412))
}
