package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

type lineHeap []string

func (h lineHeap) Len() int           { return len(h) }
func (h lineHeap) Less(i, j int) bool { return len(h[i]) < len(h[j]) }
func (h lineHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *lineHeap) Push(x interface{}) {
	*h = append(*h, x.(string))
}

func (h *lineHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("read number of lines error:", scanner.Err())
	}
	numberStr := scanner.Text()
	n, err := strconv.ParseInt(numberStr, 10, 64)
	if err != nil {
		log.Fatalf("number of lines(%s) error: %s", numberStr, err)
	}
	var h lineHeap
	for scanner.Scan() {
		heap.Push(&h, scanner.Text())
		for h.Len() > int(n) {
			heap.Pop(&h)
		}
	}
	sort.Sort(sort.Reverse(h))
	for _, s := range h {
		fmt.Println(s)
	}
}
