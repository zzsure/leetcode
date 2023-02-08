package main

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	output := twoSum(numbers, target)
	for _, out := range output {
		println(out)
	}
}

func twoSum(numbers []int, target int) []int {
	mapValue := make(map[int]int)
	for idx, val := range numbers {
		if match, ok := mapValue[target - val]; ok {
			return []int{match + 1, idx + 1}
		}
		mapValue[val] = idx
	}
	return []int{0, 0}
}
