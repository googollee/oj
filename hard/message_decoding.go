package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

var sTol = map[string]int{
	"001": 1,
	"010": 2,
	"011": 3,
	"100": 4,
	"101": 5,
	"110": 6,
	"111": 7,
}

func seperateHeadBody(str string) (string, string) {
	s := strings.LastIndexFunc(str, func(r rune) bool {
		if '0' <= r && r <= '9' {
			return false
		}
		return true
	})
	head := str[:s+1]
	body := str[s+1:]
	return head, body
}

func nTos(n int) string {
	b := uint(1)
	for {
		max := 1<<b - 1
		if n < max {
			break
		}
		n -= max
		b++
	}
	f := fmt.Sprintf("%%0%db", b)
	ret := fmt.Sprintf(f, n)
	return ret
}

func parseHead(head string) map[string]rune {
	ret := make(map[string]rune)
	for i, r := range head {
		ret[nTos(i)] = r
	}
	return ret
}

func segEnd(s string) bool {
	for _, r := range s {
		if r != '1' {
			return false
		}
	}
	return true
}

func decoding(head map[string]rune, body string) string {
	ret := bytes.NewBuffer(nil)
	for len(body) > 0 {
		segHead := body[:3]
		body = body[3:]
		if segHead == "000" {
			break
		}
		l := sTol[segHead]
		for {
			data := body[:l]
			body = body[l:]
			if segEnd(data) {
				break
			}
			ret.WriteRune(head[data])
		}
	}
	return ret.String()
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		headS, body := seperateHeadBody(str)
		head := parseHead(headS)
		fmt.Println(decoding(head, body))
	}
}
