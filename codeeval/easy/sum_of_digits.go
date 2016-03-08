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
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		var num int
		fmt.Sscanf(scanner.Text(), "%d", &num)
		ret := 0
		for num > 0 {
			ret += num % 10
			num /= 10
		}
		fmt.Println(ret)
	}
}
