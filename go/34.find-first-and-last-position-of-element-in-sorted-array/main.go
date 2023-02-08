package main

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	//println(binarySearch(nums, 0, len(nums)-1, 8))
	res := searchRange(nums, 8)
	println(res[0])
	println(res[1])
	res = searchRange(nums, 6)
	println(res[0])
	println(res[1])
}

func searchRange(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen == 0 {
		return []int{-1, -1}
	}
	if numsLen == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	mid := binarySearch(nums, 0, numsLen-1, target)
	if mid == -1 {
		return []int{-1, -1}
	}
	min, max := mid, mid
	for idx := min; idx != -1; {
		idx = binarySearch(nums, 0, idx-1, target)
		if idx != -1 {
			min = idx
		}
	}
	for idx := max; idx != -1; {
		idx = binarySearch(nums, idx+1, numsLen-1, target)
		if idx != -1 {
			max = idx
		}
	}
	return []int{min, max}
}

func binarySearch(nums []int, i, j, target int) int {
	if i > j {
		return -1
	}
	mid := (i + j) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return binarySearch(nums, i, mid-1, target)
	}
	return binarySearch(nums, mid+1, j, target)
}
