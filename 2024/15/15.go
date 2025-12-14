// https://adventofcode.com/2024/day/15
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type pos struct {
	i, j int
}

func main() {
	warehouse, moves, err := readWarehouse(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println("part1:", part1(copyWarehouse(warehouse), moves))
	fmt.Println("part2:", part2(copyWarehouse(warehouse), moves))
}

func readWarehouse(r io.Reader) ([][]byte, string, error) {
	scanner := bufio.NewScanner(r)

	var warehouse [][]byte
	var moves strings.Builder
	readWarehouse := true
	for scanner.Scan() {
		switch s := strings.TrimSpace(scanner.Text()); {
		case s == "":
			readWarehouse = false
		case readWarehouse:
			warehouse = append(warehouse, []byte(s))
		default:
			moves.WriteString(s)
		}
	}

	return warehouse, moves.String(), scanner.Err()
}

func copyWarehouse(warehouse [][]byte) [][]byte {
	w := make([][]byte, 0, len(warehouse))
	for _, r := range warehouse {
		a := make([]byte, len(r))
		copy(a, r)
		w = append(w, a)
	}
	return w
}

func part1(warehouse [][]byte, moves string) int {
	robot := findRobot(warehouse)
	for _, m := range moves {
		robot = moveRobot(warehouse, robot, byte(m))
	}

	n := 0
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == 'O' {
				n += 100*i + j
			}
		}
	}

	return n
}

func findRobot(warehouse [][]byte) pos {
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '@' {
				return pos{i, j}
			}
		}
	}
	return pos{-1, -1}
}

var dir map[byte]pos = map[byte]pos{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func moveRobot(warehouse [][]byte, robot pos, m byte) pos {
	d := dir[m]
	np := pos{robot.i + d.i, robot.j + d.j}

	switch warehouse[np.i][np.j] {
	case '#':
		return robot
	case '.':
		warehouse[np.i][np.j] = '@'
		warehouse[robot.i][robot.j] = '.'
		return np
	case 'O':
		i := np.i
		j := np.j
		for {
			i += d.i
			j += d.j
			switch warehouse[i][j] {
			case 'O':
				continue
			case '.':
				warehouse[np.i][np.j] = '@'
				warehouse[robot.i][robot.j] = '.'
				warehouse[i][j] = 'O'
				return pos{np.i, np.j}
			default: // '#'
				return robot
			}
		}
	}

	panic("invalid move")
}

func part2(warehouse [][]byte, moves string) int {
	warehouse = expand(warehouse)

	robot := findRobot(warehouse)
	for _, m := range moves {
		robot = moveRobot2(warehouse, robot, byte(m))
	}

	n := 0
	for i := 0; i < len(warehouse); i++ {
		for j := 0; j < len(warehouse[i]); j++ {
			if warehouse[i][j] == '[' {
				n += 100*i + j
			}
		}
	}
	return n
}

func expand(warehouse [][]byte) [][]byte {
	var newWarehouse [][]byte
	for _, row := range warehouse {
		var newRow []byte
		for _, c := range row {
			switch c {
			case '#':
				newRow = append(newRow, '#', '#')
			case 'O':
				newRow = append(newRow, '[', ']')
			case '.':
				newRow = append(newRow, '.', '.')
			case '@':
				newRow = append(newRow, '@', '.')
			}
		}
		newWarehouse = append(newWarehouse, newRow)
	}
	return newWarehouse
}

type box struct {
	l, r pos
}

func moveBox(warehouse [][]byte, b box, m byte) {
	b1 := box{
		pos{b.l.i + dir[m].i, b.l.j + dir[m].j},
		pos{b.r.i + dir[m].i, b.r.j + dir[m].j},
	}

	if warehouse[b1.l.i][b1.l.j] == '#' || warehouse[b1.r.i][b1.r.j] == '#' {
		return
	}

	if warehouse[b1.l.i][b1.l.j] == '[' && warehouse[b1.r.i][b1.r.j] == ']' {
		if canMoveBox(warehouse, b1, m) {
			moveBox(warehouse, b1, m)
		}

	} else if warehouse[b1.l.i][b1.l.j] == ']' && warehouse[b1.r.i][b1.r.j] == '[' {
		if canMoveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m) &&
			canMoveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m) {

			moveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m)
			moveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m)
		}

	} else if warehouse[b1.l.i][b1.l.j] == ']' {
		if canMoveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m) {
			moveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m)
		}

	} else if warehouse[b1.r.i][b1.r.j] == '[' {
		if canMoveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m) {
			moveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m)
		}
	}

	if warehouse[b1.l.i][b1.l.j] == '.' && warehouse[b1.r.i][b1.r.j] == '.' {
		warehouse[b1.l.i][b1.l.j] = warehouse[b.l.i][b.l.j]
		warehouse[b1.r.i][b1.r.j] = warehouse[b.r.i][b.r.j]
		warehouse[b.l.i][b.l.j] = '.'
		warehouse[b.r.i][b.r.j] = '.'
		return
	}
}

