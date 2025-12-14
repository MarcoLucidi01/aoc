// https://adventofcode.com/2025/day/6
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	problems, problems2, err := readProblems(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Printf("part1: %d\n", part1(problems))
	fmt.Printf("part2: %d\n", part2(problems, problems2))
}

func readProblems(r io.Reader) ([][]int, [][]byte, error) {
	scanner := bufio.NewScanner(r)

	var problems [][]int
	var problems2 [][]byte
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		var row []int
		for _, f := range fields {
			n := 0
			switch f {
			case "+":
				n = 0
			case "*":
				n = 1
			default:
				var err error
				if n, err = strconv.Atoi(f); err != nil {
					return nil, nil, err
				}
			}
			row = append(row, n)
		}
		problems = append(problems, row)
		problems2 = append(problems2, []byte(scanner.Text()))
	}

	return problems, problems2, scanner.Err()
}

func part1(problems [][]int) int {
	tot := 0
	ops := problems[len(problems)-1]
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		n := op
		for j := 0; j < len(problems)-1; j++ {
			switch op {
			case 0:
				n += problems[j][i]
			case 1:
				n *= problems[j][i]
			}
		}
		tot += n
	}

	return tot
}

func part2(problems [][]int, problems2 [][]byte) int {
	tot := 0
	ops := problems[len(problems)-1]
	j := len(ops) -1
	var numbers []int
	for i := len(problems2[0])-1; i >= 0; i-- {
		var number []byte
		for j := 0; j < len(problems2)-1; j++ {
			number = append(number, problems2[j][i])
		}

		allSpaces := true
		for _, d := range number {
			if d != ' ' {
				allSpaces = false
				break
			}
		}
		if !allSpaces {
			n, _ := strconv.Atoi(strings.TrimSpace(string(number)))
			numbers = append(numbers, n)
			if i > 0 {
				continue
			}
		}

		op := ops[j]
		result := op
		for _, n := range numbers {
			switch op {
			case 0:
				result += n
			case 1:
				result *= n
			}
		}
		tot += result

		numbers = nil
		j--
	}

	return tot
}
