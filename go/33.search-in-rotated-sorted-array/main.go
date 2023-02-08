package main

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	println(search(arr, 0))
	println(search(arr, 3))
	arr = []int{5, 1, 3}
	println(search(arr, 5))
	arr = []int{3, 1}
	println(search(arr, 1))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, i, j, target int) int {
	if j < i {
		return -1
	}
	mid := (i + j) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[i] == target {
		return i
	}
	if nums[j] == target {
		return j
	}

	if nums[mid] > nums[i] {
		if target < nums[mid] && target >= nums[i] {
			return binarySearch(nums, i, mid-1, target)
		} else {
			return binarySearch(nums, mid+1, j, target)
		}
	} else {
		if target > nums[mid] && target <= nums[j] {
			return binarySearch(nums, mid+1, j, target)
		} else {
			return binarySearch(nums, i, mid-1, target)
		}
	}
	return -1
}
