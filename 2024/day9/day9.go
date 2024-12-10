package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	filename := os.Args[1]

	part1(readInput(filename))
	part2(readInput(filename))
}

func readInput(filename string) []int {
	defer TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := ""
	for scanner.Scan() {
		input += scanner.Text()
	}
	inputArray := strings.Split(input, "")
	disk := []int{}
	for i, val := range inputArray {
		num, _ := strconv.Atoi(val)
		for range num {
			if i%2 == 0 {
				disk = append(disk, i/2)
			} else {
				disk = append(disk, -1)
			}
		}
	}
	return disk
}

func part1(disk []int) {
	defer TimeTrack(time.Now())

	start := 0
	end := len(disk) - 1
	for start < end {
		if disk[start] != -1 {
			start++
			continue
		}
		if disk[end] == -1 {
			end--
			continue
		}
		disk[start] = disk[end]
		disk[end] = -1
		start++
		end--
	}

	println(calcChecksum(disk))
}

func part2(disk []int) {
	defer TimeTrack(time.Now())

	end := len(disk) - 1
	start := end

	for end >= 0 {
		val := disk[end]
		for val == -1 {
			end--
			val = disk[end]
		}

		start = end
		for start > 0 {
			start--
			if disk[start] != disk[end] {
				start++
				break
			}
		}

		l := 0
		r := l + (end - start)
		for r < start {
			if disk[l] != -1 || disk[r] != -1 {
				l++
				r++
				continue
			}
			isEmpty := true
			for _, val := range disk[l : r+1] {
				if val != -1 {
					isEmpty = false
				}
			}
			if !isEmpty {
				l++
				r++
				continue
			}

			for i := range disk[l : r+1] {
				disk[l+i] = disk[start+i]
				disk[start+i] = -1
			}

			break
		}

		end = start - 1

	}
	println(calcChecksum(disk))
}

func calcChecksum(disk []int) int {
	checksum := 0
	for i, num := range disk {
		if num == -1 {
			continue
		}
		checksum += num * i
	}
	return checksum
}

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	fmt.Printf("%s completed in %s\n", name, elapsed)
}
