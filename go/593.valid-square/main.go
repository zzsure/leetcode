package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := []int{0,0}
	p2 := []int{1,1}
	p3 := []int{1,0}
	p4 := []int{0,1}
	fmt.Println(validSquare(p1, p2, p3, p4))
}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	v := []int{dist(p1, p2), dist(p2, p3), dist(p3, p4), dist(p4, p1), dist(p1, p3), dist(p2, p4)}
	fmt.Println(v)
	sort.Ints(v)
	fmt.Println(v)
	return v[0] != 0 && v[0] == v[2] && v[1] == v[3] && v[4] == v[5] && 2 * v[0] == v[4]
}

func dist(p1[]int, p2[]int) int {
	return (p2[0] - p1[0]) * (p2[0] - p1[0]) + (p2[1] - p1[1]) * (p2[1] - p1[1])
}