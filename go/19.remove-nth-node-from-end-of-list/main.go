package main

import "fmt"

func main() {
	n5 := &ListNode{
		Val:  5,
		Next: nil,
	}
	n4 := &ListNode{
		Val:  4,
		Next: n5,
	}
	n3 := &ListNode{
		Val:  3,
		Next: n4,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	h := removeNthFromEnd(n1, 2)
	for p := h; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	pre := &ListNode{
		Val:  0,
		Next: head,
	}
	head = pre
	p1, p2 := head, head
	for ; p2 != nil; p2 = p2.Next {
		if i == n {
			pre = p1
			p1 = p1.Next
		} else {
			i++
		}
	}

	pre.Next = p1.Next

	return head.Next
}
