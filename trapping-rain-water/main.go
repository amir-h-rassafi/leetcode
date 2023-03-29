package main

import (
	"fmt"
)

type Pair struct {
	k int
	v int
}

type Center struct {
	mid Pair
	left Pair
	right Pair
}

func min(i int, j int) int {
	if i < j {
		return i
	}
	return j
}


func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func cal(i int, j int, cols []int, m int) int {

	if m == 0 {
		return 0
	}

	sum := 0
	for n:=i; n <= j; n ++ {
		sum += max(m - cols[n], 0)
	}
	
	return sum
}

func getMax(arr []int) Pair {

	if len(arr) == 0 {
		return Pair{0, 0}
	}

	max := 0
	maxi := 0

	for index, element := range arr {
		if element > max {
			max = element
			maxi = index
		}
	}

	return Pair{maxi, max}
}

func getLeftMidRight(arr []int) Center {
	mid := getMax(arr)
	left := getMax(arr[:mid.k])
	right := getMax(arr[mid.k+1:])
	right.k = mid.k + right.k + 1
	return Center{mid, left, right}
}

func trap(height []int) int {

	if len(height) <=2 {
		return 0
	}

	c := getLeftMidRight(height)
	left_rain := cal(c.left.k, c.mid.k, height, min(c.left.v, c.mid.v))
	right_rain := cal(c.mid.k, c.right.k, height, min(c.right.v, c.mid.v))

	return left_rain + right_rain + trap(height[:c.left.k+1]) + trap(height[c.right.k:])
}


func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
	fmt.Println(trap([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}))
	fmt.Println(trap([]int{15,14,13,12,11,10,9,8,7}))
}