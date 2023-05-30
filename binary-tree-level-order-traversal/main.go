package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type TreeNodeLevel struct {
	node  *TreeNode
	level int
}

func levelOrder(root *TreeNode) [][]int {
	rootL := TreeNodeLevel{root, 0}
	var ans [][]int
	ans = append(ans, []int{root.Val})
	BFS(&rootL, &ans)
	return ans
}

func BFS(root *TreeNodeLevel, ans *[][]int) {
	queue := make(chan *TreeNodeLevel, 100)
	queue <- root
	var nodePointer *TreeNodeLevel
	var left, right *TreeNode
	for {
		nodePointer = <-queue
		left = nodePointer.node.Left
		right = nodePointer.node.Right
		if left == nil && right == nil && len(queue) == 0 {
			break
		}

		if left != nil {
			if len(*ans) < nodePointer.level+2 {
				*ans = append(*ans, []int{left.Val})
			} else {
				(*ans)[nodePointer.level+1] = append((*ans)[nodePointer.level+1], left.Val)
			}

			queue <- &TreeNodeLevel{left, nodePointer.level + 1}
		}

		if right != nil {
			if len(*ans) < nodePointer.level+2 {
				*ans = append(*ans, []int{right.Val})
			} else {
				(*ans)[nodePointer.level+1] = append((*ans)[nodePointer.level+1], right.Val)
			}
			queue <- &TreeNodeLevel{right, nodePointer.level + 1}
		}
	}
}

func main() {
	n7 := TreeNode{7, nil, nil}
	n15 := TreeNode{15, nil, nil}
	n20 := TreeNode{20, &n15, &n7}
	n9 := TreeNode{9, nil, nil}
	root := TreeNode{3, &n9, &n20}
	fmt.Println("salam")
	fmt.Println(levelOrder(&root))
}
