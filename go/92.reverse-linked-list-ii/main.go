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
	n := reverseBetween(n1, 2, 4)
	for n != nil {
		fmt.Println(n.Val)
		n = n.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	dummy.Next = head
	pre, last, cur := dummy, head, head
	idx := 1
	for cur != nil {
		tmp := cur.Next
		if idx == m {
			last = cur
		} else if idx > m && idx <= n {
			cur.Next = pre.Next
			pre.Next = cur
			if idx == n {
				last.Next = tmp
				break
			}
		} else {
			pre = cur
		}
		idx++
		cur = tmp
	}
	return dummy.Next
}
