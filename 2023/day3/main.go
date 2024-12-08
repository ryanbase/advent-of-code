package main

import (
	"bufio"
	"os"
	"strconv"
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

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0

	for i, line := range lines {
		currentNumber := ""
		for j, c := range line {
			char := byte(c)
			if isNumber(char) {
				currentNumber += string(char)
			}
			if currentNumber != "" && (!isNumber(char) || j == len(line)-1) {
				start := j - (len(currentNumber) + 1)
				if j == len(line)-1 && isNumber(char) {
					start = j - len(currentNumber)
				}
				if start < 0 {
					start = 0
				}
				lineIndicesToCheck := make([]int, 0)
				if i > 0 {
					lineIndicesToCheck = append(lineIndicesToCheck, i-1)
				}
				lineIndicesToCheck = append(lineIndicesToCheck, i)
				if i < len(lines)-1 {
					lineIndicesToCheck = append(lineIndicesToCheck, i+1)
				}
				symbolFound := false
				for _, lineIndexToCheck := range lineIndicesToCheck {
					if symbolFound {
						break
					}
					lineToCheck := lines[lineIndexToCheck]
					for k := start; k <= j; k++ {
						if isSymbol(lineToCheck[k]) {
							symbolFound = true
							break
						}
					}
				}

				if symbolFound {
					num, err := strconv.Atoi(currentNumber)
					if err != nil {
						panic(err)
					}
					total += num
				}

				currentNumber = ""
			}
		}
	}

	println(total)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var lines []string

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	total := 0

	for i, line := range lines {
		for j, c := range line {
			char := byte(c)
			if !isStar(char) {
				continue
			}
			numbers := make([]string, 0)
			// Check above
			if i > 0 {
				number := ""
				if isNumber(lines[i-1][j]) {
					number = number + string(lines[i-1][j])
					for k := j - 1; k >= 0; k-- {
						if isNumber(lines[i-1][k]) {
							number = string(lines[i-1][k]) + number
						} else {
							break
						}
					}
					for k := j + 1; k < len(line); k++ {
						if isNumber(lines[i-1][k]) {
							number = number + string(lines[i-1][k])
						} else {
							break
						}
					}
					numbers = append(numbers, number)
				} else {
					// Upper left
					start := j - 1
					if start >= 0 && isNumber(lines[i-1][start]) {
						number := ""
						for k := start; k >= 0; k-- {
							if isNumber(lines[i-1][k]) {
								number = string(lines[i-1][k]) + number
							} else {
								break
							}
						}
						numbers = append(numbers, number)
					}

					// Upper right
					start = j + 1
					if start < len(line) && isNumber(lines[i-1][start]) {
						number := ""
						for k := start; k < len(line); k++ {
							if isNumber(lines[i-1][k]) {
								number = number + string(lines[i-1][k])
							} else {
								break
							}
						}
						numbers = append(numbers, number)
					}
				}
			}
			// Check left
			if j > 0 {
				start := j - 1
				if isNumber(line[start]) {
					number := ""
					for k := start; k >= 0; k-- {
						if isNumber(line[k]) {
							number = string(line[k]) + number
						} else {
							break
						}
					}
					numbers = append(numbers, number)
				}
			}

			// Check right
			if j < len(line)-1 {
				start := j + 1
				if isNumber(line[start]) {
					number := ""
					for k := start; k < len(line); k++ {
						if isNumber(line[k]) {
							number = number + string(line[k])
						} else {
							break
						}
					}
					numbers = append(numbers, number)
				}
			}

			// Check below
			if i < len(lines)-1 {
				if isNumber(lines[i+1][j]) {
					number := string(lines[i+1][j])
					for k := j - 1; k >= 0; k-- {
						if isNumber(lines[i+1][k]) {
							number = string(lines[i+1][k]) + number
						} else {
							break
						}
					}
					for k := j + 1; k < len(line); k++ {
						if isNumber(lines[i+1][k]) {
							number = number + string(lines[i+1][k])
						} else {
							break
						}
					}
					numbers = append(numbers, number)
				} else {
					// Below left
					start := j - 1
					if start >= 0 && isNumber(lines[i+1][start]) {
						number := ""
						for k := start; k >= 0; k-- {
							if isNumber(lines[i+1][k]) {
								number = string(lines[i+1][k]) + number
							} else {
								break
							}
						}
						numbers = append(numbers, number)
					}

					// Below right
					start = j + 1
					if start < len(line) && isNumber(lines[i+1][start]) {
						number := ""
						for k := start; k < len(line); k++ {
							if isNumber(lines[i+1][k]) {
								number = number + string(lines[i+1][k])
							} else {
								break
							}
						}
						numbers = append(numbers, number)
					}
				}
			}

			if len(numbers) == 2 {
				num1, err := strconv.Atoi(numbers[0])
				if err != nil {
					panic(err)
				}
				num2, err := strconv.Atoi(numbers[1])
				if err != nil {
					panic(err)
				}
				total += num1 * num2
			}
		}
	}

	println(total)

}

func isNumber(value byte) bool {
	return value >= 48 && value <= 57
}

func isSymbol(value byte) bool {
	return !isNumber(value) && value != 46
}

func isStar(value byte) bool {
	return value == 42
}
