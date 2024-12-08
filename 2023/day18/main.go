package main

import (
	"bufio"
	"math"
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

	input := make([]Input, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		dir := split[0]
		dist, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		color := strings.ReplaceAll(split[2], "(", "")
		color = strings.ReplaceAll(color, ")", "")
		input = append(input, Input{dir, int64(dist), color})
	}

	grid := createGrid(input)

	count := 0
	floodFill(22, 207, grid)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] != "" {
				count++
			}
		}
	}

	// printGrid(grid)

	println(count)

}

func partTwo() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := make([]Input, 0)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		hex := strings.ReplaceAll(split[2], "(", "")
		hex = strings.ReplaceAll(hex, ")", "")
		dist, err := strconv.ParseInt(hex[1:6], 16, 64)
		if err != nil {
			panic(err)
		}
		dirNum := hex[6:]
		var dir string = ""
		switch dirNum {
		case "0":
			dir = "R"
			break
		case "1":
			dir = "D"
			break
		case "2":
			dir = "L"
			break
		case "3":
			dir = "U"
			break
		}
		input = append(input, Input{dir, dist, "#"})
	}

	println(calcArea(input))
}

func calcArea(input []Input) int64 {
	var x int64 = 0
	var y int64 = 0
	coords := []Coord{}
	for _, in := range input {
		if in.dir == "R" {
			x += in.dist
		} else if in.dir == "L" {
			x -= in.dist
		} else if in.dir == "U" {
			y += in.dist
		} else if in.dir == "D" {
			y -= in.dist
		}
		coords = append(coords, Coord{x, y})
		// fmt.Println(in, x, y)
	}

	var sum1 int64 = 0
	var sum2 int64 = 0

	for i := 0; i < len(coords)-1; i++ {
		sum1 = sum1 + coords[i].x*coords[i+1].y
		sum2 = sum2 + coords[i].y*coords[i+1].x
	}
	sum1 = sum1 + coords[len(coords)-1].x*coords[0].y
	sum2 = sum2 + coords[0].x*coords[len(coords)-1].y

	return int64(math.Abs(float64(sum1-sum2)) / 2)
}

func printGrid(grid [][]string) {
	for _, bytes := range grid {
		for _, str := range bytes {
			if str == "X" {
				print("X")
			} else if str != "" {
				print("#")
			} else {
				print(".")
			}
		}
		println()
	}
}

func createGrid(input []Input) [][]string {
	grid := make([][]string, 0)
	grid = append(grid, []string{""})

	x := 0
	y := 0
	width := 1

	for _, line := range input {

		// println(dir, dist, color)
		dir := line.dir
		dist := line.dist
		color := line.color

		if line.dir == "U" {
			for i := 0; i < int(dist); i++ {
				y--
				if y < 0 {
					grid = addRowsToTop(grid, width)
					y = 0
				}
				grid[y][x] = color
			}
		} else if dir == "R" {
			for i := 0; i < int(dist); i++ {
				x++
				if x > width-1 {
					grid = addColToRight(grid)
					width++
				}
				grid[y][x] = color
			}
		} else if dir == "D" {
			for i := 0; i < int(dist); i++ {
				y++
				if y > len(grid)-1 {
					grid = addRowToBottom(grid, width)
				}
				grid[y][x] = color
			}
		} else if dir == "L" {
			for i := 0; i < int(dist); i++ {
				x--
				if x < 0 {
					grid = addColToLeft(grid)
					x = 0
				}
				grid[y][x] = color
			}
		}
	}
	return grid
}

func addPadding(grid [][]string) [][]string {
	grid = addRowsToTop(grid, len(grid[0]))
	grid = addRowToBottom(grid, len(grid[0]))
	grid = addColToLeft(grid)
	grid = addColToRight(grid)
	return grid
}

func floodFill(row int, col int, grid [][]string) {
	// println(string(pipes[row][col]), exists)
	if grid[row][col] != "" {
		return
	}
	if grid[row][col] == "" {
		grid[row][col] = "X"
	}
	if row > 0 {
		floodFill(row-1, col, grid)
	}
	if row < len(grid)-1 {
		floodFill(row+1, col, grid)
	}
	if col > 0 {
		floodFill(row, col-1, grid)
	}
	if col < len(grid[0])-1 {
		floodFill(row, col+1, grid)
	}
}

func addRowsToTop(grid [][]string, rowWidth int) [][]string {
	newRow := make([]string, 0)
	for i := 0; i < rowWidth; i++ {
		newRow = append(newRow, "")
	}
	grid = append([][]string{newRow}, grid...)
	return grid
}

func addRowToBottom(grid [][]string, rowWidth int) [][]string {
	newRow := make([]string, 0)
	for i := 0; i < rowWidth; i++ {
		newRow = append(newRow, "")
	}
	grid = append(grid, newRow)
	return grid
}

func addColToLeft(grid [][]string) [][]string {
	for i := 0; i < len(grid); i++ {
		grid[i] = append([]string{""}, grid[i]...)
	}
	return grid
}

func addColToRight(grid [][]string) [][]string {
	for i := 0; i < len(grid); i++ {
		grid[i] = append(grid[i], "")
	}
	return grid
}

type Input struct {
	dir   string
	dist  int64
	color string
}

type Coord struct {
	x int64
	y int64
}
