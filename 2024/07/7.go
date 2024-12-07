// https://adventofcode.com/2024/day/7
package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	equations, err := readEquations(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", solve(equations, []rune{'+', '*'}))
	fmt.Println("part2:", solve(equations, []rune{'+', '*', '|'}))
}

func readEquations(r io.Reader) ([][]int, error) {
	scanner := bufio.NewScanner(r)

	var equations [][]int
	for scanner.Scan() {
		f := strings.Fields(strings.TrimSpace(scanner.Text()))
		f[0] = strings.TrimSuffix(f[0], ":")
		var nums []int
		for _, s := range f {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			nums = append(nums, n)
		}
		equations = append(equations, nums)
	}

	return equations, scanner.Err()
}

func solve(equations [][]int, ops []rune) int {
	tot := 0
	for _, eq := range equations {
		w := len(eq) - 2
		x := int(math.Pow(float64(len(ops)), float64(w)))
		for i := 0; i < x; i++ {
			n := i
			res := eq[1]
			for j := 0; j < w; j++ {
				op := ops[n%len(ops)]
				n /= len(ops)
				res = do(op, res, eq[j+2])
			}
			if res == eq[0] {
				tot += res
				break
			}
		}
	}
	return tot
}

func do(op rune, a, b int) int {
	switch op {
	case '+':
		return a + b
	case '*':
		return a * b
	case '|':
		for x := b; x > 0; x /= 10 {
			a *= 10
		}
		return a + b
	}
	panic("unknown op: " + string(op))
}
