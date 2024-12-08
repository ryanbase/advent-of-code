package main

import (
	"bufio"
	"os"
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

	platform := make([]string, 0)

	for scanner.Scan() {
		platform = append(platform, scanner.Text())
	}

	platform = rollNorth(platform)
	load := calculateLoad(platform)

	println("Total:", load)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	platform := make([]string, 0)

	for scanner.Scan() {
		platform = append(platform, scanner.Text())
	}

	cycles := 1000000000
	// cycles := 1000 // 1000 works without finding the loop size

	states := make(map[string]int)

	loopSize := 0
	cyclesNeeded := cycles
	for i := 0; i < loopSize+cyclesNeeded; i++ {
		platform = cycle(platform)
		if loopSize == 0 {
			str := arrayAsString(platform)
			_, ok := states[str]
			if !ok {
				states[str] = 1
			} else {
				loopSize = i + 1
				cyclesNeeded = loopSize - (cycles % loopSize)
			}
		}
	}

	load := calculateLoad(platform)

	println("Total:", load)
}

func cycle(platform []string) []string {
	platform = rollNorth(platform)
	platform = rollWest(platform)
	platform = rollSouth(platform)
	platform = rollEast(platform)
	return platform
}

func rollNorth(platform []string) []string {
	for col := 0; col < len(platform[0]); col++ {
		location := 0
		for row := 0; row < len(platform); row++ {
			if platform[row][col] == '#' {
				location = row + 1
			} else if platform[row][col] == 'O' {
				platform[row] = replaceAtIndex(platform[row], '.', col)
				platform[location] = replaceAtIndex(platform[location], 'O', col)
				location++
			}
		}
	}

	return platform
}

func rollSouth(platform []string) []string {
	for col := 0; col < len(platform[0]); col++ {
		location := len(platform[0]) - 1
		for row := len(platform) - 1; row >= 0; row-- {
			if platform[row][col] == '#' {
				location = row - 1
			} else if platform[row][col] == 'O' {
				platform[row] = replaceAtIndex(platform[row], '.', col)
				platform[location] = replaceAtIndex(platform[location], 'O', col)
				location--
			}
		}
	}

	return platform
}

func rollWest(platform []string) []string {
	for row := 0; row < len(platform); row++ {
		location := 0
		for col := 0; col < len(platform[row]); col++ {
			if platform[row][col] == '#' {
				location = col + 1
			} else if platform[row][col] == 'O' {
				platform[row] = replaceAtIndex(platform[row], '.', col)
				platform[row] = replaceAtIndex(platform[row], 'O', location)
				location++
			}
		}
	}

	return platform
}

func rollEast(platform []string) []string {
	for row := 0; row < len(platform); row++ {
		location := len(platform[row]) - 1
		for col := len(platform[row]) - 1; col >= 0; col-- {
			if platform[row][col] == '#' {
				location = col - 1
			} else if platform[row][col] == 'O' {
				platform[row] = replaceAtIndex(platform[row], '.', col)
				platform[row] = replaceAtIndex(platform[row], 'O', location)
				location--
			}
		}
	}

	return platform
}

func calculateLoad(platform []string) int {
	load := 0
	for col := 0; col < len(platform[0]); col++ {
		for row := 0; row < len(platform); row++ {
			if platform[row][col] == 'O' {
				load += len(platform) - row
			}
		}
	}
	return load
}

func arrayAsString(arr []string) string {
	str := ""
	for _, s := range arr {
		str += s
	}
	return str
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
