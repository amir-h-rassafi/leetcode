package main

import (
	"fmt"
)

type DataStream struct {
	value int
	k     int
	cnt   int
}

func Constructor(value int, k int) DataStream {
	return DataStream{value, k, 0}
}

func (this *DataStream) Consec(num int) bool {
	if num == this.value {
		this.cnt++
	} else {
		this.cnt = 0
	}
	if this.k <= this.cnt {
		return true
	}
	return false
}

func main() {
	fmt.Println("salam")
	var test DataStream = Constructor(4, 3)
	fmt.Println(test.Consec(4))
	fmt.Println(test.Consec(4))
	fmt.Println(test.Consec(4))
	fmt.Println(test.Consec(3))
}
