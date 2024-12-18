package main

import (
	"bufio"
	"os"
	"slices"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type state struct {
	walls map[point]struct{}
	boxes map[point]struct{}
	robot point
}

type point struct {
	row int
	col int
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	state, moves := readInput(filename)
	part1(state, moves)
}

func readInput(filename string) (*state, []byte) {
	defer utils.TimeTrack(time.Now())

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	state := state{make(map[point]struct{}), make(map[point]struct{}), point{}}
	moves := []byte{}
	first := true
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			first = false
			continue
		}
		bytes := []byte(line)
		if first {
			for col, val := range bytes {
				if val == '#' {
					state.walls[point{row, col}] = struct{}{}
				} else if val == '@' {
					state.robot = point{row, col}
				} else if val == 'O' {
					state.boxes[point{row, col}] = struct{}{}
				}
			}
		} else {
			moves = append(moves, bytes...)
		}
		row++
	}
	return &state, moves
}

func part1(state *state, moves []byte) {
	defer utils.TimeTrack(time.Now())
	for _, move := range moves {
		// fmt.Println(state.robot, string(move), state.boxes)
		if move == '<' {
			moveRobot(state, 0, -1)
		} else if move == '^' {
			moveRobot(state, -1, 0)
		} else if move == '>' {
			moveRobot(state, 0, 1)
		} else if move == 'v' {
			moveRobot(state, 1, 0)
		}
	}
	res := 0
	for box := range state.boxes {
		res += (box.row * 100) + box.col
	}
	println(res)
}

func moveRobot(state *state, rowMove int, colMove int) {
	newPoint := point{state.robot.row + rowMove, state.robot.col + colMove}
	if _, ok := state.walls[newPoint]; ok {
		return
	}
	if _, ok := state.boxes[newPoint]; !ok {
		state.robot = newPoint
		return
	}
	boxes := []point{}
	lastBox := newPoint
	_, isBox := state.boxes[newPoint]
	for isBox {
		boxes = append(boxes, newPoint)
		lastBox = newPoint
		newPoint = point{newPoint.row + rowMove, newPoint.col + colMove}
		_, isBox = state.boxes[newPoint]
	}
	if _, ok := state.walls[point{lastBox.row + rowMove, lastBox.col + colMove}]; ok {
		return
	}
	slices.Reverse(boxes)
	for _, box := range boxes {
		delete(state.boxes, box)
		state.boxes[point{box.row + rowMove, box.col + colMove}] = struct{}{}
	}
	state.robot = point{state.robot.row + rowMove, state.robot.col + colMove}
}
