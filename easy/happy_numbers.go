package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func happyNext(n uint64) uint64 {
	ret := uint64(0)
	for n > 0 {
		i := n % 10
		ret += i * i
		n = n / 10
	}
	return ret
}

func checkHappy(n uint64) bool {
	if n == 0 {
		return false
	}
	path := make(map[uint64]struct{})
	for {
		if n == 1 {
			return true
		}
		if _, ok := path[n]; ok {
			return false
		}
		path[n] = struct{}{}
		n = happyNext(n)
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
		in := scanner.Text()
		n, err := strconv.ParseUint(in, 10, 64)
		if err != nil {
			log.Fatal("invalid input:", in)
		}
		if checkHappy(n) {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
