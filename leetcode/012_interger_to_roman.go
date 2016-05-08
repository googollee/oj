package main

import "fmt"

func intToRoman(num int) string {
	if num < 1 || num > 3999 {
		return ""
	}
	pair := [][]byte{
		{'I', 'V'},
		{'X', 'L'},
		{'C', 'D'},
		{'M', ' '},
		{' ', ' '},
	}
	var ret []byte
	pairIndex := 0
	for num > 0 {
		index := num % 10
		romanIndex := make([]byte, 0, 4)
		one := pair[pairIndex][0]
		five := pair[pairIndex][1]
		ten := pair[pairIndex+1][0]
		gtFive := false
		if index >= 5 {
			gtFive = true
			index -= 5
			if index < 4 {
				romanIndex = append(romanIndex, five)
			}
		}
		switch index {
		case 4:
			romanIndex = append(romanIndex, one)
			if !gtFive {
				romanIndex = append(romanIndex, five)
			} else {
				romanIndex = append(romanIndex, ten)
			}
		case 3:
			romanIndex = append(romanIndex, one)
			fallthrough
		case 2:
			romanIndex = append(romanIndex, one)
			fallthrough
		case 1:
			romanIndex = append(romanIndex, one)
		}
		ret = append(romanIndex, ret...)
		num /= 10
		pairIndex++
	}
	return string(ret)
}

func main() {
	tests := []struct {
		input  int
		output string
	}{
		{0, ""},
		{4000, ""},
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{11, "XI"},
		{19, "XIX"},
		{40, "XL"},
		{50, "L"},
		{80, "LXXX"},
		{90, "XC"},
		{98, "XCVIII"},
		{99, "XCIX"},
		{100, "C"},
		{500, "D"},
		{1000, "M"},
		{3999, "MMMCMXCIX"},
	}
	for _, test := range tests {
		ret := intToRoman(test.input)
		fmt.Println(test.output, ret)
	}
}
