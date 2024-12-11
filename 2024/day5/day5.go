package main

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	filename := utils.GetFileNameFromArgument()

	orders, updates := readInput(filename)
	part1(orders, updates)
	part2(orders, updates)
}

func readInput(filename string) (map[string][]string, [][]string) {
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	section := 1
	orders := make(map[string][]string)
	updates := [][]string{}
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			section = 2
			continue
		}

		if section == 1 {
			order := strings.Split(line, "|")
			existing, ok := orders[order[0]]
			if !ok {
				existing = []string{}
			}
			existing = append(existing, order[1])
			orders[order[0]] = existing
		}

		if section == 2 {
			updates = append(updates, strings.Split(line, ","))
		}
	}

	return orders, updates
}

func part1(orders map[string][]string, updates [][]string) {
	result := 0
	for _, update := range updates {
		if isOrdered(orders, update) == -1 {
			middle, _ := strconv.Atoi(update[len(update)/2])
			result += middle
		}
	}
	println(result)
}

func part2(orders map[string][]string, updates [][]string) {
	unordered := [][]string{}
	for _, update := range updates {
		if isOrdered(orders, update) > -1 {
			unordered = append(unordered, update)
		}
	}

	result := 0
	for _, update := range unordered {
		badIndex := isOrdered(orders, update)
		for badIndex > -1 {
			page := update[badIndex]
			update = append(update[:badIndex], update[badIndex+1:]...)
			after := orders[page]
			insertIndex := -1
			for i, p := range update {
				if slices.Contains(after, p) {
					insertIndex = i
					break
				}
			}
			update = slices.Insert(update, insertIndex, page)
			badIndex = isOrdered(orders, update)
		}
		middle, _ := strconv.Atoi(update[len(update)/2])
		result += middle
	}

	println(result)
}

func findIndex() {

}

// Returns the index of the page that is in the wrong place, or -1 if it's ordered correctly
func isOrdered(orders map[string][]string, update []string) int {
	for i, page := range update {
		after := orders[page]
		// for each page, check that each page before it is not in the after list
		for j := 0; j < i; j++ {
			for _, a := range after {
				if update[j] == a {
					return i
				}
			}
		}
	}
	return -1
}
