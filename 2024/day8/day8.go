package main

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]
	input := readInput(filename)
	part1(input)
	part2(input)
}

func readInput(filename string) [][]byte {
	defer utils.TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := [][]byte{}
	for scanner.Scan() {
		input = append(input, []byte(scanner.Text()))
	}
	return input
}

func createKey(i int, j int) string {
	return strconv.Itoa(i) + "," + strconv.Itoa(j)
}

func getAntennaMap(input [][]byte) map[byte][][]int {
	antennas := make(map[byte][][]int)

	for i, row := range input {
		for j, val := range row {
			if val == '.' {
				continue
			}
			vals, ok := antennas[val]
			if !ok {
				vals = [][]int{}
			}
			vals = append(vals, []int{i, j})
			antennas[val] = vals
		}
	}

	return antennas
}

func part1(input [][]byte) {
	defer utils.TimeTrack(time.Now())
	antinodes := make(map[string]struct{})
	antennas := getAntennaMap(input)

	for antenna := range antennas {
		locations := antennas[antenna]
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				loc1 := locations[i]
				loc2 := locations[j]
				anti1 := []int{loc1[0] + (loc1[0] - loc2[0]), loc1[1] + (loc1[1] - loc2[1])}
				if anti1[0] >= 0 && anti1[0] < len(input[0]) && anti1[1] >= 0 && anti1[1] < len(input) {
					antinodes[createKey(anti1[0], anti1[1])] = struct{}{}
				}
				anti2 := []int{loc2[0] + (loc2[0] - loc1[0]), loc2[1] + (loc2[1] - loc1[1])}
				if anti2[0] >= 0 && anti2[0] < len(input[0]) && anti2[1] >= 0 && anti2[1] < len(input) {
					antinodes[createKey(anti2[0], anti2[1])] = struct{}{}
				}
			}
		}
	}
	println(len(antinodes))
}

func part2(input [][]byte) {
	defer utils.TimeTrack(time.Now())
	antinodes := make(map[string]struct{})
	antennas := getAntennaMap(input)

	// Add all antennas to the antinodes map
	for antenna := range antennas {
		locs := antennas[antenna]
		for _, loc := range locs {
			antinodes[createKey(loc[0], loc[1])] = struct{}{}
		}
	}

	for antenna := range antennas {
		locations := antennas[antenna]
		for i := 0; i < len(locations); i++ {
			for j := i + 1; j < len(locations); j++ {
				loc1 := locations[i]
				loc2 := locations[j]
				xDiff := loc1[0] - loc2[0]
				yDiff := loc1[1] - loc2[1]
				anti1 := []int{loc1[0] + xDiff, loc1[1] + yDiff}
				for anti1[0] >= 0 && anti1[0] < len(input[0]) && anti1[1] >= 0 && anti1[1] < len(input) {
					antinodes[createKey(anti1[0], anti1[1])] = struct{}{}
					anti1 = []int{anti1[0] + xDiff, anti1[1] + yDiff}
				}

				xDiff = loc2[0] - loc1[0]
				yDiff = loc2[1] - loc1[1]
				anti2 := []int{loc2[0] + xDiff, loc2[1] + yDiff}
				for anti2[0] >= 0 && anti2[0] < len(input[0]) && anti2[1] >= 0 && anti2[1] < len(input) {
					antinodes[createKey(anti2[0], anti2[1])] = struct{}{}
					anti2 = []int{anti2[0] + xDiff, anti2[1] + yDiff}
				}
			}
		}
	}
	println(len(antinodes))
}
