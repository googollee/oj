package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func substr(main, sub []byte) int {
	lsub := len(sub)
	for i := 0; i < len(main)-lsub+1; i++ {
		if main[i] == sub[0] {
			match := true
			for j, b := range sub {
				if main[i+j] != b {
					match = false
					break
				}
			}
			if match {
				return i
			}
		}
	}
	return -1
}

func unescape(pattern string) [][]byte {
	bpattern := bytes.Trim([]byte(pattern), "*")
	if len(bpattern) == 0 {
		return nil
	}
	split := bytes.Split(bpattern, []byte("*"))
	ret := make([][]byte, 1, len(split))
	ret[0] = split[0]
	for i := 1; i < len(split); i++ {
		if len(split[i]) == 0 {
			continue
		}
		l := len(ret)
		last := ret[l-1]
		lastc := last[len(last)-1]
		if lastc == '\\' {
			ret[l-1] = append(ret[l-1][:len(last)-1], '*')
			ret[l-1] = append(ret[l-1], split[i]...)
		} else {
			ret = append(ret, split[i])
		}
	}
	return ret
}

func match(str string, pattern [][]byte) bool {
	bstr := []byte(str)
	for _, p := range pattern {
		ret := substr(bstr, p)
		if ret < 0 {
			return false
		}
		bstr = bstr[ret+len(p):]
	}
	return true
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
		split := strings.SplitN(str, ",", 2)
		org, sub := split[0], split[1]
		pattern := unescape(sub)
		fmt.Println(match(org, pattern))
	}
}
