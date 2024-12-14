package main

import (
	"strconv"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type cacheval struct {
	Blinks   int
	StoneNum string
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	inputStr := utils.ReadInputAsString(filename)
	input := strings.Split(inputStr, " ")
	cache := make(map[cacheval]int)
	part1(input, cache)
	part2(input, cache)
}

func part1(input []string, cache map[cacheval]int) {
	defer utils.TimeTrack(time.Now())
	res := calculate(input, 25, cache)
	println(res)
}

func part2(input []string, cache map[cacheval]int) {
	defer utils.TimeTrack(time.Now())
	res := calculate(input, 75, cache)
	println(res)
}

func calculate(input []string, blinks int, cache map[cacheval]int) int {
	result := 0
	for _, val := range input {
		result += applyRules(val, blinks, cache)
	}
	return result
}

func applyRules(val string, blinks int, cache map[cacheval]int) int {
	if blinks == 0 {
		return 1
	}
	res := cache[cacheval{blinks, val}]
	if res > 0 {
		return res
	}
	if val == "0" {
		res += applyRules("1", blinks-1, cache)
	} else if len(val)%2 == 0 {
		newNum1, _ := strconv.Atoi(val[:len(val)/2])
		newStr1 := strconv.Itoa(newNum1)
		newNum2, _ := strconv.Atoi(val[len(val)/2:])
		newStr2 := strconv.Itoa(newNum2)
		res += applyRules(newStr1, blinks-1, cache) + applyRules(newStr2, blinks-1, cache)
	} else {
		num, _ := strconv.Atoi(val)
		res += applyRules(strconv.Itoa(num*2024), blinks-1, cache)

	}
	cache[cacheval{blinks, val}] = res
	return res
}
