package main

import (
	"bufio"
	"os"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

}

func partTwo() {
	// f, err := os.Open("input.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// inputs := ""

	// for scanner.Scan() {
	// 	inputs += scanner.Text()
	// }
}
