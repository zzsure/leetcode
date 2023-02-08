package main

import "fmt"

func main() {
	n6 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n5 := &ListNode{
		Val:  5,
		Next: n6,
	}
	n4 := &ListNode{
		Val:  2,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  4,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	// 0,1,2,2,3,4
	n := partition(n1, 3)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	left := new(ListNode)
	right := new(ListNode)
	cur, leftCur, rightCur := head, left, right

	for cur != nil {
		if cur.Val < x {
			leftCur.Next = cur
			leftCur = cur
		} else {
			rightCur.Next = cur
			rightCur = cur
		}
		cur = cur.Next
	}
	rightCur.Next = nil
	leftCur.Next = right.Next

	return left.Next
}
