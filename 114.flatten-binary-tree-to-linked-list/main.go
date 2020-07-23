package main

import "fmt"

func main() {
	n1 := &TreeNode{
		Val: 1,
	}
	n2 := &TreeNode{
		Val: 2,
	}
	n3 := &TreeNode{
		Val: 3,
	}
	n4 := &TreeNode{
		Val: 4,
	}
	n5 := &TreeNode{
		Val: 5,
	}
	n6 := &TreeNode{
		Val: 6,
	}
	n1.Left = n2
	n2.Left = n3
	n2.Right = n4
	n1.Right = n5
	n5.Right = n6
	flatten(n1)
	inOrder(n1)
}

func inOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	inOrder(root.Left)
	inOrder(root.Right)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	l := root.Left
	r := root.Right
	root.Left = nil
	root.Right = nil
	flatten(l)
	flatten(r)
	if l != nil {
		root.Right = l
		for l.Right != nil {
			l = l.Right
		}
		l.Right = r
	} else {
		root.Right = r
	}
}
