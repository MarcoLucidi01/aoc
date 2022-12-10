// https://adventofcode.com/2022/day/10
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

func main() {
	strength, frame, err := run(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(strength)
	fmt.Print(frame)
}

func run(r io.Reader) (int, string, error) {
	x := 1
	clock := 0
	strength := 0
	i := 0
	var frame strings.Builder

	sumStrength := func() {
		if clock%40 == 20 {
			strength += clock * x
		}
	}

	draw := func() {
		if i == x-1 || i == x || i == x+1 {
			frame.WriteRune('#')
		} else {
			frame.WriteRune('.')
		}

		i = (i + 1) % 40
		if i == 0 {
			frame.WriteRune('\n')
		}
	}

	tick := func() {
		clock++
		sumStrength()
		draw()
	}

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		inst := strings.Fields(strings.TrimSpace(scanner.Text()))
		tick()

		if inst[0] == "addx" {
			if len(inst) != 2 {
				return 0, "", fmt.Errorf("%v: missing number", inst)
			}

			n, err := strconv.Atoi(inst[1])
			if err != nil {
				return 0, "", err
			}

			tick()
			x += n
		}
	}
	return strength, frame.String(), scanner.Err()
}
