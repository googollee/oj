package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		splits := strings.Split(scanner.Text(), ",")
		if len(splits) < 3 {
			log.Fatalf("invalid line: %s", scanner.Text())
			continue
		}
		ints := make([]uint, len(splits))
		for i, str := range splits {
			i64, err := strconv.ParseUint(str, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			ints[i] = uint(i64)
		}
		num, p1, p2 := ints[0], ints[1], ints[2]
		fmt.Println(bit(num, p1) == bit(num, p2))
	}
}

func bit(num, position uint) uint {
	num = num >> (position - 1)
	return num & 0x1
}
