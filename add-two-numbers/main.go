package main

import (
	"fmt"
	"strconv"
)

//todo handle it with string manipulation

type ListNode struct {
	Val int
	Next *ListNode
}

func getNum(n *ListNode) string {
	num := ""
	for {
		num += strconv.Itoa(n.Val)
		if n.Next == nil {
			break
		}
		n = n.Next
	}

	return num
}

func genLL(num string) *ListNode {
	i := len(num) - 1
	var n *ListNode
	var prev *ListNode
	var first *ListNode
	var err error
	var s int
	for {		
		if num[i] == 0 {
			break
		}
		s, err = strconv.Atoi(string(num[i]))
		if err != nil {
			panic(err)
		}
		n = &ListNode{s, nil}

		if first == nil {
			first = n
		}
		if prev != nil {
			prev.Next = n
		}
		prev = n
		if i == 0 {
			break
		}
		i--
	}

	return first
}

func addString(num1 string, num2 string) string {
    b1 := []byte(num1)
    b2 := []byte(num2)
    if len(b1) > len(b2) {
        b2 = append(make([]byte, len(b1)-len(b2)), b2...)
    } else {
        b1 = append(make([]byte, len(b2)-len(b1)), b1...)
    }
    res := make([]byte, len(b1)+1)
    carry := 0
    for i := len(b1) - 1; i >= 0; i-- {
        v := int(b1[i]-'0') + int(b2[i]-'0') + carry
        res[i+1] = byte(v%10) + '0'
        carry = v / 10
    }
    if res[0] == '0' {
        res = res[1:]
    }
    return string(res)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return genLL(addString(getNum(l1), getNum(l2)))
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