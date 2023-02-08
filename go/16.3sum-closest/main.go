package main

import (
	"math"
	"sort"
)

func main() {
	arr := []int{-1, 2, 1, -4}
	println(threeSumClosest(arr, 1))
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length < 3 {
		return 0
	}
	sort.Ints(nums)
	res := nums[0]+nums[1]+nums[length-1]
	for i := 0; i < length - 2; i++ {
		j := i+1
		k := length -1
		for ; j < k; {
			num := nums[i]+nums[j]+nums[k]
			if num == target {
				return target
			}
			if num < target {
				j++
			} else if num > target {
				k--
			}
			if math.Abs(float64(target-num)) < math.Abs(float64(target-res)) {
				res = num
			}
		}
	}
	return res
}
