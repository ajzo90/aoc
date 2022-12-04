package main

import (
	"aoc2022/pkg/aoc"
	"io"
)

func main() {

	aoc.
		New(2021, 3, example).
		Part(1, exResP1, func(r io.Reader) int {
			var count [32][2]int
			var nCols int
			aoc.ReadLines(r, func(line []byte) {
				if nCols == 0 {
					nCols = len(line)
				}
				for i, c := range line {
					count[i][c-'0']++
				}
			})
			gamma, epsilon := 0, 0
			for i, v := range count[:nCols] {
				mask := 1 << (nCols - 1 - i)
				if v[1] > v[0] {
					gamma |= mask
				} else {
					epsilon |= mask
				}
			}

			return gamma * epsilon
		}).
		Part(2, exResP2, func(r io.Reader) int {
			var res int
			return res
		})
}

const example = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

const exResP1 = `198`
const exResP2 = `230`
