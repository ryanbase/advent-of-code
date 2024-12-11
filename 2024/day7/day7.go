package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	filename := utils.GetFileNameFromArgument()

	part1(filename)
	part2(filename)
}

func part1(filename string) {
	defer utils.TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		expected, _ := strconv.Atoi(split[0])
		valuesStrs := strings.Split(strings.Trim(split[1], " "), " ")
		values := []int{}
		for _, val := range valuesStrs {
			num, _ := strconv.Atoi(val)
			values = append(values, num)
		}

		calc := calculate1(values, 1, values[0], expected)
		if calc {
			res += expected
		}
	}

	println(res)
}

func calculate1(values []int, index int, curr int, expected int) bool {
	if index == len(values) {
		return curr == expected
	}
	add := calculate1(values, index+1, curr+values[index], expected)
	mult := calculate1(values, index+1, curr*values[index], expected)
	return add || mult
}

func part2(filename string) {
	defer utils.TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	res := 0

	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ":")
		expected, _ := strconv.Atoi(split[0])
		valuesStrs := strings.Split(strings.Trim(split[1], " "), " ")
		values := []int{}
		for _, val := range valuesStrs {
			num, _ := strconv.Atoi(val)
			values = append(values, num)
		}

		calc := calculate2(values, 1, values[0], expected)
		if calc {
			res += expected
		}
	}

	println(res)
}

func calculate2(values []int, index int, curr int, expected int) bool {
	if index == len(values) {
		return curr == expected
	}
	add := calculate2(values, index+1, curr+values[index], expected)
	mult := calculate2(values, index+1, curr*values[index], expected)
	concatVal, _ := strconv.Atoi(strconv.Itoa(curr) + strconv.Itoa(values[index]))
	concat := calculate2(values, index+1, concatVal, expected)
	return add || mult || concat
}
