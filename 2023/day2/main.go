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

	for scanner.Scan() {
		line := scanner.Text()

		possible := true
		game := strings.Split(line, ": ")
		gameId := strings.Split(game[0], " ")[1]
		draws := strings.Split(game[1], "; ")

		for _, draw := range draws {
			cubes := strings.Split(draw, ", ")
			red := 0
			blue := 0
			green := 0
			for _, color := range cubes {
				number := strings.Split(color, " ")
				num, err := strconv.Atoi(number[0])
				if err != nil {
					panic(err)
				}
				if number[1] == "red" {
					red += num
				} else if number[1] == "blue" {
					blue += num
				} else if number[1] == "green" {
					green += num
				}

				if red > 12 || green > 13 || blue > 14 {
					possible = false
				}
			}
		}

		if possible {
			id, err := strconv.Atoi(gameId)
			if err != nil {
				panic(err)
			}
			total += id
		}

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

	for scanner.Scan() {
		line := scanner.Text()

		game := strings.Split(line, ": ")
		draws := strings.Split(game[1], "; ")

		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, draw := range draws {
			cubes := strings.Split(draw, ", ")
			for _, color := range cubes {
				number := strings.Split(color, " ")
				num, err := strconv.Atoi(number[0])
				if err != nil {
					panic(err)
				}
				if number[1] == "red" && num > minRed {
					minRed = num
				} else if number[1] == "blue" && num > minBlue {
					minBlue = num
				} else if number[1] == "green" && num > minGreen {
					minGreen = num
				}
			}
		}

		power := minRed * minGreen * minBlue

		total += power

	}

	println(total)
}
