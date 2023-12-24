// https://adventofcode.com/2023/day/10
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	maze, start, err := readMaze(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	p1, loop := part1(maze, start)
	fmt.Println("part1:", p1)
	fmt.Println("part2:", part2(maze, loop))
}

type pos struct {
	x, y int
}

func readMaze(r io.Reader) ([]string, pos, error) {
	var maze []string
	var start pos
	scanner := bufio.NewScanner(r)
	for x := 0; scanner.Scan(); x++ {
		s := strings.TrimSpace(scanner.Text())
		if s == "" {
			continue
		}
		for y, r := range s {
			if r == 'S' {
				start.x = x
				start.y = y
			}
		}
		maze = append(maze, s)
	}
	return maze, start, scanner.Err()
}

func part1(maze []string, start pos) (int, map[pos]bool) {
	visited := make(map[pos]bool)
	p := start
	for {
		pipe := maze[p.x][p.y]
		visited[p] = true

		up := pos{x: p.x - 1, y: p.y}
		if up.x >= 0 && !visited[up] {
			q := maze[up.x][up.y]
			switch pipe {
			case 'S', '|', 'L', 'J':
				if q == '|' || q == '7' || q == 'F' {
					p = up
					continue
				}
			}
		}

		down := pos{x: p.x + 1, y: p.y}
		if down.x < len(maze) && !visited[down] {
			q := maze[down.x][down.y]
			switch pipe {
			case 'S', '|', '7', 'F':
				if q == '|' || q == 'L' || q == 'J' {
					p = down
					continue
				}
			}
		}

		left := pos{x: p.x, y: p.y - 1}
		if left.y >= 0 && !visited[left] {
			q := maze[left.x][left.y]
			switch pipe {
			case 'S', '-', 'J', '7':
				if q == '-' || q == 'L' || q == 'F' {
					p = left
					continue
				}
			}
		}

		right := pos{x: p.x, y: p.y + 1}
		if right.y < len(maze[0]) && !visited[right] {
			q := maze[right.x][right.y]
			switch pipe {
			case 'S', '-', 'L', 'F':
				if q == '-' || q == 'J' || q == '7' {
					p = right
					continue
				}
			}
		}

		if up == start || down == start || left == start || right == start {
			break
		}
	}

	return len(visited) / 2, visited
}

func part2(maze []string, loop map[pos]bool) int {
	// TODO
	return 0
}
