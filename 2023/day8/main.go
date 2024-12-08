package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan() // Skip empty line

	nodes := make(map[string]Node)

	currentNode := "AAA"
	destination := "ZZZ"

	for scanner.Scan() {
		line := scanner.Text()
		name, node := parseNode(line)
		nodes[name] = node
	}

	steps := 0

	for true {
		for _, direction := range instructions {
			if currentNode == destination {
				break
			}
			if direction == 'L' {
				currentNode = nodes[currentNode].left
			}
			if direction == 'R' {
				currentNode = nodes[currentNode].right
			}
			steps++
		}
		if currentNode == destination {
			break
		}
	}

	println(steps)
}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	instructions := scanner.Text()
	scanner.Scan() // Skip empty line

	nodes := make(map[string]Node)

	startingNodes := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		name, node := parseNode(line)
		nodes[name] = node
		if endsWith(name, 'A') {
			startingNodes = append(startingNodes, name)
		}
	}

	nodeSteps := make([]int, len(startingNodes))

	for i, node := range startingNodes {
		currentNode := node
		steps := 0
		for true {
			for _, direction := range instructions {
				if endsWith(currentNode, 'Z') {
					break
				}
				if direction == 'L' {
					currentNode = nodes[currentNode].left
				}
				if direction == 'R' {
					currentNode = nodes[currentNode].right
				}
				steps++
			}
			if endsWith(currentNode, 'Z') {
				break
			}
		}
		nodeSteps[i] = steps
	}

	println(LCM(nodeSteps[0], nodeSteps[1], nodeSteps[2:]...))
}

func parseNode(line string) (string, Node) {
	split := strings.Split(line, " = ")
	name := split[0]
	directions := strings.ReplaceAll(split[1], "(", "")
	directions = strings.ReplaceAll(directions, ")", "")
	directionsSplit := strings.Split(directions, ", ")
	return name, Node{directionsSplit[0], directionsSplit[1]}
}

func endsWith(str string, char byte) bool {
	return str[len(str)-1] == char
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// https://siongui.github.io/2017/06/03/go-find-lcm-by-gcd/
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Node struct {
	left  string
	right string
}
