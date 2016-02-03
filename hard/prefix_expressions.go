package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Expression struct {
	symbols []string
	numbers []float64
}

func (e *Expression) Push(token string) {
	switch token {
	case "+":
		fallthrough
	case "*":
		fallthrough
	case "/":
		e.symbols = append(e.symbols, token)
	default:
		n, err := strconv.ParseFloat(token, 64)
		if err != nil {
			log.Fatalf("invalid number(%s): %s", token, err)
		}
		e.numbers = append(e.numbers, n)
		e.Calculate()
	}
}

func (e *Expression) Calculate() {
	for {
		sl := len(e.symbols)
		if sl <= 0 {
			return
		}
		s := e.symbols[sl-1]
		nl := len(e.numbers)
		if nl < 2 {
			return
		}
		n1 := e.numbers[nl-2]
		n2 := e.numbers[nl-1]
		e.symbols = e.symbols[:sl-1]
		e.numbers = e.numbers[:nl-2]
		v := float64(0)
		switch s {
		case "+":
			v = n1 + n2
		case "*":
			v = n1 * n2
		case "/":
			v = n1 / n2
		}
		e.numbers = append(e.numbers, v)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strs := strings.Split(scanner.Text(), " ")
		exp := Expression{}
		for _, s := range strs {
			exp.Push(s)
		}
		exp.Calculate()
		fmt.Println(int(exp.numbers[0] + 0.1))
	}
}
