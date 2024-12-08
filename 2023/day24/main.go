package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := []Input{}

	for scanner.Scan() {
		line := scanner.Text()
		posVel := strings.Split(line, " @ ")
		pos := strings.Split(posVel[0], ", ")
		x, err := strconv.Atoi(strings.TrimSpace(pos[0]))
		checkErr(err)
		y, err := strconv.Atoi(strings.TrimSpace(pos[1]))
		checkErr(err)
		z, err := strconv.Atoi(strings.TrimSpace(pos[2]))
		checkErr(err)
		vel := strings.Split(posVel[1], ", ")
		vx, err := strconv.Atoi(strings.TrimSpace(vel[0]))
		checkErr(err)
		vy, err := strconv.Atoi(strings.TrimSpace(vel[1]))
		checkErr(err)
		vz, err := strconv.Atoi(strings.TrimSpace(vel[2]))
		checkErr(err)
		input = append(input, Input{x, y, z, vx, vy, vz})
	}

	// for _, in := range input {

	// }

}

func partTwo() {
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Input struct {
	x1, y1, z1, vx, vy, vz int
}
