package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]

	part1(filename)
	part2(filename)
}

func part1(filename string) {
	defer TimeTrack(time.Now())
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

		stack := []int{}
		stack = append(stack, values[0])
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
	defer TimeTrack(time.Now())
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

		stack := []int{}
		stack = append(stack, values[0])
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

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)

	// Skip this function, and fetch the PC and file for its parent.
	pc, _, _, _ := runtime.Caller(1)

	// Retrieve a function object this functions parent.
	funcObj := runtime.FuncForPC(pc)

	// Regex to extract just the function name (and not the module path).
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")

	fmt.Printf("%s completed in %s\n", name, elapsed)
}
