package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]
	partOne(filename)
	partTwo(filename)
}

func partOne(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")

		if isSafe(vals) {
			safeCount++
		}
	}

	println(safeCount)
}

func partTwo(filename string) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	safeCount := 0

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")

		if isSafe(vals) {
			safeCount++
			continue
		}

		for i := 0; i < len(vals); i++ {
			newCopy := make([]string, len(vals))
			copy(newCopy, vals)
			newCopy = append(newCopy[:i], newCopy[i+1:]...)
			if isSafe(newCopy) {
				safeCount++
				break
			}
		}

	}

	println(safeCount)
}

func isSafe(vals []string) bool {
	increasing := true
	num1, _ := strconv.Atoi(string(vals[0]))
	num2, _ := strconv.Atoi(string(vals[1]))
	if num2 < num1 {
		increasing = false
	}
	prev := num1
	safe := true
	for i := 1; i < len(vals); i++ {
		num, _ := strconv.Atoi(string(vals[i]))
		if (increasing && (num <= prev || num-prev > 3)) || (!increasing && (num >= prev || prev-num > 3)) {
			safe = false
			break
		}
		prev = num
	}
	return safe
}
