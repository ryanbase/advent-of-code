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

	workflows := make(map[string]workflow)
	parts := []part{}
	parseWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseWorkflows = false
		} else if parseWorkflows {
			workflow := parseWorkflow(line)
			workflows[workflow.name] = workflow
		} else {
			line = strings.ReplaceAll(line, "{", "")
			line = strings.ReplaceAll(line, "}", "")
			lineSplit := strings.Split(line, ",")
			x := 0
			m := 0
			a := 0
			s := 0
			for _, c := range lineSplit {
				opVal := strings.Split(c, "=")
				val, err := strconv.Atoi(opVal[1])
				if err != nil {
					panic(err)
				}
				if opVal[0] == "x" {
					x = val
				} else if opVal[0] == "m" {
					m = val
				} else if opVal[0] == "a" {
					a = val
				} else if opVal[0] == "s" {
					s = val
				}
			}
			parts = append(parts, part{x, m, a, s})
		}
	}

	accepted := []part{}
	rejected := []part{}
	for _, part := range parts {
		workflow := workflows["in"]
		for true {
			result := ""
			for _, rule := range workflow.rules {
				val := 0
				if rule.category == 'x' {
					val = part.x
				} else if rule.category == 'm' {
					val = part.m
				} else if rule.category == 'a' {
					val = part.a
				} else if rule.category == 's' {
					val = part.s
				}
				if (rule.operator == '>' && val > rule.value) || (rule.operator == '<' && val < rule.value) {
					result = rule.result
					break
				}
			}
			if result != "" {
				if result == "A" {
					accepted = append(accepted, part)
					break
				} else if result == "R" {
					rejected = append(rejected, part)
					break
				} else {
					workflow = workflows[result]
				}
			} else {
				if workflow.fallback == "A" {
					accepted = append(accepted, part)
					break
				} else if workflow.fallback == "R" {
					rejected = append(rejected, part)
					break
				} else {
					workflow = workflows[workflow.fallback]
				}
			}
		}
	}
	total := 0
	for _, acc := range accepted {
		sum := acc.x + acc.m + acc.a + acc.s
		total += sum
	}

	println(total)
}

func partTwo() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	workflows := make(map[string]workflow)
	parseWorkflows := true
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseWorkflows = false
		} else if parseWorkflows {
			workflow := parseWorkflow(line)
			workflows[workflow.name] = workflow
		}
	}

}

func parseWorkflow(line string) workflow {
	index := strings.IndexRune(line, '{')
	name := line[:index]
	line = line[index+1:]
	line = strings.ReplaceAll(line, "}", "")
	split := strings.Split(line, ",")
	fallback := split[len(split)-1]
	split = split[:len(split)-1]
	rules := []rule{}
	for _, r := range split {
		category := r[0]
		operator := r[1]
		ruleSplit := strings.Split(r[2:], ":")
		result := ruleSplit[1]
		value, err := strconv.Atoi(ruleSplit[0])
		if err != nil {
			panic(err)
		}
		rules = append(rules, rule{category, operator, value, result})
	}
	return workflow{name, rules, fallback}
}

type part struct {
	x, m, a, s int
}

type workflow struct {
	name     string
	rules    []rule
	fallback string
}

type rule struct {
	category byte
	operator byte
	value    int
	result   string
}
