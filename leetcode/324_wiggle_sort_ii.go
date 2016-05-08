package main

import "fmt"

func wiggleSort(nums []int) {
	for i, j := 0, 1; j < len(nums); {
		for nums[i] == nums[j] {
			j++
		}
		if i%2 == 0 {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		} else {
			if nums[i] < nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		i++
		for j <= i {
			j++
		}
	}
}

func main() {
	tests := []struct {
		input  []int
		output []int
	}{
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 4, 1, 5, 1, 6}},
		{[]int{1, 3, 2, 2, 3, 1}, []int{2, 3, 1, 3, 1, 2}},
	}
	for _, test := range tests {
		wiggleSort(test.input)
		fmt.Println(test.input, test.output)
	}
}
