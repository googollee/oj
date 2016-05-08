package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	l := len(nums)
	ret := -1
	minDiff := -1
	for i := 0; i < l-2; i++ {
		j := i + 1
		k := l - 1
		t := target - nums[i]
		for j < k {
			sum := nums[j] + nums[k]
			if sum == t {
				return target
			}
			diff := sum - t
			if diff < 0 {
				diff = -diff
			}
			if minDiff < 0 || minDiff > diff {
				minDiff = diff
				ret = nums[i] + nums[j] + nums[k]
			}
			if sum > t {
				k--
			} else {
				j++
			}
		}
	}
	return ret
}

func main() {
	tests := []struct {
		nums   []int
		target int
		output int
	}{
		{[]int{}, 0, -1},
		{[]int{0, 0, 0}, 2, 0},
		{[]int{-1, 2, 1, -4}, 1, 2},
	}
	for _, test := range tests {
		ret := threeSumClosest(test.nums, test.target)
		fmt.Println(ret, test.output)
	}
}
