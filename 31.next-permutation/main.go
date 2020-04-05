package main

import "fmt"

func main() {
	a1 := []int{1, 2, 3}
	nextPermutation(a1)
	printArr(a1)
	a2 := []int{3, 2, 1}
	nextPermutation(a2)
	printArr(a2)
	a3 := []int{1, 1, 5}
	nextPermutation(a3)
	printArr(a3)
	// 2,1,3
	a4 := []int{1, 3, 2}
	nextPermutation(a4)
	printArr(a4)
	// 5,1,1
	a5 := []int{1, 5, 1}
	nextPermutation(a5)
	printArr(a5)
}

func printArr(nums []int) {
	for _, n := range nums {
		fmt.Println(n)
	}
}

func nextPermutation(nums []int) {
	count := len(nums)
	if count == 0 || count == 1 {
		return
	}
	i := count - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		j := count - 1
		for ; j >= 0; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1 : count])
}

func reverse(nums []int) {
	count := len(nums)
	for i := 0; i < count/2; i++ {
		nums[i], nums[count-i-1] = nums[count-i-1], nums[i]
	}
}
