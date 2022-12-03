// https://adventofcode.com/2022/day/3
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

func main() {
	sacks, err := readSacks(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(sacks))
	fmt.Println(part2(sacks))
}

func readSacks(r io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(r)
	var sacks []string
	for scanner.Scan() {
		sacks = append(sacks, strings.TrimSpace(scanner.Text()))
	}
	return sacks, scanner.Err()
}

func part1(sacks []string) int {
	sum := 0
	for _, sack := range sacks {
		second := sack[len(sack)/2:]
		for _, item := range sack {
			if strings.ContainsRune(second, item) {
				sum += priority(item)
				break
			}
		}
	}
	return sum
}

func priority(item rune) int {
	if unicode.IsLower(item) {
		return int(item-'a') + 1
	}
	return int(item-'A') + 27
}

func part2(sacks []string) int {
	sum := 0
	for i := 0; i < len(sacks); i += 3 {
		for _, item := range sacks[i] {
			if strings.ContainsRune(sacks[i+1], item) && strings.ContainsRune(sacks[i+2], item) {
				sum += priority(item)
				break
			}
		}
	}
	return sum
}
