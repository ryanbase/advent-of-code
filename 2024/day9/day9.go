package main

import (
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

func main() {
	filename := utils.GetFileNameFromArgument()

	part1(utils.ReadInputAsIntArray(filename))
	part2(utils.ReadInputAsIntArray(filename))
}

func part1(disk []int) {
	defer utils.TimeTrack(time.Now())

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
	defer utils.TimeTrack(time.Now())

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
