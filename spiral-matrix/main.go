package main

import "fmt"


func spiralOrder(matrix [][]int) []int {

	var ans []int

	touchrx := len(matrix) - 1
	touchlx := 1
	touchry := len(matrix[0]) - 1
	touchly := 0

	//can set 1 forward or -1 backwards or 0 stopping
	dirx := 0
	//can set 1 down or -1 up or 0 stopping
	diry := 1

	i := 0
	j := 0
	k := 0
	
	for {

		ans = append(ans, matrix[i][j])
		
		if j == touchry && diry == 1{
			diry = 0
			dirx = 1
			touchry--
		}

		if i == touchrx && dirx == 1{
			diry = -1
			dirx = 0
			touchrx--
		} 

		if j == touchly && diry == -1{
			diry = 0
			dirx = -1
			touchly++
		} 

		if i == touchlx && dirx == -1{
			diry = 1
			dirx = 0
			touchlx++
		}

		if dirx == 1 {
			i++
		} else if dirx == -1 {
			i--
		}

		if diry == 1 {
			j++
		} else if diry == -1 {
			j--
		}

		
		k++
		if k == len(matrix) * len(matrix[0]) {
			break
		}
	}
	
	return ans
}

func main() {
	fmt.Println(spiralOrder([][]int{{1,2,3},{4,5,6},{7,8,9}}))
	//[1,2,3,6,9,8,7,4,5]
	fmt.Println(spiralOrder([][]int{{1,2,3,4},{5,6,7,8},{9,10,11,12}}))
	//[1,2,3,4,8,12,11,10,9,5,6,7]
}