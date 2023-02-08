package main

import "fmt"

func main() {
	//n7 := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	n6 := &ListNode{
		Val:  4,
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
	dumy := new(ListNode)
	dumy.Next = head
	ll, last, cur := dumy, dumy, head
	occu := false
	for cur != nil {
		if cur.Val != last.Val {
			if ll != last {
				if occu == true {
					ll.Next = cur
					occu = false
				} else {
					ll = last
				}
			}
		} else {
			if cur != head {
				occu = true
			}
		}
		last = cur
		cur = cur.Next
		if occu == true && cur == nil {
			ll.Next = nil
		}
	}
	return dumy.Next
}
