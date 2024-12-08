package main

import (
	"bufio"
	"fmt"
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

	times := make([]int, 0)
	dists := make([]int, 0)

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	text := strings.Split(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]), " ")
	for _, t := range text {
		if t != "" {
			num, err := strconv.Atoi(strings.TrimSpace(t))
			if err != nil {
				panic(err)
			}
			times = append(times, num)
		}
	}
	scanner.Scan()
	text = strings.Split(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]), " ")
	for _, t := range text {
		if t != "" {
			num, err := strconv.Atoi(strings.TrimSpace(t))
			if err != nil {
				panic(err)
			}
			dists = append(dists, num)
		}
	}

	counts := make([]int, len(times))

	for i := 0; i < len(times); i++ {
		time := times[i]
		record := dists[i]
		for j := 1; j <= time; j++ {
			dist := j * (time - j)
			if dist > record {
				counts[i]++
			}
		}
	}

	fmt.Println(counts)

	total := 1

	for _, count := range counts {
		total *= count
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
	scanner.Scan()
	text := strings.ReplaceAll(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]), " ", "")
	num, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		panic(err)
	}
	time := num

	scanner.Scan()
	text = strings.ReplaceAll(strings.TrimSpace(strings.Split(scanner.Text(), ":")[1]), " ", "")
	num, err = strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		panic(err)
	}
	record := num

	count := 0

	for j := 1; j <= time; j++ {
		dist := j * (time - j)
		if dist > record {
			count++
		}
	}

	println(count)
}
