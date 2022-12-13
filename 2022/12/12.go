// https://adventofcode.com/2022/day/12
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type pos struct {
	x, y int
}

func main() {
	hmap, start, end, err := readHeightMap(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bfs(hmap, start, end))
	fmt.Println(part2(hmap, end))
}

func readHeightMap(r io.Reader) ([][]rune, pos, pos, error) {
	scanner := bufio.NewScanner(r)
	var hmap [][]rune
	var start pos
	var end pos

	for y := 0; scanner.Scan(); y++ {
		s := []rune(strings.TrimSpace(scanner.Text()))
		for x := range s {
			switch s[x] {
			case 'S':
				start = pos{x, y}
				s[x] = 'a'
			case 'E':
				end = pos{x, y}
				s[x] = 'z'
			}
		}
		hmap = append(hmap, s)
	}

	return hmap, start, end, scanner.Err()
}

func bfs(hmap [][]rune, start, end pos) int {
	queue := []pos{start}
	dist := map[pos]int{}
	dist[start] = 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p == end {
			return dist[p]
		}

		for _, incr := range []pos{{+1, 0}, {-1, 0}, {0, +1}, {0, -1}} {
			adj := pos{p.x + incr.x, p.y + incr.y}

			if adj.x < 0 || adj.x >= len(hmap[0]) || adj.y < 0 || adj.y >= len(hmap) {
				continue
			}
			if _, visited := dist[adj]; visited {
				continue
			}
			if hmap[adj.y][adj.x] > hmap[p.y][p.x]+1 {
				continue
			}

			dist[adj] = dist[p] + 1
			queue = append(queue, adj)
		}
	}

	return math.MaxInt
}

func part2(hmap [][]rune, end pos) int {
	min := math.MaxInt
	for y := range hmap {
		for x, h := range hmap[y] {
			if h != 'a' {
				continue
			}
			dist := bfs(hmap, pos{x, y}, end)
			if dist < min {
				min = dist
			}
		}
	}
	return min
}
