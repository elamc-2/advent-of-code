package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Stack []string
type Command struct {
	n      int
	origin int
	end    int
}

var cmd []Command

func parse() (map[int]Stack, error) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("reading file", err)
	}

	split := strings.Split(string(input), "\n\n")

	//crates
	lastIdx := strings.LastIndex(split[0], "\n")
	crates := strings.Split(split[0][:lastIdx], "\n")
	var stack = make(map[int]Stack)
	for i := len(crates) - 1; i >= 0; i-- {
		row := crates[i]
		idx := 0
		for j := 0; j < len(row); j += 3 {
			if string(row[j]) == " " {
				j++
				idx++
			}
			crate := row[j : j+3]
			if len(strings.TrimSpace(crate)) > 0 {
				replacer := strings.NewReplacer("[", "", "]", "")
				stack[idx] = append(stack[idx], replacer.Replace(crate))
			}
		}
	}

	//cmds
	for _, v := range strings.Split(split[1], "\n") {
		reg := regexp.MustCompile("[0-9]+").FindAllString(v, -1)
		tmp := make([]int, 3)
		for i, n := range reg {
			tmp[i], _ = strconv.Atoi(n)
		}
		cmd = append(cmd, Command{tmp[0], tmp[1], tmp[2]})
	}
	return stack, nil
}

func (l *Stack) pop() (string, Stack) {
	last := (*l)[len(*l)-1]
	*l = (*l)[:len(*l)-1]
	return last, *l
}

func copy(in map[int]Stack) map[int]Stack {
	out := make(map[int]Stack)
	buf := new(bytes.Buffer)
	gob.NewEncoder(buf).Encode(in)
	gob.NewDecoder(buf).Decode(&out)
	return out
}

func print(m map[int]Stack) {
	for i := 0; i < len(m); i++ {
		stackAt := m[i]
		last, _ := stackAt.pop()
		fmt.Print(last)
	}
}

func One(m map[int]Stack) {
	for _, cmd := range cmd {
		for i := 0; i < cmd.n; i++ {
			stackAt := m[cmd.origin-1]
			last, updated := stackAt.pop()
			m[cmd.origin-1] = updated
			m[cmd.end-1] = append(m[cmd.end-1], last)
		}
	}
	print(m)
}

func Two(m map[int]Stack) {
	for _, cmd := range cmd {
		stackAt := m[cmd.origin-1]
		nStack := stackAt[len(stackAt)-cmd.n:]
		updated := stackAt[:len(stackAt)-cmd.n]
		m[cmd.origin-1] = updated
		m[cmd.end-1] = append(m[cmd.end-1], nStack...)
	}
	print(m)
}

func main() {
	data, err := parse()
	if err != nil {
		panic(err)
	}

	One(copy(data))
	Two(copy(data))
}
