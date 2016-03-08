package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	ret := 0
	strBytes := []byte(s)
	for i := range strBytes {
		j := i + 1
		marks := make(map[byte]struct{})
		marks[strBytes[i]] = struct{}{}
		for ; j < len(strBytes); j++ {
			if _, ok := marks[strBytes[j]]; ok {
				break
			}
			marks[strBytes[j]] = struct{}{}
		}
		fmt.Println(string(strBytes[i:j]), i, j)
		if j-i > ret {
			ret = j - i
		}
	}
	return ret
}

func main() {
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
}
