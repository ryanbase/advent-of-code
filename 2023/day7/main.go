package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
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

	handTypes := make([][]Hand, 7)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{split[0], bid}
		handType := getHandType(hand.cards, false)
		handTypes[handType] = append(handTypes[handType], hand)
	}

	// Sort hand type arrays, flatten, calc scores
	rankedHands := make([]Hand, 0)
	for _, handType := range handTypes {
		slices.SortFunc(handType, sortHands)
		rankedHands = append(rankedHands, handType...)
	}

	score := 0
	for i, hand := range rankedHands {
		score += hand.bid * (i + 1)
	}

	println(score)

}

func partTwo() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	handTypes := make([][]Hand, 7)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		bid, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		hand := Hand{split[0], bid}
		handType := getHandType(hand.cards, true)
		handTypes[handType] = append(handTypes[handType], hand)
	}

	// Sort hand type arrays, flatten, calc scores
	rankedHands := make([]Hand, 0)
	for _, handType := range handTypes {
		slices.SortFunc(handType, sortHandsWithJokers)
		rankedHands = append(rankedHands, handType...)
	}

	score := 0
	for i, hand := range rankedHands {
		score += hand.bid * (i + 1)
	}

	println(score)
}

func sortHands(hand1 Hand, hand2 Hand) int {
	for i := 0; i < len(hand1.cards); i++ {
		compare := cmp.Compare(getCardValue(hand1.cards[i]), getCardValue(hand2.cards[i]))
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func sortHandsWithJokers(hand1 Hand, hand2 Hand) int {
	for i := 0; i < len(hand1.cards); i++ {
		compare := cmp.Compare(getCardValueWithJokers(hand1.cards[i]), getCardValueWithJokers(hand2.cards[i]))
		if compare != 0 {
			return compare
		}
	}
	return 0
}

func getCardValue(card byte) int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'J':
		return 9
	case 'T':
		return 8
	case '9':
		return 7
	case '8':
		return 6
	case '7':
		return 5
	case '6':
		return 4
	case '5':
		return 3
	case '4':
		return 2
	case '3':
		return 1
	case '2':
		return 0
	default:
		return -1
	}
}

func getCardValueWithJokers(card byte) int {
	switch card {
	case 'A':
		return 12
	case 'K':
		return 11
	case 'Q':
		return 10
	case 'T':
		return 9
	case '9':
		return 8
	case '8':
		return 7
	case '7':
		return 6
	case '6':
		return 5
	case '5':
		return 4
	case '4':
		return 3
	case '3':
		return 2
	case '2':
		return 1
	case 'J':
		return 0
	default:
		return -1
	}
}

func getCardCounts(cards string, jokers bool) []int {
	counts := make([]int, 13)
	for i := range cards {
		if jokers {
			counts[getCardValueWithJokers(cards[i])]++
		} else {
			counts[getCardValue(cards[i])]++
		}
	}
	return counts
}

/**
* 0 = High Card
* 1 = One Pair
* 2 = Two Pair
* 3 = Three of a kind
* 4 = Full House
* 5 = Four of a kind
* 6 = Five of a kind
**/
func getHandType(cards string, jokers bool) int {
	counts := getCardCounts(cards, jokers)

	if jokers {
		jokerCount := counts[0]
		if jokerCount > 0 {
			highestCount := 0
			highestCountIndex := 0
			for i, count := range counts {
				if i == 0 {
					continue
				}
				if count >= highestCount {
					highestCount = count
					highestCountIndex = i
				}
			}
			counts[highestCountIndex] += jokerCount
			counts[0] -= jokerCount
		}
	}

	possibleFullHouse := 0
	for _, count := range counts {
		if count == 5 {
			return 6
		}
		if count == 4 {
			return 5
		}
		if count == 3 {
			if possibleFullHouse == 2 {
				return 4
			}
			possibleFullHouse = 3
		}
		if count == 2 {
			if possibleFullHouse == 3 {
				return 4
			}
			if possibleFullHouse == 2 {
				return 2
			}
			possibleFullHouse = 2
		}
	}
	if possibleFullHouse == 3 {
		return 3
	}
	if possibleFullHouse == 2 {
		return 1
	}
	return 0
}

type Hand struct {
	cards string
	bid   int
}
