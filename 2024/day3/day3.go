package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]
	input := readInput(filename)

	partOne(input)
	partTwo(input)
}

func readInput(filename string) string {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := ""
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}
	return input
}

func partOne(input string) {
	regex, _ := regexp.Compile("mul\\(\\d+,\\d+\\)")
	matches := regex.FindAllStringIndex(input, -1)

	result := 0
	for _, match := range matches {
		nums := strings.Split(input[match[0]+4:match[1]-1], ",")
		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])
		result += num1 * num2
	}

	println(result)
}

func partTwo(input string) {
	regex, _ := regexp.Compile("don't\\(\\)|do\\(\\)|mul\\(\\d+,\\d+\\)")
	matches := regex.FindAllStringIndex(input, -1)
	do := true
	result := 0

	for _, match := range matches {
		command := input[match[0]:match[1]]
		if command == "do()" {
			do = true
		} else if command == "don't()" {
			do = false
		} else if do {
			nums := strings.Split(input[match[0]+4:match[1]-1], ",")
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			result += num1 * num2
		}
	}

	println(result)
}
