package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
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

	input := []string{}

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	start, err := getStartingPosition(input)
	if err != nil {
		panic(err)
	}

	end, err := getEndPosition(input)
	if err != nil {
		panic(err)
	}

	println(count(input, start, end))

}

func partTwo() {
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

func getStartingPosition(input []string) (coord, error) {
	for col, char := range input[0] {
		if char == '.' {
			return coord{0, col, 0}, nil
		}
	}
	return coord{-1, -1, -1}, errors.New("Starting position not found")
}

func getEndPosition(input []string) (coord, error) {
	for col, char := range input[len(input)-1] {
		if char == '.' {
			return coord{len(input) - 1, col, 0}, nil
		}
	}
	return coord{-1, -1, -1}, errors.New("End position not found")
}

func count(input []string, start coord, end coord) int {
	visited := make(map[string]int)
	queue := []coord{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if input[pos.row][pos.col] == '#' {
			continue
		}

		mapKey := getMapKey(pos)
		_, visit := visited[mapKey]
		if visit {
			continue
		}

		visited[mapKey] = 1
		if pos.row > 0 {
			queue = append(queue, coord{pos.row - 1, pos.col, pos.step + 1})
		}
		if pos.row < len(input)-1 {
			queue = append(queue, coord{pos.row + 1, pos.col, pos.step + 1})
		}
		if pos.col > 0 {
			queue = append(queue, coord{pos.row, pos.col - 1, pos.step + 1})
		}
		if pos.col < len(input[pos.row])-1 {
			queue = append(queue, coord{pos.row, pos.col + 1, pos.step + 1})
		}
	}

	return visited[getMapKey(end)]
}

func getMapKey(c coord) string {
	return strconv.Itoa(c.row) + strconv.Itoa(c.col)
}

type coord struct {
	row  int
	col  int
	step int
}
