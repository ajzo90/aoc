package main

import (
	"aoc2022/pkg/aoc"
	"bytes"
	"io"
)

var digits = [][]byte{
	[]byte("one"),
	[]byte("two"),
	[]byte("three"),
	[]byte("four"),
	[]byte("five"),
	[]byte("six"),
	[]byte("seven"),
	[]byte("eight"),
	[]byte("nine"),
}

func main() {

	var xxx = func(r io.Reader) int {
		var sum int
		aoc.ReadLines(r, func(line []byte) {
			var first, last *int
			for idx, c := range line {
				if c >= '0' && c <= '9' {
					v := int(c - '0')
					last = &v
					if first == nil {
						first = &v
					}
				}
				suffix := line[idx:]
				for num, d := range digits {
					if bytes.Equal(suffix[:len(d)], d) {
						v := num + 1
						last = &v
						if first == nil {
							first = &v
						}
					}
				}
			}
			sum += *first*10 + *last
		})
		return sum
	}
	aoc.
		New(2023, 1, example).
		Part(1, "281", xxx).
		Part(2, "281", xxx)
}

const example = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`
