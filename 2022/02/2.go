// https://adventofcode.com/2022/day/2
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	score1 = map[string]int{
		"A X": 1 + 3,
		"A Y": 2 + 6,
		"A Z": 3 + 0,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 1 + 6,
		"C Y": 2 + 0,
		"C Z": 3 + 3,
	}

	score2 = map[string]int{
		"A X": 3 + 0,
		"A Y": 1 + 3,
		"A Z": 2 + 6,
		"B X": 1 + 0,
		"B Y": 2 + 3,
		"B Z": 3 + 6,
		"C X": 2 + 0,
		"C Y": 3 + 3,
		"C Z": 1 + 6,
	}
)

func main() {
	part1, part2, err := play(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1)
	fmt.Println(part2)
}

func play(r io.Reader) (int, int, error) {
	scanner := bufio.NewScanner(r)
	part1 := 0
	part2 := 0
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		part1 += score1[s]
		part2 += score2[s]
	}
	return part1, part2, scanner.Err()
}
