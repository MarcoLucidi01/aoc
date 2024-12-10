// https://adventofcode.com/2024/day/10
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
	topoMap, err := readTopoMap(os.Stdin)
	if err != nil {
		panic(err)
	}

	score, rating := analyze(topoMap)
	fmt.Println("part1:", score)
	fmt.Println("part2:", rating)
}

func readTopoMap(r io.Reader) (map[pos]int, error) {
	scanner := bufio.NewScanner(r)
	topoMap := make(map[pos]int)
	for i := 0; scanner.Scan(); i++ {
		for j, r := range strings.TrimSpace(scanner.Text()) {
			topoMap[pos{i, j}] = int(r - '0')
		}
	}
	return topoMap, scanner.Err()
}

func analyze(topoMap map[pos]int) (int, int) {
	score, rating := 0, 0
	for p, h := range topoMap {
		if h != 0 {
			continue
		}
		tops := make(map[pos]int)
		findPath(topoMap, p, tops)

		score += len(tops)
		for _, n := range tops {
			rating += n
		}
	}
	return score, rating
}

func findPath(topoMap map[pos]int, p pos, tops map[pos]int) {
	h := topoMap[p]
	if h == 9 {
		tops[p]++
		return
	}

	next := []pos{
		{p.i - 1, p.j}, // up
		{p.i + 1, p.j}, // down
		{p.i, p.j - 1}, // left
		{p.i, p.j + 1}, // right
	}
	for _, np := range next {
		if topoMap[np] == h+1 {
			findPath(topoMap, np, tops)
		}
	}
}
