package main

import (
	"bufio"
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

	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	// history := make(map[string]int)

	// total := nextLoc(Coord{0, 0}, 'R', input, history)
	total := getEnergized(input)

	println(total)

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

// func nextLoc(current Coord, direction byte, input []string, history map[string]int) int {
// 	if current.row >= len(input) || current.row < 0 || current.col >= len(input[0]) || current.col < 0 {
// 		return 0
// 	}
// 	// fmt.Println(current, createMapKey(current, direction))
// 	_, looped := history[createMapKey(current, direction)]
// 	if looped {
// 		return 0
// 	}
// 	count := 1
// 	_, energized := history[createMapKey(current, 'E')]
// 	if energized {
// 		count = 0
// 	}
// 	history[createMapKey(current, direction)] = 1
// 	history[createMapKey(current, 'E')] = 1
// 	currentChar := input[current.row][current.col]
// 	// println(string(currentChar), string(direction))
// 	if currentChar == '.' {
// 		if direction == 'U' {
// 			return count + nextLoc(Coord{current.row - 1, current.col}, direction, input, history)
// 		} else if direction == 'D' {
// 			return count + nextLoc(Coord{current.row + 1, current.col}, direction, input, history)
// 		} else if direction == 'L' {
// 			return count + nextLoc(Coord{current.row, current.col - 1}, direction, input, history)
// 		} else if direction == 'R' {
// 			return count + nextLoc(Coord{current.row, current.col + 1}, direction, input, history)
// 		}
// 	} else if currentChar == '/' {
// 		if direction == 'U' {
// 			return count + nextLoc(Coord{current.row, current.col + 1}, 'R', input, history)
// 		} else if direction == 'D' {
// 			return count + nextLoc(Coord{current.row, current.col - 1}, 'L', input, history)
// 		} else if direction == 'L' {
// 			return count + nextLoc(Coord{current.row + 1, current.col}, 'D', input, history)
// 		} else if direction == 'R' {
// 			return count + nextLoc(Coord{current.row - 1, current.col}, 'U', input, history)
// 		}
// 	} else if currentChar == '\\' {
// 		if direction == 'U' {
// 			return count + nextLoc(Coord{current.row, current.col - 1}, 'L', input, history)
// 		} else if direction == 'D' {
// 			return count + nextLoc(Coord{current.row, current.col + 1}, 'R', input, history)
// 		} else if direction == 'L' {
// 			return count + nextLoc(Coord{current.row - 1, current.col}, 'U', input, history)
// 		} else if direction == 'R' {
// 			return count + nextLoc(Coord{current.row + 1, current.col}, 'D', input, history)
// 		}
// 	} else if currentChar == '-' {
// 		if direction == 'U' || direction == 'D' {
// 			return count + nextLoc(Coord{current.row, current.col - 1}, 'L', input, history) + nextLoc(Coord{current.row, current.col + 1}, 'R', input, history)
// 		} else if direction == 'L' {
// 			return count + nextLoc(Coord{current.row, current.col - 1}, 'L', input, history)
// 		} else if direction == 'R' {
// 			return count + nextLoc(Coord{current.row, current.col + 1}, 'R', input, history)
// 		}
// 	} else if currentChar == '|' {
// 		if direction == 'L' || direction == 'R' {
// 			return count + nextLoc(Coord{current.row - 1, current.col}, 'U', input, history) + nextLoc(Coord{current.row + 1, current.col}, 'D', input, history)
// 		} else if direction == 'U' {
// 			return count + nextLoc(Coord{current.row - 1, current.col}, 'U', input, history)
// 		} else if direction == 'D' {
// 			return count + nextLoc(Coord{current.row + 1, current.col}, 'D', input, history)
// 		}
// 	}

// 	return 0
// }

func getEnergized(input []string) int {
	visited := make(map[string]int)
	energized := make(map[string]int)
	queue := make([]Coord, 0)
	queue = append(queue, Coord{0, 0, 'R'})

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		energized[createMapKey(Coord{current.row, current.col, 'E'})] = 1
		currentChar := input[current.row][current.col]

		nextCoords := make([]Coord, 0)

		if currentChar == '.' {
			if current.dir == 'U' {
				nextCoords = append(nextCoords, Coord{current.row - 1, current.col, 'U'})
			} else if current.dir == 'D' {
				nextCoords = append(nextCoords, Coord{current.row + 1, current.col, 'D'})
			} else if current.dir == 'L' {
				nextCoords = append(nextCoords, Coord{current.row, current.col - 1, 'L'})
			} else if current.dir == 'R' {
				nextCoords = append(nextCoords, Coord{current.row, current.col + 1, 'R'})
			}
		} else if currentChar == '/' {
			if current.dir == 'U' {
				nextCoords = append(nextCoords, Coord{current.row, current.col + 1, 'R'})
			} else if current.dir == 'D' {
				nextCoords = append(nextCoords, Coord{current.row, current.col - 1, 'L'})
			} else if current.dir == 'L' {
				nextCoords = append(nextCoords, Coord{current.row + 1, current.col, 'D'})
			} else if current.dir == 'R' {
				nextCoords = append(nextCoords, Coord{current.row - 1, current.col, 'U'})
			}
		} else if currentChar == '\\' {
			if current.dir == 'U' {
				nextCoords = append(nextCoords, Coord{current.row, current.col - 1, 'L'})
			} else if current.dir == 'D' {
				nextCoords = append(nextCoords, Coord{current.row, current.col + 1, 'R'})
			} else if current.dir == 'L' {
				nextCoords = append(nextCoords, Coord{current.row - 1, current.col, 'U'})
			} else if current.dir == 'R' {
				nextCoords = append(nextCoords, Coord{current.row + 1, current.col, 'D'})
			}
		} else if currentChar == '-' {
			if current.dir == 'U' || current.dir == 'D' {
				nextCoords = append(nextCoords, Coord{current.row, current.col - 1, 'L'})
				nextCoords = append(nextCoords, Coord{current.row, current.col + 1, 'R'})
			} else if current.dir == 'L' {
				nextCoords = append(nextCoords, Coord{current.row, current.col - 1, 'L'})
			} else if current.dir == 'R' {
				nextCoords = append(nextCoords, Coord{current.row, current.col + 1, 'R'})
			}
		} else if currentChar == '|' {
			if current.dir == 'L' || current.dir == 'R' {
				nextCoords = append(nextCoords, Coord{current.row - 1, current.col, 'U'})
				nextCoords = append(nextCoords, Coord{current.row + 1, current.col, 'D'})
			} else if current.dir == 'U' {
				nextCoords = append(nextCoords, Coord{current.row - 1, current.col, 'U'})
			} else if current.dir == 'D' {
				nextCoords = append(nextCoords, Coord{current.row + 1, current.col, 'D'})
			}
		}

		for _, coord := range nextCoords {
			_, exists := visited[createMapKey(coord)]
			if !exists && coord.row >= 0 && coord.row < len(input) && coord.col >= 0 && coord.col < len(input[0]) {
				queue = append(queue, coord)
				visited[createMapKey(coord)] = 1
			}
		}

	}

	return len(energized)
}

type Coord struct {
	row int
	col int
	dir byte
}

func createMapKey(coord Coord) string {
	return strconv.Itoa(coord.row) + strconv.Itoa(coord.col) + string(coord.dir)
}
