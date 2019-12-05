package main

func main() {
	n3 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n2 := &ListNode{
		Val:  2,
		Next: n3,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	l1 := n1

	n6 := &ListNode{
		Val:  4,
		Next: nil,
	}
	n5 := &ListNode{
		Val:  3,
		Next: n6,
	}
	n4 := &ListNode{
		Val:  1,
		Next: n5,
	}
	l2 := n4

	traverse(mergeTwoLists(l1, l2))
}

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var l, h *ListNode
	cl1, cl2 := l1, l2

	for ; cl1 != nil && cl2 != nil; {
		val := cl1.Val
		if cl2.Val < cl1.Val {
			val = cl2.Val
			cl2 = cl2.Next
		} else {
			cl1 = cl1.Next
		}
		if h == nil {
			h = &ListNode{
				Val:  val,
				Next: nil,
			}
			l = h
		} else {
			h.Next = &ListNode{
				Val:  val,
				Next: nil,
			}
			h = h.Next
		}
	}

	last := cl1
	if cl1 == nil {
		last = cl2
	}

	for ; last != nil; {
		val := last.Val
		last = last.Next
		h.Next = &ListNode{
			Val:  val,
			Next: nil,
		}
		h = h.Next
	}

	return l
}

func traverse(l *ListNode) {
	c := l
	for ; c != nil; c = c.Next {
		println(c.Val)
	}
}
