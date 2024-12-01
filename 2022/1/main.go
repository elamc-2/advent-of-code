package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parse() ([]int, error) {
	input, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("reading file", err)
	}

	var o []int
	for _, v := range strings.Split(string(input), "\n\n") {
		sum := 0
		for _, s := range strings.Split(v, "\n") {
			conv, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			sum += conv
		}
		o = append(o, sum)
	}
	return o, nil
}

func One(data []int) {
	sort.Ints(data)
	fmt.Println(data[len(data)-1])
}

func Two(data []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(data)))
	sum := 0
	for _, v := range data[:3] {
		sum += v
	}
	fmt.Println(sum)
}

func main() {
	data, err := parse()
	if err != nil {
		panic(err)
	}

	One(data)
	Two(data)
}
