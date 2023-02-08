package main

func main() {
	//nums := []int{}
	nums := []int{1, 1}
	//nums := []int{1, 2, 3, 1}
	//nums := []int{2, 7, 9, 3, 1}
	println(rob(nums))
}

// bp[i] = math.max(num[i] + bp[i-2], bp[i-1])
func rob(nums []int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}
	bp := make([]int, len(nums))
	for i, val := range nums {
		if i == 0 {
			bp[i] = val
		} else if i == 1 {
			if val > bp[0] {
				bp[i] = val
			} else {
				bp[i] = bp[0]
			}
		} else {
			bp[i] = getMax(val + bp[i - 2], bp[i - 1])
		}
	}
	return bp[len(nums) - 1]
}

func getMax(v1, v2 int) int {
	if v1 > v2 {
		return v1
	}
	return v2
}
