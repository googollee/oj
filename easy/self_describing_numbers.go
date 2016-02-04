package main

import (
	"fmt"
	"log"
)
import "bufio"
import "os"

func convertToBytes(s string) []byte {
	ret := make([]byte, len(s))
	for i := range ret {
		ret[i] = s[i] - '0'
	}
	return ret
}

func checkSelf(n []byte) int {
	var count [10]byte
	for _, c := range n {
		count[c]++
	}
	for i := range n {
		if n[i] != count[i] {
			return 0
		}
	}
	return 1
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		b := convertToBytes(scanner.Text())
		fmt.Println(checkSelf(b))
	}
}
