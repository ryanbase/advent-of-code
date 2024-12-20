package main

import (
	"fmt"
	"maps"
	"math"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type point struct {
	row int
	col int
	dir byte
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	part1(filename)
}

func part1(filename string) {
	defer utils.TimeTrack(time.Now())

	input := utils.ReadInputAsByteMatrix(filename)
	start := point{-1, -1, 'E'}
	end := point{-1, -1, 'E'}

	for i, row := range input {
		for j, val := range row {
			if val == 'S' {
				start = point{i, j, 'E'}
			} else if val == 'E' {
				end = point{i, j, 'E'}
			}
		}
	}

	visited := make(map[point]struct{})
	res := doMaze(input, start, end, visited, 0, 0)
	println(res)

}

func doMaze(maze [][]byte, curr point, end point, visited map[point]struct{}, points int, steps int) int {
	if curr == end {
		return points
	}

	fmt.Println(steps, curr)

	if steps > len(maze)*len(maze[0]) {
		return math.MaxInt
	}

	// if _, ok := visited[curr]; ok {
	// 	return math.MaxInt
	// }

	// visited[curr] = struct{}{}

	res := math.MaxInt
	if curr.dir == 'E' {
		if maze[curr.row][curr.col+1] == '.' || maze[curr.row][curr.col+1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col + 1, 'E'}, end, maps.Clone(visited), points+1, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row-1][curr.col] == '.' || maze[curr.row-1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'N'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row+1][curr.col] == '.' || maze[curr.row+1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'S'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
	} else if curr.dir == 'S' {
		if maze[curr.row+1][curr.col] == '.' || maze[curr.row+1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row + 1, curr.col, 'S'}, end, maps.Clone(visited), points+1, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row][curr.col-1] == '.' || maze[curr.row][curr.col-1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'W'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row][curr.col+1] == '.' || maze[curr.row+1][curr.col+1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'E'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
	} else if curr.dir == 'W' {
		if maze[curr.row][curr.col-1] == '.' || maze[curr.row][curr.col-1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col - 1, 'W'}, end, maps.Clone(visited), points+1, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row-1][curr.col] == '.' || maze[curr.row-1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'N'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row+1][curr.col] == '.' || maze[curr.row+1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'S'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
	} else if curr.dir == 'N' {
		if maze[curr.row-1][curr.col] == '.' || maze[curr.row-1][curr.col] == 'E' {
			tmp := doMaze(maze, point{curr.row - 1, curr.col, 'N'}, end, maps.Clone(visited), points+1, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row][curr.col-1] == '.' || maze[curr.row][curr.col-1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'W'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
		if maze[curr.row][curr.col+1] == '.' || maze[curr.row][curr.col+1] == 'E' {
			tmp := doMaze(maze, point{curr.row, curr.col, 'E'}, end, maps.Clone(visited), points+1000, steps+1)
			if tmp < res {
				res = tmp
			}
		}
	}

	return res
}
