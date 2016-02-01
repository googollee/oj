package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func splitAndReverse(line string) []string {
	strs := strings.Split(line, " ")
	reverse(strs)
	return strs
}

func reverse(strs []string) {
	l := len(strs)
	for i := 0; i < l/2; i++ {
		strs[i], strs[l-i-1] = strs[l-i-1], strs[i]
	}
}

func checkCycle(strs []string, from int) []string {
	i, j := 0, from
	for {
		if i == from {
			ret := strs[:from]
			reverse(ret)
			return ret
		}
		if j >= len(strs) {
			return nil
		}
		if strs[i] != strs[j] {
			return nil
		}
		i++
		j++
	}
}

func findCycle(strs []string) []string {
	for j := 1; j < len(strs); j++ {
		if strs[j] == strs[0] {
			if ret := checkCycle(strs, j); ret != nil {
				return ret
			}
		}
	}
	return strs[0:1]
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := splitAndReverse(scanner.Text())
		cycle := findCycle(strs)
		fmt.Println(strings.Join(cycle, " "))
	}
}
