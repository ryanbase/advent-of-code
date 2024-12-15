package main

import (
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type plot struct {
	row int
	col int
}

type region struct {
	perimiter int
	plots     []plot
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	input := utils.ReadInputAsByteMatrix(filename)
	regions := getRegions(input)
	part1(regions)
	// part2(input, regions)
}

func part1(regions []region) {
	defer utils.TimeTrack(time.Now())
	price := 0
	for _, region := range regions {
		price += len(region.plots) * region.perimiter
	}
	println(price)
}

// Part 2 not complete
func part2(input [][]byte, regions []region) {
	defer utils.TimeTrack(time.Now())
	price := 0
	for _, region := range regions {
		// TODO Find new perimiters
		// for i,row := range input {
		// 	for j, val := range row {

		// 	}
		// }
		price += len(region.plots) * region.perimiter
	}
	println(price)
}

func getRegions(input [][]byte) []region {
	defer utils.TimeTrack(time.Now())
	regions := []region{}
	visited := make(map[plot]struct{})
	for i, row := range input {
		for j, _ := range row {
			if _, ok := visited[plot{i, j}]; !ok {
				region := getRegion(input, i, j, visited)
				regions = append(regions, region)
			}
		}
	}
	return regions
}

func getRegion(input [][]byte, row int, col int, visited map[plot]struct{}) region {
	region := region{0, []plot{}}
	stack := []plot{plot{row, col}}
	for len(stack) > 0 {
		curr := stack[0]
		stack = stack[1:]
		if _, ok := visited[curr]; ok {
			continue
		}
		visited[curr] = struct{}{}
		region.plots = append(region.plots, curr)
		currValue := input[curr.row][curr.col]
		if curr.row > 0 {
			if input[curr.row-1][curr.col] == currValue {
				stack = append(stack, plot{curr.row - 1, curr.col})
			} else {
				region.perimiter++
			}
		} else {
			region.perimiter++
		}
		if curr.row < len(input)-1 {
			if input[curr.row+1][curr.col] == currValue {
				stack = append(stack, plot{curr.row + 1, curr.col})
			} else {
				region.perimiter++
			}
		} else {
			region.perimiter++
		}
		if curr.col > 0 {
			if input[curr.row][curr.col-1] == currValue {
				stack = append(stack, plot{curr.row, curr.col - 1})
			} else {
				region.perimiter++
			}
		} else {
			region.perimiter++
		}
		if curr.col < len(input[curr.row])-1 {
			if input[curr.row][curr.col+1] == currValue {
				stack = append(stack, plot{curr.row, curr.col + 1})
			} else {
				region.perimiter++
			}
		} else {
			region.perimiter++
		}
	}
	return region
}
