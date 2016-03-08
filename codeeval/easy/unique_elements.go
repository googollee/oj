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
		line := scanner.Text()
		nums := strings.Split(line, ",")
		var elements []string
		hash := make(map[string]bool)
		for _, num := range nums {
			if !hash[num] {
				elements = append(elements, num)
				hash[num] = true
			}
		}
		fmt.Println(strings.Join(elements, ","))
	}
}
