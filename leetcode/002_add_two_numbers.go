package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ret ListNode
	addition := 0
	result := &ret
	for l1 != nil && l2 != nil {
		result.Next = &ListNode{
			Val: l1.Val + l2.Val + addition,
		}
		result = result.Next
		addition = 0
		if result.Val >= 10 {
			addition = 1
			result.Val -= 10
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	rest := l1
	if l2 != nil {
		rest = l2
	}
	for rest != nil {
		result.Next = &ListNode{}
		result = result.Next
		result.Val = rest.Val + addition
		addition = 0
		if result.Val >= 10 {
			addition = 1
			result.Val -= result.Val
		}
		rest = rest.Next
	}
	if addition == 1 {
		result.Next = &ListNode{
			Val: 1,
		}
	}
	return ret.Next
}

func printList(l *ListNode) {
	for l != nil {
		fmt.Print(l.Val)
		l = l.Next
	}
	fmt.Println()
}

func main() {
	l1 := &ListNode{
		Val: 9,
		// Next: &ListNode{
		// 	Val: 4,
		// 	Next: &ListNode{
		// 		Val: 9,
		// 	},
		// },
	}
	l2 := &ListNode{
		Val: 9,
		// Next: &ListNode{
		// 	Val: 6,
		// 	Next: &ListNode{
		// 		Val: 4,
		// 	},
		// },
	}
	r := addTwoNumbers(l1, l2)
	printList(r)
}
