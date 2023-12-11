// https://adventofcode.com/2023/day/11
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

type pair struct {
	a, b pos
}

func main() {
	pairs, ex, ey, err := readUniverse(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", solve(pairs, ex, ey, 2))
	fmt.Println("part2:", solve(pairs, ex, ey, 1000000))
}

func readUniverse(r io.Reader) ([]pair, []int, []int, error) {
	scanner := bufio.NewScanner(r)

	var universe [][]int
	var galaxies []pos
	for x := 0; scanner.Scan(); x++ {
		var row []int
		for y, c := range strings.TrimSpace(scanner.Text()) {
			switch c {
			case '.':
				row = append(row, 0)
			case '#':
				galaxies = append(galaxies, pos{x, y})
				row = append(row, 1)
			}
		}
		universe = append(universe, row)
	}

	pairsSet := make(map[pair]bool)
	for _, a := range galaxies {
		for _, b := range galaxies {
			if a == b || pairsSet[pair{a, b}] || pairsSet[pair{b, a}] {
				continue
			}
			pairsSet[pair{a, b}] = true
		}
	}
	var pairs []pair
	for p := range pairsSet {
		pairs = append(pairs, p)
	}

	n := 0
	var ex []int
loop:
	for _, row := range universe {
		ex = append(ex, n)
		for _, c := range row {
			if c != 0 {
				continue loop
			}
		}
		n++
	}

	n = 0
	var ey []int
loop2:
	for i := 0; i < len(universe[0]); i++ {
		ey = append(ey, n)
		for j := 0; j < len(universe); j++ {
			if universe[j][i] != 0 {
				continue loop2
			}
		}
		n++
	}

	return pairs, ex, ey, scanner.Err()
}

func solve(pairs []pair, ex, ey []int, factor int) int {
	sum := 0
	for _, p := range pairs {
		sum += manhattan(p.a, p.b, ex, ey, factor-1)
	}
	return sum
}

func manhattan(a, b pos, ex, ey []int, factor int) int {
	x := (a.x + ex[a.x]*factor) - (b.x + ex[b.x]*factor)
	if x < 0 {
		x *= -1
	}

	y := (a.y + ey[a.y]*factor) - (b.y + ey[b.y]*factor)
	if y < 0 {
		y *= -1
	}

	return x + y
}