func canMoveBox(warehouse [][]byte, b box, m byte) bool {
	b1 := box{
		pos{b.l.i + dir[m].i, b.l.j + dir[m].j},
		pos{b.r.i + dir[m].i, b.r.j + dir[m].j},
	}

	if warehouse[b1.l.i][b1.l.j] == '#' || warehouse[b1.r.i][b1.r.j] == '#' {
		return false
	}

	if warehouse[b1.l.i][b1.l.j] == '[' && warehouse[b1.r.i][b1.r.j] == ']' {
		return canMoveBox(warehouse, b1, m)

	} else if warehouse[b1.l.i][b1.l.j] == ']' && warehouse[b1.r.i][b1.r.j] == '[' {
		return canMoveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m) &&
			canMoveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m)

	} else if warehouse[b1.l.i][b1.l.j] == ']' {
		return canMoveBox(warehouse, box{pos{b1.l.i + dir['<'].i, b1.l.j + dir['<'].j}, b1.l}, m)

	} else if warehouse[b1.r.i][b1.r.j] == '[' {
		return canMoveBox(warehouse, box{b1.r, pos{b1.r.i + dir['>'].i, b1.r.j + dir['>'].j}}, m)
	}

	if warehouse[b1.l.i][b1.l.j] == '.' && warehouse[b1.r.i][b1.r.j] == '.' {
		return true
	}

	return false
}

func moveRobot2(warehouse [][]byte, robot pos, m byte) pos {
	d := dir[m]
	np := pos{robot.i + d.i, robot.j + d.j}

	switch warehouse[np.i][np.j] {
	case '#':
		return robot
	case '.':
		warehouse[np.i][np.j] = '@'
		warehouse[robot.i][robot.j] = '.'
		return np
	case '[', ']':
		if m == '<' || m == '>' {
			i := np.i
			j := np.j
			for {
				i += d.i
				j += d.j
				switch warehouse[i][j] {
				case ']', '[':
					continue
				case '.':
					x, y := np.i, np.j
					for {
						x += d.i
						y += d.j
						switch warehouse[x][y] {
						case ']':
							warehouse[x][y] = '['
						case '[':
							warehouse[x][y] = ']'
						case '.':
							if warehouse[x-d.i][y-d.j] == ']' {
								warehouse[x][y] = '['
							} else {
								warehouse[x][y] = ']'
							}
							warehouse[np.i][np.j] = '@'
							warehouse[robot.i][robot.j] = '.'
							return np
						}
					}
				default: // '#'
					return robot
				}
			}
		}

		if warehouse[np.i][np.j] == '[' {
			if canMoveBox(warehouse, box{np, pos{np.i + dir['>'].i, np.j + dir['>'].j}}, m) {
				moveBox(warehouse, box{np, pos{np.i + dir['>'].i, np.j + dir['>'].j}}, m)
				warehouse[np.i][np.j] = '@'
				warehouse[robot.i][robot.j] = '.'
				return np
			}
		} else {
			if canMoveBox(warehouse, box{pos{np.i + dir['<'].i, np.j + dir['<'].j}, np}, m) {
				moveBox(warehouse, box{pos{np.i + dir['<'].i, np.j + dir['<'].j}, np}, m)
				warehouse[np.i][np.j] = '@'
				warehouse[robot.i][robot.j] = '.'
				return np
			}
		}

		return robot
	}

	panic("invalid move")
}
