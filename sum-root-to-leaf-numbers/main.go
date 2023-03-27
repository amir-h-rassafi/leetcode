package main

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

var d = TreeNode{4, nil, nil}
var a = TreeNode{1, &d, nil}
var b = TreeNode{3, nil, nil}
var c = TreeNode{2, &a, &b}

func calc(root *TreeNode, val int) int {
  if root == nil {
    return 0
  }

  if root.Left == nil && root.Right == nil {
    return val * 10 + root.Val
  }
  
  return calc(root.Left, val * 10 + root.Val) + calc(root.Right, val * 10 + root.Val)
}

func sumNumbers(root *TreeNode) int {
  return calc(root, 0)
}

func main() {

  fmt.Print(sumNumbers(&c))

}