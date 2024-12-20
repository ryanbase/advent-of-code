package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type state struct {
	a     int
	b     int
	c     int
	input []int
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	part1(filename)
	part2(filename)
}

func readInput(filename string) state {
	defer utils.TimeTrack(time.Now())

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	state := state{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "Register A") {
			state.a = toInt(strings.Split(line, ": ")[1])
			continue
		}
		if strings.Contains(line, "Register B") {
			state.b = toInt(strings.Split(line, ": ")[1])
			continue
		}
		if strings.Contains(line, "Register C") {
			state.c = toInt(strings.Split(line, ": ")[1])
			continue
		}
		if strings.Contains(line, "Program") {
			vals := strings.Split(strings.Split(line, ": ")[1], ",")
			for _, val := range vals {
				state.input = append(state.input, toInt(val))
			}
			continue
		}
	}

	return state
}

func part1(filename string) {
	defer utils.TimeTrack(time.Now())

	state := readInput(filename)
	res := runProgram(state)
	fmt.Println(res)
}

func part2(filename string) {
	defer utils.TimeTrack(time.Now())

	origState := readInput(filename)

	a := 1
	for true {
		if a == origState.a {
			a++
			continue
		}
		state := state{a, origState.b, origState.c, origState.input}
		out := runProgram(state)
		if out == intsToStr(origState.input) {
			break
		}
		a++
	}

	println(a)
}

func runProgram(state state) string {
	out := []int{}

	for i := 0; i < len(state.input); i += 2 {
		opcode := state.input[i]
		operand := state.input[i+1]

		switch opcode {
		case 0:
			state.a = state.a / int(math.Pow(2, float64(getComboOperand(state, operand))))
			break
		case 1:
			state.b = state.b ^ operand
			break
		case 2:
			state.b = getComboOperand(state, operand) % 8
			break
		case 3:
			if state.a == 0 {
				break
			}
			i = operand - 2
			break
		case 4:
			state.b = state.b ^ state.c
			break
		case 5:
			out = append(out, getComboOperand(state, operand)%8)
			break
		case 6:
			state.b = state.a / int(math.Pow(2, float64(getComboOperand(state, operand))))
			break
		case 7:
			state.c = state.a / int(math.Pow(2, float64(getComboOperand(state, operand))))
			break
		default:
			panic("We should not be here")

		}
	}

	return intsToStr(out)
}

func getComboOperand(state state, operand int) int {
	switch operand {
	case 0:
		return operand
	case 1:
		return operand
	case 2:
		return operand
	case 3:
		return operand
	case 4:
		return state.a
	case 5:
		return state.b
	case 6:
		return state.c
	case 7:
		panic("Operand 7 received. Invalid program.")
	default:
		panic("We should not be here")
	}
}

func intsToStr(ints []int) string {
	strs := []string{}
	for _, num := range ints {
		strs = append(strs, strconv.Itoa(num))
	}

	return strings.Join(strs, ",")
}

func toInt(val string) int {
	num, _ := strconv.Atoi(val)
	return num
}
