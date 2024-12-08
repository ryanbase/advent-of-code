package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]

	input, start := readInput(filename)
	part1(input, start)
	part2BruteForce(input, start)
}

func readInput(filename string) ([][]byte, []int) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := [][]byte{}
	start := []int{-1, -1}
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []byte(line))
		if start[0] == -1 {
			for j, char := range line {
				if char == '^' {
					start[0] = i
					start[1] = j
				}
			}
			i++
		}
	}
	return input, start
}

func createKey(i int, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func createDirKey(i int, j int, dir rune) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j) + "," + string(dir)
}

// Returns next value or Z if unable
func getNext(input [][]byte, dir rune, i int, j int) byte {
	next := byte('Z')
	if dir == 'N' && i > 0 {
		next = input[i-1][j]
	} else if dir == 'E' && j < len(input[i])-1 {
		next = input[i][j+1]
	} else if dir == 'S' && i < len(input)-1 {
		next = input[i+1][j]
	} else if dir == 'W' && j > 0 {
		next = input[i][j-1]
	}
	return next
}

func part1(input [][]byte, start []int) {
	visited := make(map[string]bool)
	i := start[0]
	j := start[1]
	dir := 'N'
	for i >= 0 && i < len(input) && j >= 0 && j < len(input[i]) {
		visited[createKey(i, j)] = true
		next := getNext(input, dir, i, j)

		if next == 'Z' {
			break
		}
		if next == '.' || next == '^' {
			if dir == 'N' {
				i = i - 1
			} else if dir == 'E' {
				j = j + 1
			} else if dir == 'S' {
				i = i + 1
			} else if dir == 'W' {
				j = j - 1
			}
		} else if next == '#' {
			if dir == 'N' {
				dir = 'E'
			} else if dir == 'E' {
				dir = 'S'
			} else if dir == 'S' {
				dir = 'W'
			} else if dir == 'W' {
				dir = 'N'
			}
		}
	}

	println(len(visited))
}

func repeatPath(input [][]byte, visited map[string]bool, i int, j int, dir rune) bool {
	if dir == 'N' {
		i--
		for i >= 0 && input[i][j] != '#' {
			if visited[createDirKey(i, j, dir)] {
				return true
			}
			i--
		}
	} else if dir == 'E' {
		j++
		for j <= len(input[i])-1 && input[i][j] != '#' {
			if visited[createDirKey(i, j, dir)] {
				return true
			}
			j++
		}
	} else if dir == 'S' {
		i++
		for i <= len(input)-1 && input[i][j] != '#' {
			if visited[createDirKey(i, j, dir)] {
				return true
			}
			i++
		}
	} else if dir == 'W' {
		j--
		for j >= 0 && input[i][j] != '#' {
			if visited[createDirKey(i, j, dir)] {
				return true
			}
			j--
		}
	}
	return false
}

func part2(input [][]byte, start []int) {
	visited := make(map[string]bool)
	obstructions := make(map[string]bool)
	i := start[0]
	j := start[1]
	dir := 'N'
	for i >= 0 && i < len(input) && j >= 0 && j < len(input[i]) {
		next := getNext(input, dir, i, j)

		if next != '#' {
			if dir == 'N' && repeatPath(input, visited, i, j, 'E') {
				obstructions[createKey(i-1, j)] = true
			} else if dir == 'E' && repeatPath(input, visited, i, j, 'S') {
				obstructions[createKey(i, j+1)] = true
			} else if dir == 'S' && repeatPath(input, visited, i, j, 'W') {
				obstructions[createKey(i+1, j)] = true
			} else if dir == 'W' && repeatPath(input, visited, i, j, 'N') {
				obstructions[createKey(i, j-1)] = true
			}
		}

		visited[createDirKey(i, j, dir)] = true

		if next == 'Z' {
			break
		}
		if next == '.' || next == '^' {
			if dir == 'N' {
				i = i - 1
			} else if dir == 'E' {
				j = j + 1
			} else if dir == 'S' {
				i = i + 1
			} else if dir == 'W' {
				j = j - 1
			}
		} else if next == '#' {
			if dir == 'N' {
				dir = 'E'
			} else if dir == 'E' {
				dir = 'S'
			} else if dir == 'S' {
				dir = 'W'
			} else if dir == 'W' {
				dir = 'N'
			}
		}
	}

	for key := range obstructions {
		println(key)
	}

	// for key := range visited {
	// 	println(key)
	// }

	println(len(obstructions))
}
func isInLoop(input [][]byte, start []int) bool {
	visited := make(map[string]bool)
	i := start[0]
	j := start[1]
	dir := 'N'
	steps := 0
	for i >= 0 && i < len(input) && j >= 0 && j < len(input[i]) {
		visited[createKey(i, j)] = true
		next := getNext(input, dir, i, j)

		if next == 'Z' {
			break
		}
		if next == '.' || next == '^' {
			if dir == 'N' {
				i = i - 1
			} else if dir == 'E' {
				j = j + 1
			} else if dir == 'S' {
				i = i + 1
			} else if dir == 'W' {
				j = j - 1
			}
		} else if next == '#' {
			if dir == 'N' {
				dir = 'E'
			} else if dir == 'E' {
				dir = 'S'
			} else if dir == 'S' {
				dir = 'W'
			} else if dir == 'W' {
				dir = 'N'
			}
		}
		steps++
		if steps > len(input)*len(input[0]) {
			return true
		}
	}
	return false
}

func part2BruteForce(input [][]byte, start []int) {
	obstructions := make(map[string]bool)
	for i, row := range input {
		for j := range row {
			if input[i][j] != '#' && input[i][j] != '^' {
				input[i][j] = '#'
				if isInLoop(input, start) {
					obstructions[createKey(i, j)] = true
				}
				input[i][j] = '.'
			}
		}
	}

	println(len(obstructions))
}
