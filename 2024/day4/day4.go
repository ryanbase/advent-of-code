package main

import (
	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	filename := utils.GetFileNameFromArgument()
	input := utils.ReadInputAsByteMatrix(filename)
	partOne(input)
	partTwo(input)
}

func partOne(input [][]byte) {
	count := 0

	for i, line := range input {
		for j, char := range line {
			if char == 'X' {
				// Check up
				if i >= 3 {
					if input[i-1][j] == 'M' && input[i-2][j] == 'A' && input[i-3][j] == 'S' {
						count++
					}
				}

				// Check down
				if i <= len(input)-4 {
					if input[i+1][j] == 'M' && input[i+2][j] == 'A' && input[i+3][j] == 'S' {
						count++
					}
				}

				// Check left
				if j >= 3 {
					if input[i][j-1] == 'M' && input[i][j-2] == 'A' && input[i][j-3] == 'S' {
						count++
					}
				}

				// Check right
				if j <= len(input[i])-4 {
					if input[i][j+1] == 'M' && input[i][j+2] == 'A' && input[i][j+3] == 'S' {
						count++
					}
				}

				// Check up left
				if i >= 3 && j >= 3 {
					if input[i-1][j-1] == 'M' && input[i-2][j-2] == 'A' && input[i-3][j-3] == 'S' {
						count++
					}
				}

				// Check up right
				if i >= 3 && j <= len(input[i])-4 {
					if input[i-1][j+1] == 'M' && input[i-2][j+2] == 'A' && input[i-3][j+3] == 'S' {
						count++
					}
				}

				// Check down left
				if i <= len(input)-4 && j >= 3 {
					if input[i+1][j-1] == 'M' && input[i+2][j-2] == 'A' && input[i+3][j-3] == 'S' {
						count++
					}
				}

				// Check down right
				if i <= len(input)-4 && j <= len(input[i])-4 {
					if input[i+1][j+1] == 'M' && input[i+2][j+2] == 'A' && input[i+3][j+3] == 'S' {
						count++
					}
				}
			}
		}
	}

	println(count)
}

func partTwo(input [][]byte) {
	count := 0

	for i, line := range input {
		for j, char := range line {
			if i == 0 || i == len(input)-1 || j == 0 || j == len(input[i])-1 {
				continue
			}
			if char == 'A' {
				if ((input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M')) &&
					((input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S') || (input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M')) {
					count++
				}
			}
		}
	}

	println(count)
}
