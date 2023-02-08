package main

import (
	"fmt"
	"sort"
)

func main() {
	//intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals := [][]int{{1, 4}, {4, 5}}
	fmt.Println(merge(intervals))
}

type Intervals [][]int

func (is Intervals) Len() int {
	return len(is)
}

func (is Intervals) Less(i, j int) bool {
	return is[i][0] < is[j][0]
}

func (is Intervals) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func merge(intervals [][]int) [][]int {
	sort.Sort(Intervals(intervals))
	count := len(intervals)
	var temp []int
	var res [][]int
	for i := 0; i < count; i++ {
		if len(temp) == 0 {
			temp = append(temp, intervals[i][0])
			temp = append(temp, intervals[i][1])
		} else {
			if intervals[i][0] > temp[1] {
				res = append(res, temp)
				temp = make([]int, 0)
				temp = append(temp, intervals[i][0])
				temp = append(temp, intervals[i][1])
			} else if intervals[i][1] > temp[1] {
				temp[1] = intervals[i][1]
			}
		}
		if i == count-1 {
			res = append(res, temp)
		}
	}
	return res
}
