package main

import (
	"bufio"
	"math"
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

	total := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		matches := getNumMatches(getWinningNumbers(line), getOwnNumbers(line))
		total += int(math.Pow(2, float64(matches-1)))
	}

	println(total)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	total := 0

	matchMap := make(map[int]int)

	scanner := bufio.NewScanner(f)
	cardNum := 1
	for scanner.Scan() {
		line := scanner.Text()
		matches := getNumMatches(getWinningNumbers(line), getOwnNumbers(line))
		matchMap[cardNum] = matches
		cardNum++
	}

	cardTotals := make([]int, cardNum-1)

	for i := range cardTotals {
		cardTotals[i] = cardTotals[i] + 1
		for j := i + 1; j < i+matchMap[i+1]+1; j++ {
			cardTotals[j] = cardTotals[j] + cardTotals[i]
		}
		total += cardTotals[i]
	}

	println(total)

}

func getWinningNumbers(line string) []string {
	numbers := strings.TrimSpace(strings.Split(line, ": ")[1])
	winningNumbersStr := strings.TrimSpace(strings.Split(numbers, " | ")[0])
	winningNumbersSplit := strings.Split(winningNumbersStr, " ")
	winningNumbers := make([]string, 0)
	for _, numStr := range winningNumbersSplit {
		numStr = strings.TrimSpace(numStr)
		if numStr != "" {
			winningNumbers = append(winningNumbers, numStr)
		}
	}
	return winningNumbers
}

func getOwnNumbers(line string) []string {
	numbers := strings.TrimSpace(strings.Split(line, ": ")[1])
	ownNumbersStr := strings.TrimSpace(strings.Split(numbers, " | ")[1])
	ownNumbersSplit := strings.Split(ownNumbersStr, " ")
	ownNumbers := make([]string, 0)
	for _, numStr := range ownNumbersSplit {
		numStr = strings.TrimSpace(numStr)
		if numStr != "" {
			ownNumbers = append(ownNumbers, numStr)
		}
	}
	return ownNumbers
}

func getNumMatches(winningNums []string, ownNums []string) int {
	matches := 0
	for _, ownNum := range ownNums {
		for _, winningNum := range winningNums {
			if ownNum == winningNum {
				matches++
				break
			}
		}
	}
	return matches
}
