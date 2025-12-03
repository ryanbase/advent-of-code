package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"github.com/ryanbase/advent-of-code/2024/utils"
)

type trie struct {
	root *node
}

type node struct {
	isEnd    bool
	children []*node
}

func (t *trie) insert(word string) {
	currNode := t.root
	for _, char := range word {
		next := currNode.children[char-'a']
		if next == nil {
			next = &node{false, make([]*node, 26)}
			currNode.children[char-'a'] = next
		}
		currNode = currNode.children[char-'a']
	}
	currNode.isEnd = true
}

func (t *trie) search(word string) bool {
	currNode := t.root
	for _, char := range word {
		next := currNode.children[char-'a']
		if next == nil {
			return false
		}
		currNode = next
	}
	return currNode.isEnd
}

func newTrie() *trie {
	rootNode := &node{false, make([]*node, 26)}
	return &trie{rootNode}
}

func main() {
	defer utils.TimeTrack(time.Now())
	filename := utils.GetFileNameFromArgument()
	trie, lines := readInput(filename)
	part1(trie, lines)
}

func readInput(filename string) (*trie, []string) {
	defer utils.TimeTrack(time.Now())

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	line := scanner.Text()
	vals := strings.Split(line, ", ")

	trie := newTrie()

	// Insert into trie
	for _, val := range vals {
		trie.insert(val)
	}

	// Skip empty line
	scanner.Scan()
	scanner.Text()

	lines := []string{}
	for scanner.Scan() {
		line = scanner.Text()
		lines = append(lines, line)
	}

	return trie, lines
}

func part1(trie *trie, lines []string) {
	count := 0
	for _, line := range lines {
		l := 0
		r := 1
		for r <= len(line) {
			println("checking", line[l:r])
			if trie.search(line[l:r]) {
				if r == len(line) {
					count++
					println(line)
					break
				}
				l = r
			}
			r++
		}
	}
	println(count)
}
