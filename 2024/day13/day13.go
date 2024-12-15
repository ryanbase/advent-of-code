package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type machine struct {
	aX     int
	aY     int
	bX     int
	bY     int
	prizeX int
	prizeY int
}

type cacheval struct {
	x          int
	y          int
	currTokens int
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	machines := readInput(filename)
	part1(machines)
	part2(machines)
}

func readInput(filename string) []machine {
	defer utils.TimeTrack(time.Now())

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	machines := []machine{}
	currMachine := machine{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		split := strings.Split(line, ": ")
		xyVals := strings.Split(split[1], ", ")
		x, _ := strconv.Atoi(xyVals[0][2:])
		y, _ := strconv.Atoi(xyVals[1][2:])
		if strings.Contains(split[0], "Button A") {
			currMachine.aX = x
			currMachine.aY = y
		}
		if strings.Contains(split[0], "Button B") {
			currMachine.bX = x
			currMachine.bY = y
		}
		if strings.Contains(split[0], "Prize") {
			currMachine.prizeX = x
			currMachine.prizeY = y
			machines = append(machines, currMachine)
			currMachine = machine{}
		}
	}
	return machines
}

func part1(machines []machine) {
	defer utils.TimeTrack(time.Now())
	tokens := 0
	for _, machine := range machines {
		tokens += calcMinTokens(machine)
	}
	println(tokens)
}

func part2(machines []machine) {
	defer utils.TimeTrack(time.Now())
	tokens := 0
	for _, machine := range machines {
		machine.prizeX += 10000000000000
		machine.prizeY += 10000000000000
		tokens += calcMinTokens(machine)
	}
	println(tokens)
}

// I had to chatgpt the math for this...
func calcMinTokens(mach machine) int {
	det := ((mach.aX * mach.bY) - (mach.aY * mach.bX))
	aPressesNum := ((mach.bY * mach.prizeX) - (mach.bX * mach.prizeY))
	bPressesNum := ((-mach.aY * mach.prizeX) + (mach.aX * mach.prizeY))
	if aPressesNum%det == 0 && bPressesNum%det == 0 {
		return 3*(aPressesNum/det) + (bPressesNum / det)
	}
	return 0
}
