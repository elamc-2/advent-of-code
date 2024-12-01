package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse() ([][]string, error) {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("reading file", err)
	}

	var o [][]string
	for _, v := range strings.Split(string(input), "\n") {
		pair := strings.Split(v, ",")
		o = append(o, pair)
	}
	return o, nil
}

func One(data [][]string) {
	pairs := 0
	for _, v := range data {
		pA := strings.Split(v[0], "-")
		pB := strings.Split(v[1], "-")

		a0, _ := strconv.Atoi(pA[0])
		a1, _ := strconv.Atoi(pA[1])
		b0, _ := strconv.Atoi(pB[0])
		b1, _ := strconv.Atoi(pB[1])

		if b0 >= a0 && b1 <= a1 || a0 >= b0 && a1 <= b1 {
			pairs++
		}
	}
	fmt.Println(pairs)
}

func Two(data [][]string) {
	pairs := 0
	for _, v := range data {
		pA := strings.Split(v[0], "-")
		pB := strings.Split(v[1], "-")

		a0, _ := strconv.Atoi(pA[0])
		a1, _ := strconv.Atoi(pA[1])
		b0, _ := strconv.Atoi(pB[0])
		b1, _ := strconv.Atoi(pB[1])

		if a1 >= b0 && a0 <= b1 {
			pairs++
		}
	}
	fmt.Println(pairs)
}

func main() {
	data, err := parse()
	if err != nil {
		panic(err)
	}

	One(data)
	Two(data)
}
