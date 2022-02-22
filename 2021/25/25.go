// https://adventofcode.com/2021/day/25
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	cucumap, err := readCucumap()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(cucumap))
	fmt.Printf("part 2: %d\n", partTwo(cucumap))
}

func readCucumap() ([][]rune, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var cucumap [][]rune
	for scanner.Scan() {
		if line := strings.TrimSpace(scanner.Text()); line != "" {
			cucumap = append(cucumap, []rune(line))
		}
	}
	return cucumap, scanner.Err()
}

func partOne(cucumap [][]rune) int {
	steps := 0
	for moved := true; moved; steps++ {
		moved = false
		for _, row := range cucumap {
			for j, c := range row {
				if c == '>' && row[(j+1)%len(row)] == '.' {
					row[j] = ':'
					row[(j+1)%len(row)] = 'x'
					moved = true
				}
			}
		}
		for _, row := range cucumap {
			for j, c := range row {
				if c == ':' {
					row[j] = '.'
				} else if c == 'x' {
					row[j] = '>'
				}
			}
		}
		for i, row := range cucumap {
			for j, c := range row {
				if c == 'v' && cucumap[(i+1)%len(cucumap)][j] == '.' {
					row[j] = ':'
					cucumap[(i+1)%len(cucumap)][j] = 'x'
					moved = true
				}
			}
		}
		for _, row := range cucumap {
			for j, c := range row {
				if c == ':' {
					row[j] = '.'
				} else if c == 'x' {
					row[j] = 'v'
				}
			}
		}
	}
	return steps
}

func partTwo(cucumap [][]rune) int {
	// TODO I don't have enough stars to unlock part two.
	return 0
}
