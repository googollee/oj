package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePart(head *ListNode, end *ListNode) *ListNode {
	var p *ListNode
	c := head
	n := c
	for c != end {
		n = n.Next
		c.Next = p
		p = c
		c = n
	}
	return p
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	var p *ListNode
	c := head
	for {
		begin := c
		end := c
		for i := 0; i < k; i++ {
			if end == nil {
				return head
			}
			end = end.Next
		}
		h := reversePart(begin, end)
		if p == nil {
			head = h
		} else {
			p.Next = h
		}
		p = begin
		p.Next = end
		c = end
	}
}

func makeListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	ret := &ListNode{}
	c := ret
	for _, i := range a {
		c.Next = &ListNode{
			Val: i,
		}
		c = c.Next
	}
	return ret.Next
}

func getListNode(head *ListNode) []int {
	var ret []int
	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
	}
	return ret
}

func main() {
	tests := []struct {
		input  []int
		k      int
		output []int
	}{
		{nil, 1, nil},
		{nil, 2, nil},
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{2, 1, 4, 3, 5}},
		{[]int{1, 2, 3, 4, 5}, 3, []int{3, 2, 1, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, 4, []int{4, 3, 2, 1, 5}},
		{[]int{1, 2, 3, 4, 5}, 5, []int{5, 4, 3, 2, 1}},
	}
	for _, test := range tests {
		list := makeListNode(test.input)
		ret := reverseKGroup(list, test.k)
		a := getListNode(ret)
		fmt.Println(a, test.output)
	}
}
