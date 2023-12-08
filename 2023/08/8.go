// https://adventofcode.com/2023/day/8
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
	dirs, nodes, err := readMap(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(dirs, nodes))
	fmt.Println("part2:", part2(dirs, nodes))
}

func readMap(r io.Reader) ([]int, map[string][2]string, error) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	s := strings.TrimSpace(scanner.Text())
	dirs := make([]int, len(s))
	for i, r := range s {
		if r == 'R' {
			dirs[i] = 1
		}
	}

	nodes := make(map[string][2]string)
	for scanner.Scan() {
		f := strings.Fields(strings.TrimSpace(scanner.Text()))
		if len(f) == 4 {
			nodes[f[0]] = [2]string{f[2][1 : len(f[2])-1], f[3][:len(f[3])-1]}
		}
	}

	return dirs, nodes, scanner.Err()
}

func part1(dirs []int, nodes map[string][2]string) int {
	steps := 0
	for node := "AAA"; node != "ZZZ"; steps++ {
		next, ok := nodes[node]
		if !ok {
			return -1
		}
		node = next[dirs[steps%len(dirs)]]
	}
	return steps
}

func part2(dirs []int, nodes map[string][2]string) int {
	var steps []int
	for node := range nodes {
		if node[len(node)-1] != 'A' {
			continue
		}

		n := 0
		for ; node[len(node)-1] != 'Z'; n++ {
			node = nodes[node][dirs[n%len(dirs)]]
		}
		steps = append(steps, n)
	}

	if len(steps) == 1 {
		return steps[0]
	}

	return lcm(steps...)
}

func lcm(n ...int) int {
	res := n[0] * n[1] / gcd(n[0], n[1])
	for i := 2; i < len(n); i++ {
		res = lcm(res, n[i])
	}
	return res
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
