package main

import "fmt"

func reverse(x int32) int32 {
	if x == 0 {
		return 0
	}
	minus := x < 0
	if minus {
		x = -x
	}
	ret := int32(0)
	for x > 0 {
		if ret > 0 && (1<<31-1)/ret < 10 {
			return 0
		}
		ret = ret*10 + x%10
		x = x / 10
		if ret < 0 {
			return 0
		}
	}
	if minus {
		ret = -ret
	}
	return ret
}

func main() {
	tests := []struct {
		input  int32
		output int32
	}{
		{123, 321},
		{-123, -321},
		{0, 0},
		{1, 1},
		{-1, -1},
		{1534236469, 0},
		{100, 1},
		{1563847412, 0},
	}
	for _, test := range tests {
		ret := reverse(test.input)
		fmt.Println(ret, test.output)
	}
}
