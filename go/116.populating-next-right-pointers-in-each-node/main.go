package main

func main() {
	connect(nil)
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	leftNode := root
	for leftNode.Left != nil {
		head := leftNode
		for head != nil {
			head.Left.Next = head.Right
			if head.Next != nil {
				head.Right.Next = head.Next.Left
			}
			head = head.Next
		}
		leftNode = leftNode.Left
	}
	return root
}
