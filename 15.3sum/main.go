package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	println(res)
	for _, a := range res {
		for _, b := range a {
			println(b)
		}
	}
}

func threeSum(nums []int) [][]int {
	numLen := len(nums)
	var res [][]int
	resMap := make(map[string]bool)

	if numLen == 0 {
		return res
	}

	numMap := make(map[int]int, numLen)
	for idx, n := range nums {
		numMap[n] = idx
	}
	for i := 0; i < numLen; i++ {
		for j := i+1; j < numLen; j++ {
			negSum := -1*(nums[i]+nums[j])
			if _, ok := numMap[negSum]; ok {
				k := numMap[negSum]
				if k > i && k > j {
					n1, n2, n3 := nums[i], nums[j], nums[k]
					key := numKey(n1, n2, n3)
					if _, ok := resMap[key]; !ok {
						res = append(res, []int{n1, n2, n3})
						resMap[key] = true
						key = numKey(n1, n3, n2)
						resMap[key] = true
						key = numKey(n2, n1, n3)
						resMap[key] = true
						key = numKey(n2, n3, n1)
						resMap[key] = true
						key = numKey(n3, n1, n2)
						resMap[key] = true
						key = numKey(n3, n2, n1)
						resMap[key] = true
					}
				}
			}
		}
	}
	return res
}

func numKey(i, j, k int) string {
	return fmt.Sprintf("%d_%d_%d", i, j, k)
}