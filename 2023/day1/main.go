package main

import (
	"bufio"
	"os"
	"strconv"
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

	total := 0

	for scanner.Scan() {
		firstNum := ""
		lastNum := ""
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			char := line[i]
			if char >= '0' && char <= '9' {
				firstNum = string(char)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if char >= '0' && char <= '9' {
				lastNum = string(char)
				break
			}
		}
		numberString := firstNum + lastNum

		num, e := strconv.Atoi(numberString)

		if e != nil {
			panic(e)
		}

		total += num

	}

	println(total)
}

func partTwo() {
	f, err := os.Open("input.txt")

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0

	for scanner.Scan() {
		line := scanner.Text()

		firstIndex := len(line)
		lastIndex := -1
		firstNumber := ""
		lastNumber := ""

		oneIndex1 := getFirstIndex(line, "1", "one")
		if oneIndex1 > -1 && oneIndex1 < firstIndex {
			firstIndex = oneIndex1
			firstNumber = "1"
		}
		oneIndex2 := getLastIndex(line, "1", "one")
		if oneIndex2 > -1 && oneIndex2 > lastIndex {
			lastIndex = oneIndex2
			lastNumber = "1"
		}

		twoIndex1 := getFirstIndex(line, "2", "two")
		if twoIndex1 > -1 && twoIndex1 < firstIndex {
			firstIndex = twoIndex1
			firstNumber = "2"
		}
		twoIndex2 := getLastIndex(line, "2", "two")
		if twoIndex2 > -1 && twoIndex2 > lastIndex {
			lastIndex = twoIndex2
			lastNumber = "2"
		}

		threeIndex1 := getFirstIndex(line, "3", "three")
		if threeIndex1 > -1 && threeIndex1 < firstIndex {
			firstIndex = threeIndex1
			firstNumber = "3"
		}
		threeIndex2 := getLastIndex(line, "3", "three")
		if threeIndex2 > -1 && threeIndex2 > lastIndex {
			lastIndex = threeIndex2
			lastNumber = "3"
		}

		fourIndex1 := getFirstIndex(line, "4", "four")
		if fourIndex1 > -1 && fourIndex1 < firstIndex {
			firstIndex = fourIndex1
			firstNumber = "4"
		}
		fourIndex2 := getLastIndex(line, "4", "four")
		if fourIndex2 > -1 && fourIndex2 > lastIndex {
			lastIndex = fourIndex2
			lastNumber = "4"
		}

		fiveIndex1 := getFirstIndex(line, "5", "five")
		if fiveIndex1 > -1 && fiveIndex1 < firstIndex {
			firstIndex = fiveIndex1
			firstNumber = "5"
		}
		fiveIndex2 := getLastIndex(line, "5", "five")
		if fiveIndex2 > -1 && fiveIndex2 > lastIndex {
			lastIndex = fiveIndex2
			lastNumber = "5"
		}

		sixIndex1 := getFirstIndex(line, "6", "six")
		if sixIndex1 > -1 && sixIndex1 < firstIndex {
			firstIndex = sixIndex1
			firstNumber = "6"
		}
		sixIndex2 := getLastIndex(line, "6", "six")
		if sixIndex2 > -1 && sixIndex2 > lastIndex {
			lastIndex = sixIndex2
			lastNumber = "6"
		}

		sevenIndex1 := getFirstIndex(line, "7", "seven")
		if sevenIndex1 > -1 && sevenIndex1 < firstIndex {
			firstIndex = sevenIndex1
			firstNumber = "7"
		}
		sevenIndex2 := getLastIndex(line, "7", "seven")
		if sevenIndex2 > -1 && sevenIndex2 > lastIndex {
			lastIndex = sevenIndex2
			lastNumber = "7"
		}

		eightIndex1 := getFirstIndex(line, "8", "eight")
		if eightIndex1 > -1 && eightIndex1 < firstIndex {
			firstIndex = eightIndex1
			firstNumber = "8"
		}
		eightIndex2 := getLastIndex(line, "8", "eight")
		if eightIndex2 > -1 && eightIndex2 > lastIndex {
			lastIndex = eightIndex2
			lastNumber = "8"
		}

		nineIndex1 := getFirstIndex(line, "9", "nine")
		if nineIndex1 > -1 && nineIndex1 < firstIndex {
			firstNumber = "9"
		}
		nineIndex2 := getLastIndex(line, "9", "nine")
		if nineIndex2 > -1 && nineIndex2 > lastIndex {
			lastNumber = "9"
		}

		numberString := firstNumber + lastNumber

		num, e := strconv.Atoi(numberString)

		if e != nil {
			panic(e)
		}

		total += num

	}

	println(total)
}

func getFirstIndex(str string, substring1 string, substring2 string) int {
	index1 := strings.Index(str, substring1)
	index2 := strings.Index(str, substring2)
	println(str, substring1, index1, substring2, index2)
	if index1 == -1 && index2 == -1 {
		return -1
	} else if index1 > -1 && index2 > -1 {
		if index1 < index2 {
			return index1
		} else {
			return index2
		}
	} else if index1 > -1 {
		return index1
	} else {
		return index2
	}
}

func getLastIndex(str string, substring1 string, substring2 string) int {
	index1 := strings.LastIndex(str, substring1)
	index2 := strings.LastIndex(str, substring2)
	if index1 == -1 && index2 == -1 {
		return -1
	} else if index1 > -1 && index2 > -1 {
		if index1 > index2 {
			return index1
		} else {
			return index2
		}
	} else if index1 > -1 {
		return index1
	} else {
		return index2
	}
}
