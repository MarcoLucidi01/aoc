// https://adventofcode.com/2024/day/8
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

func main() {
	grid, err := readGrid(os.Stdin)
	if err != nil {
		panic(err)
	}

	p1, p2 := solve(grid)
	fmt.Println("part1:", p1)
	fmt.Println("part2:", p2)
}

func readGrid(r io.Reader) ([][]byte, error) {
	scanner := bufio.NewScanner(r)

	var grid [][]byte
	for scanner.Scan() {
		grid = append(grid, []byte(strings.TrimSpace(scanner.Text())))
	}

	return grid, scanner.Err()
}

func solve(grid [][]byte) (int, int) {
	antennas := make(map[byte][]pos)
	for i := range grid {
		for j, a := range grid[i] {
			if a != '.' {
				antennas[a] = append(antennas[a], pos{i, j})
			}
		}
	}

	p1, p2 := make(map[pos]struct{}), make(map[pos]struct{})
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				for a, b, x := v[i], v[j], 0; !isOutside(grid, a); x++ {
					if x == 1 {
						p1[a] = struct{}{}
					}
					p2[a] = struct{}{}
					_, n := antinodes(a, b)
					b = a
					a = n
				}

				for a, b, x := v[i], v[j], 0; !isOutside(grid, b); x++ {
					if x == 1 {
						p1[b] = struct{}{}
					}
					p2[b] = struct{}{}
					m, _ := antinodes(a, b)
					a = b
					b = m
				}
			}
		}
	}
	return len(p1), len(p2)
}

func antinodes(a, b pos) (pos, pos) {
	i, j := b.i-a.i, a.j-b.j
	return pos{b.i + i, b.j - j}, pos{a.i - i, a.j + j}
}

func isOutside(grid [][]byte, p pos) bool {
	return p.i < 0 || p.i >= len(grid) || p.j < 0 || p.j >= len(grid[p.i])
}
