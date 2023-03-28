package main

import (
	"fmt"
	"math"
)

// We can add or sub any 2^x

func minOperations(n int) int {
 
	k := int(math.Log2(float64(n)))
	// so n is between k and k+1

	lowerband := 1 << uint(k)
	upperband := 1 << uint(k+1)
	mindiff := int(math.Min(float64(upperband - n), float64(n - lowerband)))

	if n==0 {
		return 0
	}

	return 1 + minOperations(mindiff)
}


func main() {
	fmt.Println(minOperations(39))
	fmt.Println(minOperations(54))
}