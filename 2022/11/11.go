// https://adventofcode.com/2022/day/11
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	reOp      = regexp.MustCompile(`^Operation: new = old ([+*]) (\d+|old)$`)
	reTest    = regexp.MustCompile(`^Test: divisible by (\d+)$`)
	reIfTrue  = regexp.MustCompile(`^If true: throw to monkey (\d+)$`)
	reIfFalse = regexp.MustCompile(`^If false: throw to monkey (\d+)$`)
)

type monkey struct {
	items   []int64
	op      func(worry int64) int64
	test    int64
	ifTrue  int
	ifFalse int
}

func main() {
	monkeys, err := readMonkeys(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(part1(clone(monkeys)))
	fmt.Println(part2(clone(monkeys)))
}

func readMonkeys(r io.Reader) ([]monkey, error) {
	scanner := bufio.NewScanner(r)

	next := func() string {
		scanner.Scan()
		return strings.TrimSpace(scanner.Text())
	}

	var monkeys []monkey
	var err error
	for {
		var m monkey

		if next() == "" && !scanner.Scan() { // skip monkey name and empty line
			break
		}

		// starting items line
		items := next()
		items = strings.ReplaceAll(items, "Starting items: ", "")
		nums := strings.Split(items, ", ")
		for _, s := range nums {
			n, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				return nil, err
			}
			m.items = append(m.items, n)
		}

		// operation line
		op := next()
		match := reOp.FindStringSubmatch(op)
		if len(match) != 3 {
			return nil, fmt.Errorf("unexpected %v", op)
		}

		n := int64(0)
		if match[2] != "old" {
			n, err = strconv.ParseInt(match[2], 10, 64)
			if err != nil {
				return nil, fmt.Errorf("unexpected %v", op)
			}
		}

		switch match[1] {
		case "*":
			m.op = func(worry int64) int64 { return worry * n }
			if match[2] == "old" {
				m.op = func(worry int64) int64 { return worry * worry }
			}
		case "+":
			m.op = func(worry int64) int64 { return worry + n }
			if match[2] == "old" {
				m.op = func(worry int64) int64 { return worry + worry }
			}
		}

		// test line
		test := next()
		match = reTest.FindStringSubmatch(test)
		if len(match) != 2 {
			return nil, fmt.Errorf("unexpected %v", test)
		}
		if m.test, err = strconv.ParseInt(match[1], 10, 64); err != nil {
			return nil, err
		}

		// if true line
		ifTrue := next()
		match = reIfTrue.FindStringSubmatch(ifTrue)
		if len(match) != 2 {
			return nil, fmt.Errorf("unexpected %v", ifTrue)
		}
		if m.ifTrue, err = strconv.Atoi(match[1]); err != nil {
			return nil, err
		}

		// if false line
		ifFalse := next()
		match = reIfFalse.FindStringSubmatch(ifFalse)
		if len(match) != 2 {
			return nil, fmt.Errorf("unexpected %v", ifFalse)
		}
		if m.ifFalse, err = strconv.Atoi(match[1]); err != nil {
			return nil, err
		}

		monkeys = append(monkeys, m)
	}

	return monkeys, scanner.Err()
}

func clone(monkeys []monkey) []monkey {
	c := make([]monkey, len(monkeys))
	copy(c, monkeys)
	return c
}

func part1(monkeys []monkey) int {
	return rounds(monkeys, 20, func(worry int64) int64 { return worry / 3 })
}

func part2(monkeys []monkey) int {
	lcm := int64(1)
	for _, m := range monkeys {
		// this works because the test numbers are all primes
		lcm *= m.test
	}
	return rounds(monkeys, 10000, func(worry int64) int64 { return worry % lcm })
}

func rounds(monkeys []monkey, n int, lowerWorry func(worry int64) int64) int {
	inspected := make([]int, len(monkeys))

	for i := 0; i < n; i++ {
		for j := range monkeys {
			m := &monkeys[j]
			for len(m.items) > 0 {
				worry := m.items[0]
				m.items = m.items[1:]

				inspected[j]++

				worry = m.op(worry)
				worry = lowerWorry(worry)

				x := m.ifFalse
				if worry%m.test == 0 {
					x = m.ifTrue
				}
				monkeys[x].items = append(monkeys[x].items, worry)
			}
		}
	}

	sort.Ints(inspected)
	return inspected[len(inspected)-1] * inspected[len(inspected)-2]
}
