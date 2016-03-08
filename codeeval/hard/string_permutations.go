package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type runes []rune

func isDigital(r rune) bool {
	return '0' <= r && r <= '9'
}

func isUpper(r rune) bool {
	return 'A' <= r && r <= 'Z'
}

func (r runes) Len() int      { return len(r) }
func (r runes) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool {
	if isDigital(r[i]) {
		if isDigital(r[j]) {
			return r[i] < r[j]
		}
		return true
	}
	if isUpper(r[i]) {
		if isDigital(r[j]) {
			return false
		}
		if isUpper(r[j]) {
			return r[i] < r[j]
		}
		return true
	}
	if isDigital(r[j]) {
		return false
	}
	if isUpper(r[j]) {
		return false
	}
	return r[i] < r[j]
}

func arrange(str string, start int, output *[]string) {
	if start == len(str) {
		*output = append(*output, str)
		return
	}
	runes := make([]rune, len(str))
	for i, r := range str {
		runes[i] = r
	}
	for i := start; i < len(runes); i++ {
		runes[start], runes[i] = runes[i], runes[start]
		s := string(runes)
		arrange(s, start+1, output)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		rs := make(runes, len(str))
		for i, r := range str {
			rs[i] = r
		}
		sort.Sort(rs)
		output := []string{}
		arrange(string(rs), 0, &output)
		fmt.Println(strings.Join(output, ","))
	}
}
