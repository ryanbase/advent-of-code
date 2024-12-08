package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

var totalLowPulses int64
var totalHighPulses int64

func partOne() {
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	totalLowPulses = 0
	totalHighPulses = 0
	modules := make(map[string]module)
	var broadcaster module

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "broadcaster") {
			mods := strings.Split(line, " -> ")
			mods = strings.Split(mods[1], ", ")
			broadcaster = module{Broadcaster, mods, make(map[string]int), false}
		} else {
			mods := strings.Split(line, " -> ")
			name := mods[0][1:]
			subs := strings.Split(mods[1], ", ")
			if strings.Contains(line, "%") {
				modules[name] = module{FlipFlop, subs, make(map[string]int), false}
			} else if strings.Contains(line, "&") {
				subMap := make(map[string]int)
				for _, mod := range mods {
					subMap[mod] = LowPulse
				}
				modules[name] = module{Conjuction, subs, subMap, false}
			} else {
				panic("Unexpected mod type")
			}
		}
	}

	presses := 1

	for i := 0; i < presses; i++ {
		totalLowPulses++
		for _, mod := range broadcaster.subscribers {
			totalLowPulses++
			handlePulse(modules, mod, "broadcaster", LowPulse)
		}
	}

	println(totalLowPulses, totalHighPulses)
}

func partTwo() {
	// f, err := os.Open("test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer f.Close()

	// scanner := bufio.NewScanner(f)

	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }
}

func handlePulse(modules map[string]module, modName string, fromMod string, pulseType int) {
	module := modules[modName]
	pulseToSend := LowPulse
	send := false
	println(fromMod, pulseType, modName)
	if module.modType == FlipFlop {
		if pulseType == LowPulse {
			module.on = !module.on
			send = true
			if module.on {
				pulseToSend = HighPulse
			}
		}
	} else if module.modType == Conjuction {
		module.pulses[fromMod] = pulseType
		send = true
		for _, val := range module.pulses {
			if val == LowPulse {
				pulseToSend = HighPulse
				break
			}
		}
	}

	if send {
		for _, mod := range module.subscribers {
			if pulseToSend == LowPulse {
				totalLowPulses++
			} else {
				totalHighPulses++
			}
			handlePulse(modules, mod, modName, pulseToSend)
		}
	}
}

type module struct {
	modType     int
	subscribers []string
	pulses      map[string]int
	on          bool
}

const (
	LowPulse  = iota
	HighPulse = iota
)

const (
	Broadcaster = iota
	FlipFlop    = iota
	Conjuction  = iota
)
