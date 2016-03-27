package main

import "fmt"

func checkPalindrome(b []byte) bool {
	i := 0
	j := len(b) - 1
	for i < j {
		if b[i] != b[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func longestPalindrome(s string) string {
	b := []byte(s)
	l := len(s)
	if l == 0 {
		return ""
	}
	ret := b[0:1]
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			sub := b[i : j+1]
			if checkPalindrome(sub) {
				if len(sub) > len(ret) {
					ret = sub
				}
			}
		}
	}
	return string(ret)
}

func main() {
	tests := []struct {
		input  string
		output string
	}{
		{"", ""},
		{"abcdedcba", "abcdedcba"},
		{"aa", "aa"},
		{"abcde", "a"},
		{"abccde", "cc"},
		{"abcdcde", "cdc"},
		{"abcdcdedc", "cdedc"},
	}
	for _, t := range tests {
		out := longestPalindrome(t.input)
		if out != t.output {
			fmt.Println("failed:", t.input, t.output, out)
		}
	}
}
