package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	utils "github.com/ryanbase/advent-of-code/2024/utils"
)

type location struct {
	row int
	col int
}

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]
	input := readInput(filename)
	starts := findStarts(input)
	part1(input, starts)
	part2(input, starts)
}

func readInput(filename string) [][]int {
	defer utils.TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		row := []int{}
		for _, val := range arr {
			num, _ := strconv.Atoi(val)
			row = append(row, num)
		}
		input = append(input, row)

	}
	return input
}

func findStarts(input [][]int) []location {
	starts := []location{}
	for i, row := range input {
		for j, val := range row {
			if val == 0 {
				starts = append(starts, location{i, j})
			}
		}
	}
	return starts
}

func part1(input [][]int, starts []location) {
	defer utils.TimeTrack(time.Now())
	total := 0

	for _, start := range starts {
		total += getNumTrails(input, start)
	}

	println(total)
}

func part2(input [][]int, starts []location) {
	defer utils.TimeTrack(time.Now())
	total := 0

	for _, start := range starts {
		total += getDistinctTrails(input, start)
	}

	println(total)
}

func getNumTrails(input [][]int, start location) int {
	stack := []location{start}
	trails := make(map[location]struct{})
	for len(stack) > 0 {
		currLoc := stack[0]
		stack = slices.Delete(stack, 0, 1)
		row := currLoc.row
		col := currLoc.col
		currVal := input[row][col]

		if currVal == 9 {
			trail := location{row, col}
			trails[trail] = struct{}{}
			continue
		}

		if row > 0 && input[row-1][col] == currVal+1 {
			stack = append(stack, location{row - 1, col})
		}
		if row < len(input)-1 && input[row+1][col] == currVal+1 {
			stack = append(stack, location{row + 1, col})
		}
		if col > 0 && input[row][col-1] == currVal+1 {
			stack = append(stack, location{row, col - 1})
		}
		if col < len(input[row])-1 && input[row][col+1] == currVal+1 {
			stack = append(stack, location{row, col + 1})
		}
	}
	return len(trails)
}

func getDistinctTrails(input [][]int, start location) int {
	stack := []location{start}
	trails := []location{}
	for len(stack) > 0 {
		currLoc := stack[0]
		stack = slices.Delete(stack, 0, 1)
		row := currLoc.row
		col := currLoc.col
		currVal := input[row][col]

		if currVal == 9 {
			trails = append(trails, location{row, col})
			continue
		}

		if row > 0 && input[row-1][col] == currVal+1 {
			stack = append(stack, location{row - 1, col})
		}
		if row < len(input)-1 && input[row+1][col] == currVal+1 {
			stack = append(stack, location{row + 1, col})
		}
		if col > 0 && input[row][col-1] == currVal+1 {
			stack = append(stack, location{row, col - 1})
		}
		if col < len(input[row])-1 && input[row][col+1] == currVal+1 {
			stack = append(stack, location{row, col + 1})
		}
	}
	return len(trails)
}
