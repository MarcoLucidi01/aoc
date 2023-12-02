// https://adventofcode.com/2023/day/2
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
	games, err := parseGames(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(games))
	fmt.Println("part2:", part2(games))
}

func parseGames(r io.Reader) ([][][3]int, error) {
	scanner := bufio.NewScanner(r)

	var games [][][3]int
	for scanner.Scan() {
		s := strings.TrimSpace(scanner.Text())
		i := strings.IndexRune(s, ':')
		sets := strings.Split(s[i+1:], ";")

		var game [][3]int
		for _, set := range sets {
			cubes := strings.Split(set, ",")

			var colors [3]int
			for _, cube := range cubes {
				split := strings.Split(strings.TrimSpace(cube), " ")

				n, err := strconv.Atoi(split[0])
				if err != nil {
					return nil, err
				}

				switch split[1] {
				case "red":
					colors[0] += n
				case "green":
					colors[1] += n
				case "blue":
					colors[2] += n
				}
			}
			game = append(game, colors)
		}
		games = append(games, game)
	}

	return games, scanner.Err()
}

func part1(games [][][3]int) int {
	sum := 0
loop:
	for i, g := range games {
		for _, c := range g {
			if c[0] > 12 || c[1] > 13 || c[2] > 14 {
				continue loop
			}
		}
		sum += i + 1
	}

	return sum
}

func part2(games [][][3]int) int {
	sum := 0
	for _, g := range games {
		var m [3]int
		for _, c := range g {
			if c[0] > m[0] {
				m[0] = c[0]
			}
			if c[1] > m[1] {
				m[1] = c[1]
			}
			if c[2] > m[2] {
				m[2] = c[2]
			}
		}
		sum += m[0] * m[1] * m[2]
	}

	return sum
}
