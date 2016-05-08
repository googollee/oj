package main

import "fmt"

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	ret := 0
	for i < j {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > ret {
			ret = area
		}
	}
	return ret
}

func main() {
	tests := []struct {
		height []int
		output int
	}{
		{[]int{1, 1}, 1},
		{[]int{2, 1, 1}, 2},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 1, 2}, 4},
		{[]int{2, 2, 2}, 4},
		{[]int{4, 3, 5, 2, 6}, 16},
		{[]int{1, 2, 4, 3}, 4},
	}
	for _, test := range tests {
		ret := maxArea(test.height)
		fmt.Println(ret, test.output)
	}
}
