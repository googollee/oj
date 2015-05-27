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
		lines := strings.SplitN(scanner.Text(), ",", 2)
		str, ch := lines[0], lines[1]
		fmt.Println(strings.LastIndex(str, ch))
	}
}
