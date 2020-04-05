package main

func main() {
	//n4 := &ListNode{
	//	Val:  4,
	//	Next: nil,
	//}
	//n3 := &ListNode{
	//	Val:  3,
	//	Next: n4,
	//}
	n2 := &ListNode{
		Val:  2,
		Next: nil,
	}
	n1 := &ListNode{
		Val:  1,
		Next: n2,
	}
	n := swapPairs(n1)
	for p := n; p != nil; p = p.Next {
		println(p.Val)
	}

	// [2,1]
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := &ListNode{
		Val:  0,
		Next: head,
	}
	for p, p1, p2 := tmp, tmp.Next, tmp.Next.Next; p != nil && p1 != nil && p2 != nil; p, p1, p2 = p.Next.Next, p1.Next, p1.Next.Next {
		p1.Next = p2.Next
		p2.Next = p1
		p.Next = p2
		if p.Next == nil || p.Next.Next == nil || p1.Next == nil || p1.Next.Next == nil {
			break
		}
	}
	return tmp.Next
}
