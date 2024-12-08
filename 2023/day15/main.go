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

	inputs := ""

	for scanner.Scan() {
		inputs += scanner.Text()
	}

	currentValue := 0

	for _, input := range strings.Split(inputs, ",") {
		currentValue += HASH(input)
	}

	println(currentValue)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	inputs := ""

	for scanner.Scan() {
		inputs += scanner.Text()
	}

	boxes := make([][]Lens, 256)

	for _, input := range strings.Split(inputs, ",") {
		if strings.Contains(input, "=") {
			split := strings.Split(input, "=")
			lens := Lens{split[0], split[1]}
			box := HASH(lens.label)
			exists := false
			for i, boxLens := range boxes[box] {
				if boxLens.label == lens.label {
					boxes[box][i] = lens
					exists = true
					break
				}
			}
			if !exists {
				boxes[box] = append(boxes[box], lens)
			}
		} else if strings.Contains(input, "-") {
			label := strings.Split(input, "-")[0]
			box := HASH(label)
			for i, boxLens := range boxes[box] {
				if boxLens.label == label {
					boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
					break
				}
			}
		}
	}

	total := 0

	for i, box := range boxes {
		for slot, lens := range box {
			focal, err := strconv.Atoi(lens.focal)
			if err != nil {
				panic(err)
			}
			total += (i + 1) * (slot + 1) * focal
		}
	}

	println(total)
}

func HASH(input string) int {
	currentValue := 0
	for _, ascii := range input {
		currentValue += int(ascii)
		currentValue *= 17
		currentValue %= 256
	}
	return currentValue
}

type Lens struct {
	label string
	focal string
}
