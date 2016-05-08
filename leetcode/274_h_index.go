package main

import (
	"fmt"
	"sort"
)

func hIndex(citations []int) int {
	l := len(citations)
	if l == 0 {
		return 0
	}
	sort.Ints(citations)
	for i := 0; i < l/2; i++ {
		citations[i], citations[l-i-1] = citations[l-i-1], citations[i]
	}
	if citations[l-1] >= l {
		return l
	}
	for i := range citations {
		if citations[i] <= i {
			return i
		}
	}
	return 0
}

func main() {
	tests := []struct {
		citations []int
		hIndex    int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 1, 1}, 1},
		{[]int{3, 1}, 1},
		{[]int{3, 2}, 2},
		{[]int{3, 3}, 2},
		{[]int{3, 3, 3}, 3},
		{[]int{3, 3, 3, 1}, 3},
		{[]int{3, 3, 3, 3}, 3},
		{[]int{3, 0, 6, 1, 5}, 3},
	}
	for _, test := range tests {
		ret := hIndex(test.citations)
		fmt.Println(ret, test.hIndex)
	}
}
