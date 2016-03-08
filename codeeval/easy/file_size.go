package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	info, err := os.Stat(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(info.Size())
}
