/*****
* Solution was derived from https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/
* I could not figure this out by myself :(
*****/
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
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	cache := make(map[string]int)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		groups := make([]int, 0)
		for _, val := range strings.Split(split[1], ",") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			groups = append(groups, num)
		}
		res := calc(split[0], groups, cache)
		total += res
	}

	println(total)
}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	total := 0
	cache := make(map[string]int)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		groups := make([]int, 0)
		for _, val := range strings.Split(split[1], ",") {
			num, err := strconv.Atoi(val)
			if err != nil {
				panic(err)
			}
			groups = append(groups, num)
		}
		springs := split[0]
		allGroups := make([]int, 0)
		allGroups = append(allGroups, groups...)
		for i := 0; i < 4; i++ {
			springs += "?" + split[0]
			allGroups = append(allGroups, groups...)
		}
		res := calc(springs, allGroups, cache)
		total += res
	}

	println(total)
}

func calc(springs string, groups []int, cache map[string]int) int {
	cacheKey := createCacheKey(springs, groups)
	val, exists := cache[cacheKey]
	if exists {
		return val
	}

	if len(groups) == 0 {
		if !strings.Contains(springs, "#") {
			cache[cacheKey] = 1
			return 1
		} else {
			cache[cacheKey] = 0
			return 0
		}
	}

	if len(springs) == 0 {
		cache[cacheKey] = 0
		return 0
	}

	nextChar := springs[0]
	nextGroup := groups[0]

	out := 0
	if nextChar == '#' {
		out = pound(springs, groups, nextGroup, cache)
	} else if nextChar == '.' {
		out = dot(springs, groups, cache)
	} else if nextChar == '?' {
		out = dot(springs, groups, cache) + pound(springs, groups, nextGroup, cache)
	} else {
		panic("An invalid character was read")
	}
	cache[cacheKey] = out
	return out
}

func pound(springs string, groups []int, nextGroup int, cache map[string]int) int {
	if len(springs) < nextGroup {
		return 0
	}
	currentGroup := springs[:nextGroup]
	currentGroup = strings.ReplaceAll(currentGroup, "?", "#")

	if currentGroup != fillString("#", nextGroup) {
		return 0
	}

	if len(springs) == nextGroup {
		if len(groups) == 1 {
			return 1
		} else {
			return 0
		}
	}

	if springs[nextGroup] == '?' || springs[nextGroup] == '.' {
		return calc(springs[nextGroup+1:], groups[1:], cache)
	}

	return 0
}

func dot(springs string, groups []int, cache map[string]int) int {
	return calc(springs[1:], groups, cache)
}

func fillString(char string, length int) string {
	str := ""
	for i := 0; i < length; i++ {
		str += char
	}
	return str
}

func createCacheKey(springs string, groups []int) string {
	key := springs
	for _, group := range groups {
		// This might cause the wrong answer to be given. Not adding the groups actually gave the right answer before
		key += strconv.Itoa(group)
	}
	return key
}

type Input struct {
	springs string
	groups  []int
}
