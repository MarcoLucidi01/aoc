// https://adventofcode.com/2021/day/21
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type game struct {
	turn1  bool
	pos1   int
	score1 int
	pos2   int
	score2 int
}

func main() {
	start1, start2, err := readStartPositions()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("part 1: %d\n", partOne(start1, start2))
	fmt.Printf("part 2: %d\n", partTwo(start1, start2))
}

func readStartPositions() (int, int, error) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	start1, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		return 0, 0, err
	}
	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	start2, err := strconv.Atoi(fields[len(fields)-1])
	if err != nil {
		return 0, 0, err
	}
	return start1, start2, scanner.Err()
}

func partOne(start1, start2 int) int {
	const (
		maxScore = 1000
		nrolls   = 3
		nsides   = 100
		npos     = 10
	)
	g := game{turn1: true, pos1: start1, score1: 0, pos2: start2, score2: 0}
	nrolled := 0
	for dice := 0; g.score1 < maxScore && g.score2 < maxScore; g.turn1 = !g.turn1 {
		steps := 0
		for i := 0; i < nrolls; i++ {
			dice = dice%nsides + 1
			steps += dice
		}
		nrolled += nrolls
		if g.turn1 {
			g.pos1 = (g.pos1+steps-1)%npos + 1
			g.score1 += g.pos1
		} else {
			g.pos2 = (g.pos2+steps-1)%npos + 1
			g.score2 += g.pos2
		}
	}
	if g.score1 < g.score2 {
		return g.score1 * nrolled
	}
	return g.score2 * nrolled
}

func partTwo(start1, start2 int) uint64 {
	const (
		maxScore = 21
		nrolls   = 3
		nsides   = 3
		npos     = 10
	)
	// would be nice if this was generated according to nrolls instead of
	// hardcoding 3 loops.
	rolls := make(map[int]uint64)
	for i := 1; i <= nsides; i++ {
		for j := 1; j <= nsides; j++ {
			for k := 1; k <= nsides; k++ {
				rolls[i+j+k]++
			}
		}
	}

	memo := make(map[game][2]uint64)
	var playDirac func(g game) (uint64, uint64)
	playDirac = func(g game) (uint64, uint64) {
		if g.score1 >= maxScore {
			return 1, 0
		}
		if g.score2 >= maxScore {
			return 0, 1
		}
		wins1, wins2 := uint64(0), uint64(0)
		for steps, rep := range rolls {
			gcopy := g
			if g.turn1 {
				gcopy.pos1 = (gcopy.pos1+steps-1)%npos + 1
				gcopy.score1 += gcopy.pos1
			} else {
				gcopy.pos2 = (gcopy.pos2+steps-1)%npos + 1
				gcopy.score2 += gcopy.pos2
			}
			gcopy.turn1 = !g.turn1
			wins, ok := memo[gcopy]
			if !ok {
				w1, w2 := playDirac(gcopy)
				wins = [2]uint64{w1, w2}
				memo[gcopy] = wins
			}
			wins1 += wins[0] * rep
			wins2 += wins[1] * rep
		}
		return wins1, wins2
	}

	wins1, wins2 := playDirac(game{turn1: true, pos1: start1, score1: 0, pos2: start2, score2: 0})
	if wins1 > wins2 {
		return wins1
	}
	return wins2
}
