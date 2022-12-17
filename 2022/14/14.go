// https://adventofcode.com/2022/day/14
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

func main() {
	rocks, height, err := readCave(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	part1, part2 := simulateFall(rocks, height)
	fmt.Println(part1)
	fmt.Println(part2)
}

func readCave(r io.Reader) (map[point]rune, int, error) {
	rocks := map[point]rune{}
	height := 0
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		fields := strings.FieldsFunc(s, func(r rune) bool {
			return r == ' ' || r == ',' || r == '-' || r == '>'
		})
		var path []int
		for _, s := range fields {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, 0, err
			}
			path = append(path, n)
		}

		p0 := point{x: path[0], y: path[1]}
		for i := 2; i < len(path); i += 2 {
			p1 := point{x: path[i], y: path[i+1]}

			s, e := p0.x, p1.x
			if s > e {
				s, e = e, s
			}
			for x := s; x <= e; x++ {
				rocks[point{x, p0.y}] = '#'
			}

			s, e = p0.y, p1.y
			if s > e {
				s, e = e, s
			}
			for y := s; y <= e; y++ {
				if y > height {
					height = y
				}
				rocks[point{p0.x, y}] = '#'
			}

			p0 = p1
		}
	}

	return rocks, height, scanner.Err()
}

func simulateFall(rocks map[point]rune, height int) (int, int) {
	start := point{500, 0}
	rest1 := -1
	rest2 := 0
	for {
		p := start
		for {
			step := point{p.x, p.y + 1}

			if step.y == height+1 && rest1 == -1 {
				rest1 = rest2 // first fall into "abyss" for part 1

			} else if step.y == height+2 {
				break // reached floor
			}

			if _, blocked := rocks[step]; !blocked {
				p = step
				continue
			}

			step.x = p.x - 1
			if _, blocked := rocks[step]; !blocked {
				p = step
				continue
			}

			step.x = p.x + 1
			if _, blocked := rocks[step]; !blocked {
				p = step
				continue
			}

			break // blocked in all directions
		}

		rocks[p] = 'o'
		rest2++
		if p == start {
			break // source blocked
		}
	}

	return rest1, rest2
}
