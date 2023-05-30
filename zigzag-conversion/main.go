package main

import "fmt"

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}
	arr := make([][]byte, numRows)

	i, j, p := 0, 0, 0
	dir := -1 //-1 for down and 1 for up

	for {
		if p >= len(s) {
			break
		}

		if arr[i] == nil {
			arr[i] = make([]byte, len(s))
		}

		arr[i][j] = s[p]

		if dir == -1 && i < numRows-1 {
			i++
		}

		if dir == 1 {
			i--
			j++
		}

		if i == numRows-1 {
			dir = 1
		}

		if i == 0 {
			dir = -1
		}

		p++
	}

	str := ""
	for _, row := range arr {
		for _, val := range row {
			if val != 0 {
				str += string(val)
			}
		}
	}
	return str
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
}
