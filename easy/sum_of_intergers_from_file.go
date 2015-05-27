package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		sum += num
	}
	fmt.Println(sum)
}
