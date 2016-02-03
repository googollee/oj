package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Table [][]int

func NewTable(x, y int) Table {
	ret := Table(make([][]int, x))
	for i := range ret {
		ret[i] = make([]int, y)
	}
	return ret
}

func (t Table) Get(x, y int) int {
	if x < 0 || x >= len(t) {
		return 0
	}
	l := t[x]
	if y < 0 || y >= len(l) {
		return 0
	}
	return l[y]
}

func (t Table) Set(x, y, i int) {
	if x < 0 || x >= len(t) {
		return
	}
	l := t[x]
	if y < 0 || y >= len(l) {
		return
	}
	l[y] = i
}

func (t Table) Fill(a, b string) {
	for x, ac := range a {
		for y, bc := range b {
			if ac == bc {
				t.Set(x, y, t.Get(x-1, y-1)+1)
			} else {
				x1y := t.Get(x-1, y)
				xy1 := t.Get(x, y-1)
				if x1y > xy1 {
					t.Set(x, y, x1y)
				} else {
					t.Set(x, y, xy1)
				}
			}
		}
	}
}

func (t Table) Result(a string) string {
	var ret []byte
	x, y := len(t)-1, len(t[0])-1
	v := t.Get(x, y)
	for x >= 0 && y >= 0 {
		x1y := t.Get(x-1, y)
		xy1 := t.Get(x, y-1)
		switch {
		case v == x1y:
			x = x - 1
		case v == xy1:
			y = y - 1
		default:
			ret = append(ret, a[x])
			x, y = x-1, y-1
			v = t.Get(x, y)
		}
	}
	l := len(ret)
	for i := 0; i < l/2; i++ {
		ret[i], ret[l-i-1] = ret[l-i-1], ret[i]
	}
	return string(ret)
}

func lcs(a, b string) string {
	tbl := NewTable(len(a), len(b))
	tbl.Fill(a, b)
	return tbl.Result(a)
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.SplitN(scanner.Text(), ";", 2)
		ret := lcs(strs[0], strs[1])
		fmt.Println(ret)
	}
}
