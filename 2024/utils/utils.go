package utils

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

func GetFileNameFromArgument() string {
	if len(os.Args) != 2 {
		panic("No file name provided")
	}
	return os.Args[1]
}

func ReadInputAsString(filename string) string {
	defer TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := ""
	for scanner.Scan() {
		line := scanner.Text()
		input += line
	}
	return input
}

func ReadInputAsByteMatrix(filename string) [][]byte {
	defer TimeTrack(time.Now())
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

func ReadInputAsIntMatrix(filename string) [][]int {
	defer TimeTrack(time.Now())
	f, err := os.Open(filename)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	input := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		arr := strings.Split(line, "")
		row := []int{}
		for _, val := range arr {
			num, _ := strconv.Atoi(val)
			row = append(row, num)
		}
		input = append(input, row)

	}
	return input
}

func ReadInputAsIntArray(filename string) []int {
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

func TimeTrack(start time.Time) {
	elapsed := time.Since(start)
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	runtimeFunc := regexp.MustCompile(`^.*\.(.*)$`)
	name := runtimeFunc.ReplaceAllString(funcObj.Name(), "$1")
	fmt.Printf("%s completed in %s\n", name, elapsed)
}
