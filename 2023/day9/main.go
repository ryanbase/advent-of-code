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

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strToIntArray(line)
		next := getNextValue(nums)
		sum += next
	}

	println(sum)
}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		nums := strToIntArray(line)
		prev := getPrevValue(nums)
		sum += prev
	}

	println(sum)
}

func strToIntArray(str string) []int {
	arr := make([]int, 0)
	for _, val := range strings.Split(str, " ") {
		num, err := strconv.Atoi(val)
		if err != nil {
			panic(err)
		}
		arr = append(arr, num)
	}
	return arr
}

func getDiffArray(arr []int) []int {
	diffs := make([]int, len(arr)-1)
	for i := 0; i < len(arr)-1; i++ {
		diffs[i] = arr[i+1] - arr[i]
	}
	return diffs
}

func getNextValue(arr []int) int {
	if isAllZero(arr) {
		return 0
	}
	diffs := getDiffArray(arr)
	nextVal := getNextValue(diffs)
	return arr[len(arr)-1] + nextVal
}

func getPrevValue(arr []int) int {
	if isAllZero(arr) {
		return 0
	}
	diffs := getDiffArray(arr)
	prevValue := getPrevValue(diffs)
	return arr[0] - prevValue
}

func isAllZero(arr []int) bool {
	for _, val := range arr {
		if val != 0 {
			return false
		}
	}
	return true
}
