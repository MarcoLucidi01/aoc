// https://adventofcode.com/2021/day/5
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type line struct {
	x1, y1 int
	x2, y2 int
}

func main() {
	lines, err := readLines()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", solve(lines, func(dx, dy int) bool { return dx != 0 && dy != 0 }))
	fmt.Printf("part 2: %d\n", solve(lines, func(dx, dy int) bool { return false }))
}

func readLines() ([]line, error) {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []line
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}
		points := strings.Split(s, " -> ")
		p0 := strings.Split(strings.TrimSpace(points[0]), ",")
		p1 := strings.Split(strings.TrimSpace(points[1]), ",")
		x1, _ := strconv.Atoi(p0[0])
		y1, _ := strconv.Atoi(p0[1])
		x2, _ := strconv.Atoi(p1[0])
		y2, _ := strconv.Atoi(p1[1])
		lines = append(lines, line{x1: x1, y1: y1, x2: x2, y2: y2})
	}
	return lines, scanner.Err()
}

func solve(lines []line, skip func(dx, dy int) bool) int {
	diagram := makeDiagram(lines)
	for _, l := range lines {
		dx, dy := l.x2-l.x1, l.y2-l.y1
		if skip(dx, dy) {
			continue
		}
		dx, dy = sign(dx), sign(dy)
		for x, y := l.x1, l.y1; x != l.x2 || y != l.y2; x, y = x+dx, y+dy {
			diagram[y][x]++
		}
		diagram[l.y2][l.x2]++
	}
	return countOverlap(diagram)
}

func makeDiagram(lines []line) [][]int {
	w, h := 0, 0
	for _, l := range lines {
		if l.x1 > w {
			w = l.x1
		}
		if l.x2 > w {
			w = l.x2
		}
		if l.y1 > h {
			h = l.y1
		}
		if l.y2 > h {
			h = l.y2
		}
	}
	diagram := make([][]int, h+1)
	for i := 0; i < len(diagram); i++ {
		diagram[i] = make([]int, w+1)
	}
	return diagram
}

func sign(n int) int {
	if n > 0 {
		return +1
	}
	if n < 0 {
		return -1
	}
	return n
}

func countOverlap(diagram [][]int) int {
	cnt := 0
	for _, row := range diagram {
		for _, n := range row {
			if n >= 2 {
				cnt++
			}
		}
	}
	return cnt
}
