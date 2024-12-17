// https://adventofcode.com/2024/day/16
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type pos struct {
	i, j int
}

type posdir struct {
	pos
	dir rune
}

const (
	north = '^'
	sud   = 'v'
	east  = '>'
	west  = '<'
)

func main() {
	maze, start, end, err := readMaze(os.Stdin)
	if err != nil {
		panic(err)
	}

	end, dist, prev := dijkstra(maze, start, end)
	fmt.Println("part1:", dist[end])
	fmt.Println("part2:", part2(end, prev))
}

func readMaze(r io.Reader) (map[pos]struct{}, posdir, posdir, error) {
	scanner := bufio.NewScanner(r)
	maze := make(map[pos]struct{})
	var start, end posdir
	for i := 0; scanner.Scan(); i++ {
		for j, r := range scanner.Text() {
			switch r {
			case 'S', 'E', '.':
				maze[pos{i, j}] = struct{}{}
			}
			switch r {
			case 'S':
				start = posdir{pos{i, j}, east}
			case 'E':
				end = posdir{pos{i, j}, east}
			}
		}
	}
	return maze, start, end, scanner.Err()
}

func dijkstra(maze map[pos]struct{}, start, end posdir) (posdir, map[posdir]int, map[posdir][]posdir) {
	queue := map[posdir]struct{}{start: {}}
	dist := map[posdir]int{start: 0}
	prev := map[posdir][]posdir{}
	visited := map[posdir]struct{}{}

	for len(queue) > 0 {
		m := posdir{pos{-1, -1}, -1} // min
		for p := range queue {
			if m.i == -1 || dist[p] < dist[m] {
				m = p
			}
		}
		delete(queue, m)
		visited[m] = struct{}{}

		if m.pos == end.pos {
			end = m // update end direction so that it can be found in dist
			break
		}

		next := []posdir{
			{pos{m.i - 1, m.j}, north},
			{pos{m.i + 1, m.j}, sud},
			{pos{m.i, m.j - 1}, east},
			{pos{m.i, m.j + 1}, west},
		}
		for _, p := range next {
			if _, ok := maze[p.pos]; !ok {
				continue
			}
			if _, ok := visited[p]; ok {
				continue
			}
			queue[p] = struct{}{}

			nd := dist[m] + 1
			if m.dir != p.dir {
				nd += 1000
			}

			d, ok := dist[p]
			if ok && nd > d {
				continue
			}
			// keep track of all previous nodes at the same distance to find all
			// shortest paths for part2
			if nd < d {
				prev[p] = nil // reset previous nodes when we find a shorter distance
			}
			dist[p] = nd
			prev[p] = append(prev[p], m)
		}
	}

	return end, dist, prev
}

func part2(end posdir, prev map[posdir][]posdir) int {
	all := map[pos]struct{}{end.pos: {}}
	for q := []posdir{end}; len(q) > 0; q = q[1:] {
		for _, p := range prev[q[0]] {
			q = append(q, p)
			all[p.pos] = struct{}{}
		}
	}
	return len(all)
}
