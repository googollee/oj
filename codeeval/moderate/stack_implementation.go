package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func pop(n []string) []string {
	var ret []string
	for i := 0; i < len(n); i += 2 {
		ret = append(ret, n[i])
	}
	return ret
}

func reverse(strs []string) {
	l := len(strs)
	for i := 0; i < l/2; i++ {
		strs[i], strs[l-i-1] = strs[l-i-1], strs[i]
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
		n := strings.Split(scanner.Text(), " ")
		reverse(n)
		ret := pop(n)
		fmt.Println(strings.Join(ret, " "))
	}
}
