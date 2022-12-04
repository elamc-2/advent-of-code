package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func parse() ([]string, error) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("reading file", err)
	}
	var o []string = strings.Split(string(input), "\n")
	return o, nil
}

const (
	START_LOWER rune = rune('a') - 1
	START_UPPER rune = rune('A') - 1
)

func unicodeSum(char rune) int {
	if unicode.IsLower(char) {
		return int(char - START_LOWER)
	} else {
		return int(char-START_UPPER) + 26
	}
}

func One(data []string) {
	sum := 0
	for _, v := range data {
		half := len(v) / 2
		split := []string{v[:half], v[half:]}

		var char rune
		for _, c := range split[0] {
			if strings.Contains(split[1], string(c)) {
				char = c
				break
			}
		}
		sum += unicodeSum(char)
	}
	fmt.Println(sum)
}

func Two(data []string) {
	sum := 0
	for i := 0; i < len(data); i += 3 {
		var char rune
		for _, c := range data[i] {
			if strings.Contains(data[i+1], string(c)) && strings.Contains(data[i+2], string(c)) {
				char = c
				break
			}
		}
		sum += unicodeSum(char)
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
