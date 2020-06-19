package main

import (
	"fmt"
	"sort"
)

func main() {
	//strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	strs := []string{"", "b"}
	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {
	objs := make([]*StrObj, len(strs))
	for i, str := range strs {
		s := SortString(str)
		objs[i] = &StrObj{
			Ori: str,
			Key: s,
		}
	}
	sort.Sort(StrObjs(objs))

	var res [][]string
	r := make([]string, 0)
	for i := 0; i < len(objs); i++ {
		if i == 0 || objs[i].Key == objs[i-1].Key {
			r = append(r, objs[i].Ori)
			continue
		}
		res = append(res, r)
		r = make([]string, 0)
		r = append(r, objs[i].Ori)
	}
	res = append(res, r)
	return res
}

type StrObj struct {
	Ori string
	Key string
}

type StrObjs []*StrObj

func (so StrObjs) Less(i, j int) bool {
	return so[i].Key < so[j].Key
}

func (so StrObjs) Swap(i, j int) {
	so[i], so[j] = so[j], so[i]
}

func (so StrObjs) Len() int {
	return len(so)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
