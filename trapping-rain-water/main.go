package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	k int
	v int
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

	sum := 0
	for n:=i; n <= j; n ++ {
		sum += max(m - cols[n], 0)
	}
	
	return sum
}

func getSorted(arr []int) []Pair {
	var cols []Pair

	for index, element := range arr {
		cols = append(cols, Pair{index, element})
	}

	sort.Slice(cols, func(i, j int) bool { return cols[i].v > cols[j].v })

	return cols
}


func trap(height []int) int {

	if len(height) <= 2 {
		return 0
	}

	sorted := getSorted(height)

	mid := sorted[0]
	l := false
	r := false
	first_left := Pair{0, height[0]}
	first_right := Pair{len(height)-1, height[len(height) - 1]}
	for _, p := range sorted {
		if p.k > mid.k && r==false {
			first_right = p
			r = true
		}

		if p.k < mid.k && l==false {
			first_left = p
			l = true
		}

		if r && l {
			break
		}
	}

	left_rain := cal(first_left.k, mid.k, height, min(first_left.v, mid.v))
	right_rain := cal(mid.k, first_right.k, height, min(first_right.v, mid.v))
	return left_rain + right_rain + trap(height[:first_left.k + 1]) + + trap(height[first_right.k:])
}


func main() {
	fmt.Println(trap([]int{0,1,0,2,1,0,1,3,2,1,2,1}))
	fmt.Println(trap([]int{4,2,0,3,2,5}))
}