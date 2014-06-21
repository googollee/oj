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

		a, b := 0, 1
		for num > 0 {
			a, b = b, a+b
			num--
		}
		fmt.Println(a)
	}
}
