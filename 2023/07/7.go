// https://adventofcode.com/2023/day/7
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

const jokerValue = 11

var cardValue = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': jokerValue,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

type hand struct {
	cards  map[rune]int
	values []int
	typ    int
	bid    int
}

func main() {
	hands, err := readHands(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part1:", part1(hands))
	fmt.Println("part2:", part2(hands))
}

func readHands(r io.Reader) ([]hand, error) {
	scanner := bufio.NewScanner(r)

	var hands []hand
	for scanner.Scan() {
		f := strings.Fields(strings.TrimSpace(scanner.Text()))

		cards := make(map[rune]int)
		var values []int
		for _, c := range f[0] {
			cards[c]++
			values = append(values, cardValue[c])
		}

		bid, err := strconv.Atoi(f[1])
		if err != nil {
			return nil, err
		}

		hands = append(hands, hand{
			cards:  cards,
			values: values,
			typ:    handType(cards),
			bid:    bid,
		})
	}

	return hands, scanner.Err()
}

func handType(cards map[rune]int) int {
	switch len(cards) {
	case 1:
		return fiveOfAKind
	case 2:
		for _, n := range cards {
			switch n {
			case 4, 1:
				return fourOfAKind
			case 3, 2:
				return fullHouse
			}
		}
	case 3:
		for _, n := range cards {
			switch n {
			case 3:
				return threeOfAKind
			case 2:
				return twoPair
			}
		}
	case 4:
		return onePair
	case 5:
		return highCard
	}

	panic(fmt.Sprintf("wrong hand symbols %v", cards))
}

func part1(hands []hand) int {
	sort.Slice(hands, func(i, j int) bool {
		a, b := hands[i], hands[j]
		if a.typ < b.typ {
			return true
		}
		if a.typ > b.typ {
			return false
		}
		for x := 0; x < len(a.values); x++ {
			cmp := a.values[x] - b.values[x]
			if cmp == 0 {
				continue
			}
			return cmp < 0
		}
		return false
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}

	return sum
}

type jokerHand struct {
	hand
	jokerValues []int
	jokerType   int
}

func part2(hands []hand) int {
	var jokerHands []jokerHand
	for _, h := range hands {
		jokerHands = append(jokerHands, toJokerHand(h))
	}

	sort.Slice(jokerHands, func(i, j int) bool {
		a, b := jokerHands[i], jokerHands[j]
		if a.jokerType < b.jokerType {
			return true
		}
		if a.jokerType > b.jokerType {
			return false
		}
		for x := 0; x < len(a.jokerValues); x++ {
			cmp := a.jokerValues[x] - b.jokerValues[x]
			if cmp == 0 {
				continue
			}
			return cmp < 0
		}
		return false
	})

	sum := 0
	for i, h := range jokerHands {
		sum += h.bid * (i + 1)
	}

	return sum
}

func toJokerHand(h hand) jokerHand {
	jh := jokerHand{
		hand:        h,
		jokerValues: h.values,
		jokerType:   h.typ,
	}

	if h.cards['J'] == 0 {
		return jh
	}

	jh.jokerValues = make([]int, len(h.values))
	copy(jh.jokerValues, h.values)
	for i, v := range jh.jokerValues {
		if v == jokerValue {
			jh.jokerValues[i] = 0
		}
	}

	if h.typ == fiveOfAKind {
		return jh
	}

	jokerCards := make(map[rune]int)
	var highC rune
	var highN int
	for c, n := range h.cards {
		if c == 'J' {
			continue
		}

		if n > highN || (n == highN && cardValue[c] > cardValue[highC]) {
			highC = c
			highN = n
		}

		jokerCards[c] = n
	}

	jokerCards[highC] += h.cards['J']

	jh.jokerType = handType(jokerCards)
	return jh
}
