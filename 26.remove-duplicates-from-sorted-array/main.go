package main

func main() {
	arr := []int{1,1,2}
	println(removeDuplicates(arr))
	arr = []int{0,0,1,1,1,2,2,3,3,4}
	println(removeDuplicates(arr))
}

func removeDuplicates(nums []int) int {
	if nil == nums || len(nums) <= 0 {
		return 0
	}
	i := 0
	c := nums[0]
	for _, v := range nums {
		if v != c {
			i++
			nums[i] = v
			c = v
		}
	}
	return i+1
}