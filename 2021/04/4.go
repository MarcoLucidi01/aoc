// https://adventofcode.com/2021/day/4
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type cell struct {
	n      int
	marked bool
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	numbers := readNumbers(scanner)
	boards := readBoards(scanner)
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	fmt.Printf("part 1: %d\n", partOne(numbers, boards))
	reset(boards)
	fmt.Printf("part 2: %d\n", partTwo(numbers, boards))
}

func readNumbers(scanner *bufio.Scanner) []int {
	scanner.Scan()
	var numbers []int
	for _, s := range strings.Split(strings.TrimSpace(scanner.Text()), ",") {
		if n, err := strconv.Atoi(s); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

func readBoards(scanner *bufio.Scanner) [][][]cell {
	var boards [][][]cell
	var board [][]cell
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) == 0 {
			if len(board) > 0 {
				boards = append(boards, board)
				board = nil
			}
			continue
		}
		var row []cell
		for _, f := range fields {
			if n, err := strconv.Atoi(f); err == nil {
				row = append(row, cell{n: n})
			}
		}
		board = append(board, row)
	}
	if len(board) > 0 {
		boards = append(boards, board)
	}
	return boards
}

func partOne(numbers []int, boards [][][]cell) int {
	for _, n := range numbers {
		for _, board := range boards {
			if markAndWins(board, n) {
				return sumUnmarked(board) * n
			}
		}
	}
	return -1
}

func markAndWins(board [][]cell, drawn int) bool {
	for row, _ := range board {
		for col, _ := range board[row] {
			if board[row][col].n == drawn {
				board[row][col].marked = true
				return isRowCompleted(board, row) || isColCompleted(board, col)
			}
		}
	}
	return false
}

func isRowCompleted(board [][]cell, row int) bool {
	for col, _ := range board[row] {
		if !board[row][col].marked {
			return false
		}
	}
	return true
}

func isColCompleted(board [][]cell, col int) bool {
	for row, _ := range board {
		if !board[row][col].marked {
			return false
		}
	}
	return true
}

func sumUnmarked(board [][]cell) int {
	sum := 0
	for row, _ := range board {
		for col, _ := range board[row] {
			if !board[row][col].marked {
				sum += board[row][col].n
			}
		}
	}
	return sum
}

func reset(boards [][][]cell) {
	for _, board := range boards {
		for row, _ := range board {
			for col, _ := range board[row] {
				board[row][col].marked = false
			}
		}
	}
}

func partTwo(numbers []int, boards [][][]cell) int {
	winners := make(map[int]bool)
	for _, n := range numbers {
		for i, board := range boards {
			if winners[i] || !markAndWins(board, n) {
				continue
			}
			winners[i] = true
			if len(winners) == len(boards) {
				return sumUnmarked(board) * n
			}
		}
	}
	return -1
}
