package main

import "fmt"

func main() {
	//minStack := Constructor()
	//minStack.Push(-2)
	//minStack.Push(0)
	//minStack.Push(-3)
	//fmt.Println(minStack.GetMin())
	//minStack.Pop()
	//fmt.Println(minStack.Top())
	//fmt.Println(minStack.GetMin())
	minStack := Constructor()
	minStack.Push(1)
	minStack.Push(2)
	fmt.Println(minStack.Top())
	fmt.Println(minStack.GetMin())
	minStack.Pop()
	fmt.Println(minStack.GetMin())
	fmt.Println(minStack.Top())
}

type MinStack struct {
	Stack []int
	Min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min: 0,
	}
}


func (this *MinStack) Push(x int)  {
	if len(this.Stack) == 0 {
		this.Min = x
		this.Stack = append(this.Stack, 0)
		return
	}
	min := this.Min
	if x < this.Min {
		this.Min = x
	}
	this.Stack = append(this.Stack, x - min)
}


func (this *MinStack) Pop()  {
	idx := len(this.Stack) - 1
	x := this.Stack[idx]
	if x < 0 {
		this.Min = this.Min - x
	}
	this.Stack = this.Stack[:idx]
}


func (this *MinStack) Top() int {
	idx := len(this.Stack) - 1
	x := this.Stack[idx]
	if x <= 0 {
		return this.Min
	}
	return this.Min + x
}


func (this *MinStack) GetMin() int {
	return this.Min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
