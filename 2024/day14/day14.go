package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type robot struct {
	x  int
	y  int
	vx int
	vy int
}

var WIDTH = 101
var HEIGHT = 103

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	robots := readInput(filename)
	part1(robots, 100)
	part2(robots)
}

func readInput(filename string) []robot {
	defer utils.TimeTrack(time.Now())

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	robots := []robot{}
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		pos := strings.Split(strings.Split(split[0], "=")[1], ",")
		vel := strings.Split(strings.Split(split[1], "=")[1], ",")
		x, _ := strconv.Atoi(pos[0])
		y, _ := strconv.Atoi(pos[1])
		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])
		robot := robot{x, y, vx, vy}
		robots = append(robots, robot)
	}
	return robots
}

func part1(robots []robot, seconds int) {
	defer utils.TimeTrack(time.Now())
	sf := calculateSafetyFactor(robots, seconds)
	println(sf)
}

func part2(robots []robot) {
	defer utils.TimeTrack(time.Now())

	sf := calculateSafetyFactor(robots, 1)
	min := []int{1, sf}
	for i := 2; i < 10000; i++ {
		sf = calculateSafetyFactor(robots, i)
		if sf < min[1] {
			min[0] = i
			min[1] = sf
		}
	}

	seconds := min[0]

	grid := [][]byte{}
	for range HEIGHT {
		row := []byte{}
		for range WIDTH {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}

	for _, robot := range robots {
		robot.x = (robot.x + (robot.vx * seconds)) % WIDTH
		robot.y = (robot.y + (robot.vy * seconds)) % HEIGHT
		if robot.x < 0 {
			robot.x = WIDTH + robot.x
		}
		if robot.y < 0 {
			robot.y = HEIGHT + robot.y
		}
		grid[robot.y][robot.x] = '+'
	}

	printGrid(grid)
	println(seconds)
}

func calculateSafetyFactor(robots []robot, seconds int) int {
	quadrants := make([]int, 4)
	for _, robot := range robots {
		x := (robot.x + (robot.vx * seconds)) % WIDTH
		y := (robot.y + (robot.vy * seconds)) % HEIGHT
		if x < 0 {
			x = WIDTH + x
		}
		if y < 0 {
			y = HEIGHT + y
		}
		if y < int(HEIGHT/2) {
			if x < int(WIDTH/2) {
				quadrants[0]++
			} else if x > int(WIDTH/2) {
				quadrants[1]++
			}
		} else if y > int(HEIGHT/2) {
			if x < int(WIDTH/2) {
				quadrants[2]++
			} else if x > int(WIDTH/2) {
				quadrants[3]++
			}
		}
	}
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, val := range row {
			print(string(val))
		}
		println()
	}
}
