package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Node struct {
	label    string
	size     int
	children []*Node
	parent   *Node
}

var root Node

func parse() ([]string, error) {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("reading file", err)
	}

	split := strings.Split(string(input), "\n")
	toTree(split)
	return split, nil
}

func toTree(data []string) {
	root = Node{
		label:    "/",
	}
	current := &root

	for _, v := range data {
		switch {
		case strings.Contains(v, "dir"):
			current.children = append(current.children, &Node{
				label:    strings.Split(v, " ")[1],
				parent:   current,
			})
		case strings.Contains(v, "$"):
			if strings.Contains(v, "cd") {
				arg := strings.Split(v, " ")[2]
				if arg != ".." {
					for _, child := range current.children {
						if child.label == arg {
							current = child
							break
						}
					}
				} else {
					current = current.parent
				}
			}
		default:
			file := strings.Split(v, " ")
			size, _ := strconv.Atoi(file[0])
			current.size += size
		}
	}
}

func (node *Node) totalSize() int {
	sum := 0
	for _, child := range node.children {
		sum += child.totalSize()
	}
	return node.size + sum
}

func flatten() []*Node {
	nodes := make([]*Node, 0)
	queue := make([]*Node, 0)
	queue = append(queue, &root)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodes = append(nodes, node)
		if len(node.children) > 0 {
			queue = append(queue, node.children...)
		}
	}
	return nodes
}

func One() {
	sum := 0
	for _, n := range flatten() {
		if n.totalSize() <= 100000 {
			sum += n.totalSize()
		}
	}
	fmt.Println(sum)
}

func Two() {
	unused := 70000000 - root.totalSize()
	req := 30000000 - unused

	opt := make([]int, 0)
	for _, n := range flatten() {
		if n.totalSize() >= req {
			opt = append(opt, n.totalSize())
		}
	}
	sort.Ints(opt)
	fmt.Println(opt[0])
}

func main() {
	_, err := parse()
	if err != nil {
		panic(err)
	}

	One()
	Two()
}
