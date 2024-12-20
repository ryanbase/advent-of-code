package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	filename := utils.GetFileNameFromArgument()
	input := utils.ReadInputAsString(filename)
	part1(input)
	part2(input)
}

func part1(input string) {
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

func part2(input string) {
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
