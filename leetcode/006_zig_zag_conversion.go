package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	size := 2*numRows - 2
	data := make([][]rune, numRows)
	for i := range data {
		data[i] = make([]rune, 0, len(s))
	}
	for i, r := range s {
		p := i % size
		if p >= numRows {
			p = size - p
		}
		data[p] = append(data[p], r)
	}
	ret := make([]rune, 0, len(s))
	for _, r := range data {
		ret = append(ret, r...)
	}
	return string(ret)
}

func main() {
	tests := []struct {
		input  string
		row    int
		output string
	}{
		{"", 3, ""},
		{"A", 1, "A"},
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"P", 3, "P"},
		{"1234567", 5, "1237465"},
		{"0123456789abcd", 4, "06c157bd248a39"},
		{"01234", 2, "02413"},
	}
	for _, t := range tests {
		out := convert(t.input, t.row)
		if out != t.output {
			fmt.Println("failed:", t.input, t.output, out)
		}
	}
}
