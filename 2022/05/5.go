// https://adventofcode.com/2022/day/5
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	stacks, moves, err := readStacksAndMoves(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(clone(stacks), moves))
	fmt.Println(part2(clone(stacks), moves))
}

func readStacksAndMoves(r io.Reader) (map[int][]rune, [][3]int, error) {
	scanner := bufio.NewScanner(r)

	stacks, err := readStacks(scanner)
	if err != nil {
		return nil, nil, err
	}

	var moves [][3]int
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		match := re.FindStringSubmatch(s)
		if len(match) != 4 {
			return nil, nil, fmt.Errorf("%q: does not match move line", s)
		}

		var move [3]int
		move[0], _ = strconv.Atoi(match[1])
		move[1], _ = strconv.Atoi(match[2])
		move[2], _ = strconv.Atoi(match[3])

		moves = append(moves, move)
	}
	return stacks, moves, scanner.Err()
}

func readStacks(scanner *bufio.Scanner) (map[int][]rune, error) {
	stacks := map[int][]rune{}
	for scanner.Scan() {
		s := scanner.Text()
		if s == "" {
			break
		}
		s += " " // normalize length of last crate to 4 chars
		for i := 1; len(s) >= 4; i++ {
			crate := strings.TrimSpace(s[:4])
			s = s[4:]
			if len(crate) != 3 {
				continue
			}
			stacks[i] = append(stacks[i], rune(crate[1]))
		}
	}
	return stacks, scanner.Err()
}

func clone(stacks map[int][]rune) map[int][]rune {
	cl := make(map[int][]rune, len(stacks))
	for i, stack := range stacks {
		cl[i] = make([]rune, len(stack))
		copy(cl[i], stack)
	}
	return cl
}

func part1(stacks map[int][]rune, moves [][3]int) string {
	for _, move := range moves {
		n, from, to := move[0], move[1], move[2]
		stacks[to] = append(append([]rune(nil), reverse(stacks[from][:n])...), stacks[to]...)
		stacks[from] = stacks[from][n:]
	}
	return topCrates(stacks)
}

func reverse(stack []rune) []rune {
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
}

func topCrates(stacks map[int][]rune) string {
	var tops []rune
	for i := 1; i <= len(stacks); i++ { // stacks are 1 indexed
		tops = append(tops, stacks[i][0])
	}
	return string(tops)
}

func part2(stacks map[int][]rune, moves [][3]int) string {
	for _, move := range moves {
		n, from, to := move[0], move[1], move[2]
		stacks[to] = append(append([]rune(nil), stacks[from][:n]...), stacks[to]...)
		stacks[from] = stacks[from][n:]
	}
	return topCrates(stacks)
}
