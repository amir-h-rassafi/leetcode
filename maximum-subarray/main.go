package main

import (
	"fmt"
)

func getMax(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func getMaxArr(arr *[]int, size int) int {
	max := -10000000000

	e := 0
	for i:=0; i<size; i++ {
		e = (*arr)[i]
		if e > max {
			max = i
		}
	}

	return max
}

func maxSubArray(nums []int) int {

	for i := 1; i < len(nums); i++ {
		nums[i] = getMax(nums[i], nums[i-1]+ nums[i])
	}

	return getMaxArr(&nums, len(nums))
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}