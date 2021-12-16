// https://adventofcode.com/2021/day/16
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type exType uint64

const (
	exSum exType = iota
	exProd
	exMin
	exMax
	exLiteral
	exGt
	exLt
	exEq
)

var toBinWord = map[byte]string{
	'0': "0000",
	'1': "0001",
	'2': "0010",
	'3': "0011",
	'4': "0100",
	'5': "0101",
	'6': "0110",
	'7': "0111",
	'8': "1000",
	'9': "1001",
	'A': "1010",
	'B': "1011",
	'C': "1100",
	'D': "1101",
	'E': "1110",
	'F': "1111",
}

type bitsReader struct {
	r     *bufio.Reader
	nread uint64 // number of bits read til now
	word  string // current decoded word
	i     int    // current index in word
}

type expr struct {
	ver uint64 // version number
	typ exType // type id
	n   uint64 // literal number if typ == exLiteral
	sub []expr // sub expressions of this expression
}

func main() {
	b := newBitsReader(os.Stdin)
	ex := b.readExpr()
	fmt.Printf("part 1: %d\n", sumVerNumbers(ex))
	fmt.Printf("part 2: %d\n", eval(ex))
}

func newBitsReader(r io.Reader) *bitsReader {
	return &bitsReader{r: bufio.NewReader(r)}
}

func (b *bitsReader) readExpr() expr {
	ex := expr{ver: b.readUint(3)}

	if ex.typ = exType(b.readUint(3)); ex.typ == exLiteral {
		for {
			g := b.readUint(5)
			ex.n = (ex.n << 4) | (g & 0b1111)
			if g&0b10000 == 0 {
				return ex
			}
		}
	}

	if b.readBit() == 1 { // length type ID
		nsub := b.readUint(11)
		for i := uint64(0); i < nsub; i++ {
			ex.sub = append(ex.sub, b.readExpr())
		}
		return ex
	}

	nbits := b.readUint(15)
	for stop := b.nread + nbits; b.nread < stop; {
		ex.sub = append(ex.sub, b.readExpr())
	}
	return ex
}

func (b *bitsReader) readUint(size int) uint64 {
	n := uint64(0)
	for i := 0; i < size; i++ {
		n = (n << 1) | b.readBit()
	}
	return n
}

func (b *bitsReader) readBit() uint64 {
	if b.i == len(b.word) {
		hex, err := b.r.ReadByte()
		if err != nil {
			die(fmt.Sprintf("readBit: unexpected error %s", err))
		}
		var ok bool
		if b.word, ok = toBinWord[hex]; !ok {
			die(fmt.Sprintf("readBit: unexpected char %c", hex))
		}
		b.i = 0
	}
	bit := uint64(0)
	if b.word[b.i] == '1' {
		bit = 1
	}
	b.i++
	b.nread++
	return bit
}

func die(err string) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(1)
}

func sumVerNumbers(ex expr) uint64 {
	sum := ex.ver
	for _, sub := range ex.sub {
		sum += sumVerNumbers(sub)
	}
	return sum
}

func eval(ex expr) uint64 {
	if ex.typ == exLiteral {
		return ex.n
	}
	if len(ex.sub) == 0 {
		die(fmt.Sprintf("eval: expression %d has no sub-expressions", ex.typ))
	}
	ret := eval(ex.sub[0])
	for _, sub := range ex.sub[1:] {
		switch ex.typ {
		case exSum:
			ret += eval(sub)
		case exProd:
			ret *= eval(sub)
		case exMin:
			if v := eval(sub); v < ret {
				ret = v
			}
		case exMax:
			if v := eval(sub); v > ret {
				ret = v
			}
		case exGt:
			if ret > eval(sub) {
				ret = 1
			} else {
				ret = 0
			}
		case exLt:
			if ret < eval(sub) {
				ret = 1
			} else {
				ret = 0
			}
		case exEq:
			if ret == eval(sub) {
				ret = 1
			} else {
				ret = 0
			}
		default:
			die(fmt.Sprintf("eval: unknown expression type %d", ex.typ))
		}
	}
	return ret
}
