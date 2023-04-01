package main

import "fmt"

//todo handle it with string manipulation

type ListNode struct {
	Val int
	Next *ListNode
}

func getNum(n *ListNode) int {
	num := 0
	i := 1
	for {
		num = i * n.Val + num
		i *= 10
		if n.Next == nil {
			break
		}
		n = n.Next
	}
	return num
}

func genLL(num int) *ListNode {
	remain := 0
	var n *ListNode
	var prev *ListNode
	var first *ListNode
	for {
		remain = num % 10
		num = num / 10
		n = &ListNode{remain, nil}
		if first == nil {
			first = n
		}
		if prev != nil {
			prev.Next = n
		}
		prev = n

		if num == 0 {
			break
		}
	}

	return first
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return genLL(getNum(l1) + getNum(l2))
}


func main() {
	//243
	n3 := ListNode{3, nil}
	n2 := ListNode{4, &n3}
	n1 := ListNode{2, &n2}
	//564
	m3 := ListNode{4, nil}
	m2 := ListNode{6, &m3}
	m1 := ListNode{5, &m2}

	fmt.Println(getNum(addTwoNumbers(&n1, &m1)))

}