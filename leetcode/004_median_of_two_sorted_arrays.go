package main

import "fmt"

func findKth(a, b []int, k int) float64 {
	if len(a) == 0 {
		return float64(b[k-1])
	}
	if len(b) == 0 {
		return float64(a[k-1])
	}
	if k == 1 {
		if a[0] < b[0] {
			return float64(a[0])
		}
		return float64(b[0])
	}
	al := k / 2
	bl := k / 2
	if len(a) < al {
		al = len(a)
	}
	am := a[al-1]
	if len(b) < bl {
		bl = len(b)
	}
	bm := b[bl-1]
	if am > bm {
		return findKth(a, b[bl:], k-bl)
	}
	if am < bm {
		return findKth(a[al:], b, k-al)
	}
	if k-al-bl > 0 {
		return findKth(a[al:], b[bl:], k-al-bl)
	}
	return float64(am)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l%2 == 0 {
		m1 := findKth(nums1, nums2, l/2)
		m2 := findKth(nums1, nums2, l/2+1)
		return (m1 + m2) / 2
	}
	return findKth(nums1, nums2, l/2+1)
}

func main() {
	n1 := []int{1, 2}
	n2 := []int{1, 4}
	fmt.Println(findMedianSortedArrays(n1, n2))
}
