package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) getVal() int {
	return this.Val
}

func isSymmetric(root *TreeNode) bool {

	return check(root.Left, root.Right)
}

func check(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) {
		return false
	}
	if left.getVal() != right.getVal() {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)

}

func main() {
	left := TreeNode{2, nil, nil}
	right := TreeNode{1, nil, nil}
	root := TreeNode{Val: 2, Left: &left, Right: &right}
	fmt.Println(isSymmetric(&root))
}
