package main

func main() {
	arr := []int{3,2,2,3}
	val := 3
	println(removeElement(arr, val))
	arr = []int{0,1,2,2,3,0,4,2}
	val = 2
	println(removeElement(arr, val))
}

func removeElement(nums []int, val int) int {
	if nil == nums || len(nums) == 0 {
		return 0
	}
	i := 0
	for _, v := range nums {
		if v != val {
			nums[i] = v
			i++
		}
	}
	return i
}
