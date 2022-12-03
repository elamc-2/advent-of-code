package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func parse() ([][]string, error) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("reading file", err)
	}

	var o [][]string
	for _, v := range strings.Split(string(input), "\n") {
		o = append(o, strings.Split(v, " "))
	}
	return o, nil
}

type Hand int
const (
	Rock     Hand = 1
	Paper    Hand = 2
	Scissors Hand = 3
)

type Result int
const (
	Lose Result = 0
	Draw Result = 3
	Win  Result = 6
)

func hand(i string) Hand {
	var h Hand
	switch i {
	case "A", "X":
		h = Rock
	case "B", "Y":
		h = Paper
	case "C", "Z":
		h = Scissors
	}
	return h
}

func outcome(i string) Result {
	var r Result
	switch i {
	case "X":
		r = Lose
	case "Y":
		r = Draw
	case "Z":
		r = Win
	}
	return r
}

var moves map[Hand][]Hand = map[Hand][]Hand{
	Paper:    {Rock, Scissors},
	Scissors: {Paper, Rock},
	Rock:     {Scissors, Paper},
}

func One(data [][]string) {
	total := 0
	for _, v := range data {
		score := 3
		if moves[hand(v[1])][0] == hand(v[0]) {
			score = 6
		} else if moves[hand(v[1])][1] == hand(v[0]) {
			score = 0
		}
		total += int(hand(v[1])) + score
	}
	fmt.Println(total)
}

func Two(data [][]string) {
	total := 0
	for _, v := range data {
		score := Draw
		h := hand(v[0])
		if outcome(v[1]) == Win {
			h = moves[h][1]
			score = Win
		} else if outcome(v[1]) == Lose {
			h = moves[h][0]
			score = Lose
		}
		total += int(h) + int(score)
	}
	fmt.Println(total)
}

func main() {
	data, err := parse()
	if err != nil {
		panic(err)
	}

	One(data)
	Two(data)
}
