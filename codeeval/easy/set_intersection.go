package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		lines := strings.SplitN(scanner.Text(), ";", 2)
		nums1 := strings.Split(lines[0], ",")
		nums2 := strings.Split(lines[1], ",")
		hash := make(map[string]bool)
		for _, num := range nums1 {
			hash[num] = true
		}
		var sect []string
		for _, num := range nums2 {
			if hash[num] {
				sect = append(sect, num)
			}
		}
		fmt.Println(strings.Join(sect, ","))
	}
}
