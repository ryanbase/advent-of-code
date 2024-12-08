package main

import (
	"bufio"
	"errors"
	"os"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	pipes := make([]string, 0)
	pipes2 := make([]string, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		pipes = append(pipes, line)
		pipes2 = append(pipes2, line)
	}

	startPos, err := getStartPos(pipes)
	if err != nil {
		panic(err)
	}

	prevPos := startPos
	currentPos := startPos

	loopLength := 1

	if canGoUp(currentPos, pipes) {
		currentPos = Coord{currentPos.row - 1, currentPos.col}
	} else if canGoDown(currentPos, pipes) {
		currentPos = Coord{currentPos.row + 1, currentPos.col}
	} else if canGoLeft(currentPos, pipes) {
		currentPos = Coord{currentPos.row, currentPos.col - 1}
	}

	for true {
		if currentPos == startPos {
			break
		}

		loopLength++
		// currentPipe := pipes[currentPos.row][currentPos.col]
		// println(string(currentPipe))

		tmpPos := currentPos
		currentPos = getNextPipe(currentPos, prevPos, pipes)
		prevPos = tmpPos

		if prevPos.row == currentPos.row-1 {
			pipes2[prevPos.row] = replaceAtIndex(pipes2[prevPos.row], 'D', prevPos.col)
		} else if prevPos.row == currentPos.row+1 {
			pipes2[prevPos.row] = replaceAtIndex(pipes2[prevPos.row], 'U', prevPos.col)
		} else if prevPos.col == currentPos.col-1 {
			pipes2[prevPos.row] = replaceAtIndex(pipes2[prevPos.row], 'R', prevPos.col)
		} else if prevPos.col == currentPos.col+1 {
			pipes2[prevPos.row] = replaceAtIndex(pipes2[prevPos.row], 'X', prevPos.col)
		}
	}

	println(loopLength / 2)

	// Expand graph
	expandedPipes := make([]string, 0)
	for _, row := range pipes2 {
		newRow := ""
		for i := 0; i < len(row); i++ {
			newRow = newRow + "Z" + string(row[i])
		}
		newRow += "Z"
		expandedPipes = append(expandedPipes, newRow)
		emptyRow := ""
		for i := 0; i < len(newRow); i++ {
			emptyRow += "Z"
		}
		expandedPipes = append(expandedPipes, emptyRow)
	}

	startPos, err = getStartPos(expandedPipes)
	if err != nil {
		panic(err)
	}

	// Replace S with U. I figured this out by looking at the input after figuring out the loop.
	expandedPipes[startPos.row] = replaceAtIndex(expandedPipes[startPos.row], 'U', startPos.col)
	currentPos = startPos
	// fmt.Println(currentPos)
	for true {
		// fmt.Println(currentPos, string(expandedPipes[currentPos.row][currentPos.col]))
		if expandedPipes[currentPos.row][currentPos.col] == 'U' {
			expandedPipes[currentPos.row-1] = replaceAtIndex(expandedPipes[currentPos.row-1], 'U', currentPos.col)
			currentPos = Coord{currentPos.row - 2, currentPos.col}
		} else if expandedPipes[currentPos.row][currentPos.col] == 'D' {
			expandedPipes[currentPos.row+1] = replaceAtIndex(expandedPipes[currentPos.row+1], 'D', currentPos.col)
			currentPos = Coord{currentPos.row + 2, currentPos.col}
		} else if expandedPipes[currentPos.row][currentPos.col] == 'X' {
			expandedPipes[currentPos.row] = replaceAtIndex(expandedPipes[currentPos.row], 'X', currentPos.col-1)
			currentPos = Coord{currentPos.row, currentPos.col - 2}
		} else if expandedPipes[currentPos.row][currentPos.col] == 'R' {
			expandedPipes[currentPos.row] = replaceAtIndex(expandedPipes[currentPos.row], 'R', currentPos.col+1)
			currentPos = Coord{currentPos.row, currentPos.col + 2}
		}

		if currentPos == startPos {
			break
		}
	}

	floodFill(0, 0, expandedPipes)

	validSpace := make(map[byte]struct{})
	validSpace['.'] = struct{}{}
	validSpace['7'] = struct{}{}
	validSpace['F'] = struct{}{}
	validSpace['J'] = struct{}{}
	validSpace['L'] = struct{}{}
	validSpace['-'] = struct{}{}
	validSpace['|'] = struct{}{}

	// Count number of valid spaces remaining
	count := 0
	for row := 0; row < len(expandedPipes); row++ {
		for col := 0; col < len(expandedPipes[row]); col++ {
			_, ok := validSpace[expandedPipes[row][col]]
			if ok {
				count++
			}
		}
	}
	println(count)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func floodFill(row int, col int, pipes []string) {
	replace := make(map[byte]struct{})
	replace['Z'] = struct{}{}
	replace['.'] = struct{}{}
	replace['7'] = struct{}{}
	replace['F'] = struct{}{}
	replace['J'] = struct{}{}
	replace['L'] = struct{}{}
	replace['-'] = struct{}{}
	replace['|'] = struct{}{}

	floodFillR(row, col, pipes, replace)
}

func floodFillR(row int, col int, pipes []string, m map[byte]struct{}) {
	_, exists := m[pipes[row][col]]
	// println(string(pipes[row][col]), exists)
	if exists {
		pipes[row] = replaceAtIndex(pipes[row], '0', col)
	} else {
		return
	}
	if row > 0 {
		floodFillR(row-1, col, pipes, m)
	}
	if row < len(pipes)-1 {
		floodFillR(row+1, col, pipes, m)
	}
	if col > 0 {
		floodFillR(row, col-1, pipes, m)
	}
	if col < len(pipes[row])-1 {
		floodFillR(row, col+1, pipes, m)
	}
}

func getStartPos(pipes []string) (Coord, error) {
	for i, row := range pipes {
		for j, col := range row {
			if col == 'S' {
				return Coord{i, j}, nil
			}
		}
	}
	return Coord{-1, -1}, errors.New("Starting position not found")
}

func getNextPipe(currentPos Coord, prevPos Coord, pipes []string) Coord {
	// Check up
	if prevPos.row != currentPos.row-1 && canGoUp(currentPos, pipes) {
		return Coord{currentPos.row - 1, currentPos.col}
	}

	// Check right
	if prevPos.col != currentPos.col+1 && canGoRight(currentPos, pipes) {
		return Coord{currentPos.row, currentPos.col + 1}
	}

	// Check down
	if prevPos.row != currentPos.row+1 && canGoDown(currentPos, pipes) {
		return Coord{currentPos.row + 1, currentPos.col}
	}

	// Check left
	if prevPos.col != currentPos.col-1 && canGoLeft(currentPos, pipes) {
		return Coord{currentPos.row, currentPos.col - 1}
	}

	return currentPos
}

func canGoUp(currentPos Coord, pipes []string) bool {
	if currentPos.row <= 0 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row-1][currentPos.col]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['|'] = struct{}{}
	validNextPipes['7'] = struct{}{}
	validNextPipes['F'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['|'] = struct{}{}
	validCurrentPipes['L'] = struct{}{}
	validCurrentPipes['J'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoDown(currentPos Coord, pipes []string) bool {
	if currentPos.row >= len(pipes) {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row+1][currentPos.col]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['|'] = struct{}{}
	validNextPipes['L'] = struct{}{}
	validNextPipes['J'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['|'] = struct{}{}
	validCurrentPipes['7'] = struct{}{}
	validCurrentPipes['F'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoLeft(currentPos Coord, pipes []string) bool {
	if currentPos.col <= 0 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row][currentPos.col-1]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['-'] = struct{}{}
	validNextPipes['L'] = struct{}{}
	validNextPipes['F'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['-'] = struct{}{}
	validCurrentPipes['7'] = struct{}{}
	validCurrentPipes['J'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

func canGoRight(currentPos Coord, pipes []string) bool {
	if currentPos.col >= len(pipes[currentPos.row])-1 {
		return false
	}
	currentPipe := pipes[currentPos.row][currentPos.col]
	nextPipe := pipes[currentPos.row][currentPos.col+1]

	validNextPipes := make(map[byte]struct{})
	validNextPipes['S'] = struct{}{}
	validNextPipes['-'] = struct{}{}
	validNextPipes['7'] = struct{}{}
	validNextPipes['J'] = struct{}{}

	validCurrentPipes := make(map[byte]struct{})
	validCurrentPipes['S'] = struct{}{}
	validCurrentPipes['-'] = struct{}{}
	validCurrentPipes['L'] = struct{}{}
	validCurrentPipes['F'] = struct{}{}

	_, nextValid := validNextPipes[nextPipe]
	_, currentValid := validCurrentPipes[currentPipe]
	return nextValid && currentValid
}

type Coord struct {
	row int
	col int
}
