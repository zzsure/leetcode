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
	for leftNode != nil {
		cur := leftNode
		var pre *Node
		leftNode = nil
		for cur != nil {
			// process left child
			if cur.Left != nil {
				if pre == nil {
					pre, leftNode = cur.Left, cur.Left
				} else {
					pre.Next = cur.Left
					pre = cur.Left
				}
			}
			// process right child
			if cur.Right != nil {
				if pre == nil {
					pre, leftNode = cur.Right, cur.Right
				} else {
					pre.Next = cur.Right
					pre = cur.Right
				}
			}
			// set leftNode for the next level
			cur = cur.Next
		}
	}
	return root
}
