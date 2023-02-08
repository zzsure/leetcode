package main

func main() {
	arr := []int{1,8,6,2,5,4,8,3,7}
	println(maxArea(arr))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}
	i := 0
	j := len(height)-1
	max := 0
	w, h := 0, 0

	for ; i < j; {
		w = j-i
		if height[i] < height[j] {
			h = height[i]
			i++
		} else {
			h = height[j]
			j--
		}
		if w*h > max {
			max = w*h
		}
	}

	return max
}
