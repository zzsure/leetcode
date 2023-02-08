package main

import "fmt"

func main() {
	fmt.Println(len(permute([]int{1, 2, 3})))
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	n := nums[0]
	nums = nums[1:]
	perms := permute(nums)
	var res [][]int
	for _, perm := range perms {
		for i := 0; i < len(perm)+1; i++ {
			r := make([]int, len(perm)+1)
			r[i] = n
			isAdd := 0
			for j, p := range perm {
				if i == j {
					isAdd += 1
				}
				r[j+isAdd] = p
			}
			res = append(res, r)
		}
	}
	return res
}
