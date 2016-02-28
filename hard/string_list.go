package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }

func getChars(s string) []rune {
	rs := map[rune]struct{}{}
	for _, r := range s {
		rs[r] = struct{}{}
	}
	ret := make(runes, 0, len(rs))
	for r := range rs {
		ret = append(ret, r)
	}
	sort.Sort(ret)
	return ret
}

func output(runes []rune, indexes []int) {
	for _, i := range indexes {
		fmt.Printf("%c", runes[i])
	}
}

func add(in []int, max int) bool {
	for i := len(in) - 1; i >= 0; i-- {
		in[i]++
		if in[i] < max {
			return true
		}
		in[i] = 0
	}
	return false
}

func allOrder(runes []rune, n int) {
	in := make([]int, n)
	output(runes, in)
	for add(in, len(runes)) {
		fmt.Print(",")
		output(runes, in)
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
		sp := strings.Split(scanner.Text(), ",")
		n, _ := strconv.ParseInt(sp[0], 10, 64)
		rs := getChars(sp[1])
		allOrder(rs, int(n))
		fmt.Println()
	}
}
