package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 12; i++ {
		for j := 1; j <= 12; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
