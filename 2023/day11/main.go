package main

import (
	"bufio"
	"math"
	"os"
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

	universe := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if allSame(line, '.') {
			newLine := strings.Clone(line)
			universe = append(universe, newLine)
		}
		universe = append(universe, line)
	}

	i := 0
	max := len(universe[0])
	for i < max {
		allSame := true
		for _, str := range universe {
			if str[i] != '.' {
				allSame = false
				break
			}
		}
		if allSame {
			for row, str := range universe {
				universe[row] = str[:i] + "." + str[i:]
			}
			max = len(universe[0])
			i++
		}
		i++
	}

	galaxies := make([]Coord, 0)
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == '#' {
				galaxies = append(galaxies, Coord{row, col})
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += int(math.Abs(float64(galaxies[i].row - galaxies[j].row)))
			total += int(math.Abs(float64(galaxies[i].col - galaxies[j].col)))
		}
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

	universe := make([]string, 0)

	bigRows := make([]int, 0)
	bigCols := make([]int, 0)

	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if allSame(line, '.') {
			bigRows = append(bigRows, row)
		}
		row++
		universe = append(universe, line)
	}

	i := 0
	for i < len(universe[0]) {
		allSame := true
		for _, str := range universe {
			if str[i] != '.' {
				allSame = false
				break
			}
		}
		if allSame {
			bigCols = append(bigCols, i)
		}
		i++
	}

	galaxies := make([]Coord, 0)
	for row := 0; row < len(universe); row++ {
		for col := 0; col < len(universe[row]); col++ {
			if universe[row][col] == '#' {
				galaxies = append(galaxies, Coord{row, col})
			}
		}
	}

	total := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			minRow := 0
			maxRow := 0
			minCol := 0
			maxCol := 0
			if galaxies[i].row >= galaxies[j].row {
				minRow = galaxies[j].row
				maxRow = galaxies[i].row
			} else {
				minRow = galaxies[i].row
				maxRow = galaxies[j].row
			}
			if galaxies[i].col >= galaxies[j].col {
				minCol = galaxies[j].col
				maxCol = galaxies[i].col
			} else {
				minCol = galaxies[i].col
				maxCol = galaxies[j].col
			}

			total += maxRow - minRow
			total += maxCol - minCol

			for _, bigRow := range bigRows {
				if minRow < bigRow && maxRow > bigRow {
					total += 999999
				}
			}
			for _, bigCol := range bigCols {
				if minCol < bigCol && maxCol > bigCol {
					total += 999999
				}
			}

		}
	}

	println(total)
}

func allSame(str string, char byte) bool {
	for i := 0; i < len(str); i++ {
		if str[i] != char {
			return false
		}
	}
	return true
}

type Pair struct {
	first  Coord
	second Coord
}

type Coord struct {
	row int
	col int
}
