package main

import "fmt"

func nthUglyNumber(n int) int {
	p2 := 0
	p3 := 0
	p5 := 0
	nums := make([]int, n)
	nums[0] = 1
	for i := 1; i < n; i++ {
		n2 := nums[p2] * 2
		n3 := nums[p3] * 3
		n5 := nums[p5] * 5
		fmt.Println(nums, p2, p3, p5, n2, n3, n5)
		if n2 <= n3 && n2 <= n5 {
			nums[i] = n2
			p2++
		}
		if n3 <= n2 && n3 <= n5 {
			nums[i] = n3
			p3++
		}
		if n5 <= n2 && n5 <= n3 {
			nums[i] = n5
			p5++
		}
	}
	return nums[n-1]
}

func main() {
	tests := []struct {
		n   int
		num int
	}{
		{1, 1},
		{2, 2},
		{10, 12},
	}
	for _, test := range tests {
		ret := nthUglyNumber(test.n)
		fmt.Println(ret, test.num)
	}
}
