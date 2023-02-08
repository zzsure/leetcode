package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	temp := make([]int, m)
	copy(temp, nums1[0:m])
	i, j, k := 0, 0, 0
	for i < m && j < n {
		if temp[i] < nums2[j] {
			nums1[k] = temp[i]
			i++
		} else {
			nums1[k] = nums2[j]
			j++
		}
		k++
	}
	if i == m {
		for j < n {
			nums1[k] = nums2[j]
			j++
			k++
		}
	}
	if j == n {
		for i < m {
			nums1[k] = temp[i]
			i++
			k++
		}
	}
}
