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
	f, err := os.Open("input.txt")
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

	println(count(input, start, 64))
	// println(count2(input, start, 6))

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
	for row, str := range input {
		for col, char := range str {
			if char == 'S' {
				return coord{row, col, 0}, nil
			}
		}
	}
	return coord{-1, -1, -1}, errors.New("Starting position not found")
}

// func count(input []string, pos coord, numSteps int) int {
// 	spaces := make(map[string]struct{})
// 	visited := make(map[string]struct{})
// 	countR(input, pos, numSteps, 0, spaces, visited)

// 	// for row, str := range input {
// 	// 	for col := range str {
// 	// 		_, ok := spaces[getMapKey(coord{row, col})]
// 	// 		if ok {
// 	// 			input[row] = replaceAtIndex(input[row], 'O', col)
// 	// 		}
// 	// 	}
// 	// }

// 	// for _, str := range input {
// 	// 	println(str)
// 	// }

// 	return len(spaces)
// }

// func countR(input []string, pos coord, numSteps int, stepsTaken int, spaces map[string]struct{}, visited map[string]struct{}) {
// 	visitedKey := getMapKey(pos)
// 	_, visit := visited[visitedKey]
// 	if visit || input[pos.row][pos.col] == '#' {
// 		return
// 	}
// 	if stepsTaken == numSteps {
// 		spaces[getMapKey(pos)] = struct{}{}
// 		return
// 	}

// 	visited[visitedKey] = struct{}{}
// 	if pos.row > 0 {
// 		countR(input, coord{pos.row - 1, pos.col, stepsTaken + 1}, numSteps, stepsTaken+1, spaces, visited)
// 	}
// 	if pos.row < len(input)-1 {
// 		countR(input, coord{pos.row + 1, pos.col, stepsTaken + 1}, numSteps, stepsTaken+1, spaces, visited)
// 	}
// 	if pos.col > 0 {
// 		countR(input, coord{pos.row, pos.col - 1, stepsTaken + 1}, numSteps, stepsTaken+1, spaces, visited)
// 	}
// 	if pos.col < len(input[0])-1 {
// 		countR(input, coord{pos.row, pos.col + 1, stepsTaken + 1}, numSteps, stepsTaken+1, spaces, visited)
// 	}

// }

func count(input []string, start coord, stepsNeeded int) int {
	spaces := make(map[string]struct{})
	visited := make(map[string]struct{})
	queue := []coord{start}

	for len(queue) > 0 {
		pos := queue[0]
		queue = queue[1:]

		if input[pos.row][pos.col] == '#' {
			continue
		}

		if pos.step%2 == stepsNeeded%2 {
			spaces[getSpacesMapKey(pos)] = struct{}{}
			if pos.step == stepsNeeded {
				continue
			}
		}

		visitedKey := getVisitedMapKey(pos)
		_, visit := visited[visitedKey]
		if visit {
			continue
		}

		// if pos.step == stepsNeeded {
		// 	spaces[getSpacesMapKey(pos)] = struct{}{}
		// 	continue
		// }

		visited[visitedKey] = struct{}{}
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

	return len(spaces)
}

// func count2(input []string, start coord, stepsNeeded int) int {
// 	fmt.Println(start)
// 	count := 0
// 	for row := 0; row < len(input); row++ {
// 		for col := 0; col < len(input[row]); col++ {
// 			dist := int(math.Abs(float64(start.row)-float64(row)) + math.Abs(float64(start.col)-float64(col)))
// 			println(row, col, dist)
// 			if dist <= stepsNeeded {
// 				if dist%2 == stepsNeeded%2 && input[row][col] != '#' {
// 					count++
// 				}
// 			}
// 		}
// 	}
// 	return count
// }

func getSpacesMapKey(c coord) string {
	return strconv.Itoa(c.row) + strconv.Itoa(c.col)
}

func getVisitedMapKey(c coord) string {
	return strconv.Itoa(c.row) + strconv.Itoa(c.col) + strconv.Itoa(c.step)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

type coord struct {
	row  int
	col  int
	step int
}
