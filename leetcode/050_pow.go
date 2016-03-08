package main

import "fmt"

func multiMap(x float64, n uint) []float64 {
	var ret []float64
	for n > 0 {
		if n&0x1 == 1 {
			ret = append(ret, x)
		}
		x *= x
		n = n >> 1
	}
	return ret
}

func myPow(x float64, n int) float64 {
	dim := false
	if n < 0 {
		dim = true
		n = -n
	}
	factors := multiMap(x, uint(n))
	ret := float64(1)
	for _, f := range factors {
		ret *= f
	}
	if dim {
		ret = 1 / ret
	}
	return ret
}

func main() {
	fmt.Println(myPow(8.84372, -5))
}
