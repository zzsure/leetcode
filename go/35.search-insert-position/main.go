package main

func main() {
	nums := []int{1, 3, 5, 6}
	println(searchInsert(nums, 5))
	println(searchInsert(nums, 2))
	println(searchInsert(nums, 7))
	println(searchInsert(nums, 0))
}

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	idx := search(nums, target, 0, len(nums)-1)
	if idx == -1 {
		return 0
	}
	if idx == len(nums) {
		return len(nums) - 1
	}
	if nums[idx] == target {
		return idx
	}
	if nums[idx] > target {
		return idx
	}
	return idx + 1
}

func search(nums []int, target, i, j int) int {
	if i > j {
		return j
	}
	mid := (i + j) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[mid] > target {
		return search(nums, target, i, mid-1)
	}
	return search(nums, target, mid+1, j)
}
