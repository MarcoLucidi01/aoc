// https://adventofcode.com/2020/day/2
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"unicode/utf8"
)

func main() {
	nvalidPart1 := 0
	nvalidPart2 := 0
	for {
		var a, b int
		var c rune
		var pwd string
		if _, err := fmt.Scanf("%d-%d %c: %s\n", &a, &b, &c, &pwd); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			log.Fatal(err)
		}
		if isValidPart1(a, b, c, pwd) {
			nvalidPart1++
		}
		if isValidPart2(a, b, c, pwd) {
			nvalidPart2++
		}
	}
	fmt.Printf("part 1: %d\n", nvalidPart1)
	fmt.Printf("part 2: %d\n", nvalidPart2)
}

func isValidPart1(min, max int, c rune, pwd string) bool {
	n := 0
	for _, r := range pwd {
		if c == r {
			n++
		}
	}
	return n >= min && n <= max
}

func isValidPart2(pos1, pos2 int, c rune, pwd string) bool {
	r, _ := utf8.DecodeRuneInString(pwd[pos1-1:])
	pos1Equal := r == c
	r, _ = utf8.DecodeRuneInString(pwd[pos2-1:])
	pos2Equal := r == c
	return (pos1Equal && !pos2Equal) || (!pos1Equal && pos2Equal)
}
