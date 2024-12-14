// https://adventofcode.com/2024/day/14
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	i, j int
}

type robot struct {
	p, v pos
}

func main() {
	robots, err := readRobots(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(robots, 101, 103))
	fmt.Println("part2:", part2(robots, 101, 103))
}

func readRobots(r io.Reader) ([]robot, error) {
	scanner := bufio.NewScanner(r)
	rep := strings.NewReplacer("p=", "", ",", " ", "v=", "")

	var robots []robot
	for scanner.Scan() {
		f := strings.Fields(rep.Replace(strings.TrimSpace(scanner.Text())))

		px, err := strconv.Atoi(f[0])
		if err != nil {
			return nil, err
		}
		py, err := strconv.Atoi(f[1])
		if err != nil {
			return nil, err
		}
		vx, err := strconv.Atoi(f[2])
		if err != nil {
			return nil, err
		}
		vy, err := strconv.Atoi(f[3])
		if err != nil {
			return nil, err
		}

		robots = append(robots, robot{pos{py, px}, pos{vy, vx}})
	}

	return robots, scanner.Err()
}

func part1(robots []robot, w, h int) int {
	space, robots := initialize(robots)
	for i := 0; i < 100; i++ {
		tick(space, robots, w, h)
	}

	a := countq(space, 0, w/2, 0, h/2)
	b := countq(space, w/2+1, w, 0, h/2)
	c := countq(space, 0, w/2, h/2+1, h)
	d := countq(space, w/2+1, w, h/2+1, h)
	return a * b * c * d
}

func initialize(initial []robot) (map[pos]int, []robot) {
	space := make(map[pos]int)
	robots := make([]robot, len(initial))
	for i, r := range initial {
		space[r.p]++
		robots[i] = r
	}
	return space, robots
}

func tick(space map[pos]int, robots []robot, w, h int) {
	for i := range robots {
		space[robots[i].p]--
		robots[i] = move(robots[i], w, h)
		space[robots[i].p]++
	}
}

func move(r robot, w, h int) robot {
	r.p.i = mod(r.p.i+r.v.i, h)
	r.p.j = mod(r.p.j+r.v.j, w)
	return r
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func countq(space map[pos]int, bw, w, bh, h int) int {
	n := 0
	for i := bh; i < h; i++ {
		for j := bw; j < w; j++ {
			n += space[pos{i, j}]
		}
	}
	return n
}

func print(sec int, space map[pos]int, w, h int) {
	for i := 0; i < h; i++ {
		fmt.Printf("%d: ", sec)
		for j := 0; j < w; j++ {
			switch n := space[pos{i, j}]; n {
			case 0:
				fmt.Print(".")
			default:
				fmt.Printf("%d", n)
			}
		}
		fmt.Println()
	}
}

func part2(robots []robot, w, h int) int {
	space, robots := initialize(robots)
	for sec := 0; ; sec++ {
		for i := 0; i < h; i++ {
			n := 0
			for j := 0; j < w; j++ {
				c := space[pos{i, j}]
				if n > 0 && c < 1 {
					break
				}

				n += c
				// 10 consecutive robots are enough to detect the christmas tree
				//
				// I initially found the tree by printing each "frame" and
				// piping to `grep -E '[0-9]{10}'` :)
				if n >= 10 {
					print(sec, space, w, h)
					return sec
				}
			}
		}

		tick(space, robots, w, h)
	}
}
