package main

import (
	"bufio"
	"os"
	"slices"
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

	list1 := []int{}
	list2 := []int{}

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		for i, val := range vals {
			if i == 0 {
				num, err := strconv.Atoi(string(val))
				if err != nil {
					panic("Error converting first num")
				}
				list1 = append(list1, num)
			} else if val != "" {
				num, err := strconv.Atoi(string(val))
				if err != nil {
					panic("Error converting second num")
				}
				list2 = append(list2, num)
			}
		}
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sum := 0

	for i := 0; i < len(list1); i++ {
		diff := 0
		if list1[i] > list2[i] {
			diff = list1[i] - list2[i]
		} else {
			diff = list2[i] - list1[i]
		}
		sum += diff
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

	list1 := []int{}
	counts := make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		vals := strings.Split(line, " ")
		for i, val := range vals {
			if i == 0 {
				num, err := strconv.Atoi(string(val))
				if err != nil {
					panic("Error converting first num")
				}
				list1 = append(list1, num)
			} else if val != "" {
				num, err := strconv.Atoi(string(val))
				if err != nil {
					panic("Error converting second num")
				}
				_, ok := counts[num]
				if !ok {
					counts[num] = 0
				}
				counts[num]++
			}
		}
	}

	total := 0

	for _, val := range list1 {
		total += val * counts[val]
	}

	println(total)
}
