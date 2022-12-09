package main

import (
	"fmt"
	"log"
	"os"
)

func parse() (string, error) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("reading file", err)
	}
	return string(input), nil
}

func set(str string) map[rune]struct{} {
	set := make(map[rune]struct{})
	for _, c := range str {
		set[c] = struct{}{}
	}
	return set
}

func One(data string) {
	for i := 0; i < len(data)-4; i++ {
		if len(set(data[i:i+4])) == 4 {
			fmt.Println(len(data[0 : i+4]))
			break
		}
	}
}

func Two(data string) {
	for i := 0; i < len(data)-14; i++ {
		if len(set(data[i:i+14])) == 14 {
			fmt.Println(len(data[0 : i+14]))
			break
		}
	}
}

func main() {
	data, err := parse()
	if err != nil {
		panic(err)
	}

	One(data)
	Two(data)
}
