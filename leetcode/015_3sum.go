package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	l := len(nums)
	ret := make(map[[3]int]struct{})
	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		target := 0 - nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {
				ret[[3]int{nums[i], nums[j], nums[k]}] = struct{}{}
				j++
				k--
				continue
			}
			if sum > target {
				k--
			} else {
				j++
			}
		}
	}
	r := make([][]int, 0, len(ret))
	for k := range ret {
		r = append(r, []int{k[0], k[1], k[2]})
	}
	return r
}

func main() {
	tests := []struct {
		nums []int
		sums [][]int
	}{
		{[]int{}, [][]int{}},
		{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, 0, 1}, {-1, -1, 2}}},
		{[]int{1, 1, 2, 2, -3, -3}, [][]int{{-3, 1, 2}}},
	}
	for _, test := range tests {
		ret := threeSum(test.nums)
		fmt.Println(ret, test.sums)
	}
}
