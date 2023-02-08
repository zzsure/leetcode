package main

import "fmt"

func main() {
	//n5 := &ListNode{
	//	Val:  5,
	//	Next: nil,
	//}
	//n4 := &ListNode{
	//	Val:  4,
	//	Next: n5,
	//}
	//n3 := &ListNode{
	//	Val:  3,
	//	Next: n4,
	//}
	//n2 := &ListNode{
	//	Val:  2,
	//	Next: n3,
	//}
	//n1 := &ListNode{
	//	Val:  1,
	//	Next: n2,
	//}
	//head := n1
	//head = rotateRight(head, 6)
	//n3 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//n2 := &ListNode{
	//	Val:  1,
	//	Next: n3,
	//}
	//n1 := &ListNode{
	//	Val:  0,
	//	Next: n2,
	//}
	n1 := &ListNode{
		Val:  1,
		Next: nil,
	}
	//n2 := &ListNode{
	//	Val:  2,
	//	Next: nil,
	//}
	//n1 := &ListNode{
	//	Val:  1,
	//	Next: n2,
	//}
	head := n1
	head = rotateRight(head, 0)
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	if k == 0 {
		return head
	}
	sentry := new(ListNode)
	sentry.Next = head
	i, j := 0, 0
	p1, p2 := head, sentry
	for ; p1 != nil; p1 = p1.Next {
		i++
		if i-j > k {
			j++
			p2 = p2.Next
		}
		if p1.Next == nil {
			if k%i == 0 {
				break
			}
			if k > i {
				k = i - k%i
				for j := 0; j < k; j++ {
					p2 = p2.Next
				}
			}
			p1.Next = head
			break
		}
	}
	res := p2.Next
	p2.Next = nil
	return res
}
