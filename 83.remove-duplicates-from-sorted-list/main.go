package main

import "fmt"

func main() {
	n6 := &ListNode{
		Val:  3,
		Next: nil,
	}
	n5 := &ListNode{
		Val:  3,
		Next: n6,
	}
	n4 := &ListNode{
		Val:  2,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  2,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  1,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  0,
		Next: n2,
	}
	// 0,1,2,2,3,4
	n := deleteDuplicates(n1)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	dumpy := new(ListNode)
	dumpy.Next = head
	cur, last := head, dumpy
	for cur != nil {
		if cur == head || cur.Val != last.Val {
			last.Next = cur
			last = cur
		}
		cur = cur.Next
	}
	last.Next = nil
	return dumpy.Next
}
