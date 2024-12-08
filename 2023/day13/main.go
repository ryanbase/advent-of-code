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

	pattern := make([]string, 0)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pattern = append(pattern, line)
		} else {
			mirror := findHorizontalMirror(pattern)
			if mirror > 0 {
				total += (mirror * 100)
			} else {
				mirror = findVerticalMirror(pattern)
				if mirror > 0 {
					total += mirror
				}
			}
			pattern = make([]string, 0)
		}
	}

	mirror := findHorizontalMirror(pattern)
	if mirror > 0 {
		total += (mirror * 100)
	} else {
		mirror = findVerticalMirror(pattern)
		if mirror > 0 {
			total += mirror
		}
	}

	println("Total:", total)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	pattern := make([]string, 0)
	total := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			pattern = append(pattern, line)
		} else {
			mirror := findHorizontalMirror2(pattern)
			if mirror > 0 {
				total += (mirror * 100)
			} else {
				mirror = findVerticalMirror2(pattern)
				if mirror > 0 {
					total += mirror
				}
			}
			pattern = make([]string, 0)
		}
	}

	mirror := findHorizontalMirror2(pattern)
	if mirror > 0 {
		total += (mirror * 100)
	} else {
		mirror = findVerticalMirror2(pattern)
		if mirror > 0 {
			total += mirror
		}
	}

	println("Total:", total)
}

func findHorizontalMirror(pattern []string) int {
	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == pattern[i+1] {
			isMirror := true
			for j := 1; i+j+1 < len(pattern) && i-j >= 0; j++ {
				if pattern[i-j] != pattern[i+j+1] {
					isMirror = false
					break
				}
			}
			if isMirror {
				return i + 1
			}
		}
	}
	return 0
}

func findVerticalMirror(pattern []string) int {
	col1 := ""
	col2 := ""
	for col := 0; col < len(pattern[0])-1; col++ {
		for row := 0; row < len(pattern); row++ {
			col1 = col1 + string(pattern[row][col])
			col2 = col2 + string(pattern[row][col+1])
		}
		if col1 == col2 {
			isMirror := true
			for j := 1; col-j >= 0 && col+j+1 < len(pattern[0]); j++ {
				col1 = ""
				col2 = ""
				for row := 0; row < len(pattern); row++ {
					col1 = col1 + string(pattern[row][col-j])
					col2 = col2 + string(pattern[row][col+j+1])
				}
				if col1 != col2 {
					isMirror = false
					break
				}
			}
			if isMirror {
				return col + 1
			}
		}
		col1 = ""
		col2 = ""
	}
	return 0
}

func findHorizontalMirror2(pattern []string) int {
	for i := 0; i < len(pattern)-1; i++ {
		if pattern[i] == pattern[i+1] || offByOne(pattern[i], pattern[i+1]) {
			isMirror := true
			offByOneFound := offByOne(pattern[i], pattern[i+1])
			for j := 1; i+j+1 < len(pattern) && i-j >= 0; j++ {
				if pattern[i-j] != pattern[i+j+1] {
					if offByOne(pattern[i-j], pattern[i+j+1]) {
						if offByOneFound {
							isMirror = false
							break
						}
						offByOneFound = true
					} else {
						isMirror = false
						break
					}
				}
			}
			if isMirror && offByOneFound {
				return i + 1
			}
		}
	}
	return 0
}

func findVerticalMirror2(pattern []string) int {
	col1 := ""
	col2 := ""
	for col := 0; col < len(pattern[0])-1; col++ {
		for row := 0; row < len(pattern); row++ {
			col1 = col1 + string(pattern[row][col])
			col2 = col2 + string(pattern[row][col+1])
		}
		if col1 == col2 || offByOne(col1, col2) {
			isMirror := true
			offByOneFound := offByOne(col1, col2)
			for j := 1; col-j >= 0 && col+j+1 < len(pattern[0]); j++ {
				col1 = ""
				col2 = ""
				for row := 0; row < len(pattern); row++ {
					col1 = col1 + string(pattern[row][col-j])
					col2 = col2 + string(pattern[row][col+j+1])
				}
				if col1 != col2 {
					if offByOne(col1, col2) {
						if offByOneFound {
							isMirror = false
							break
						}
						offByOneFound = true
					} else {
						isMirror = false
						break
					}
				}
			}
			if isMirror && offByOneFound {
				return col + 1
			}
		}
		col1 = ""
		col2 = ""
	}
	return 0
}

func offByOne(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	if str1 == str2 {
		return false
	}
	foundDiff := false
	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[i] {
			if foundDiff {
				return false
			}
			foundDiff = true
		}
	}
	return true
}
