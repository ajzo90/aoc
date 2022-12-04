package main

import (
	"aoc2022/pkg/aoc"
	"io"
	"strconv"
)

func countIncrements(l []int) int {
	var last = l[0]
	var count int
	for _, v := range l {
		if v > last {
			count++
		}
		last = v
	}
	return count
}

func main() {

	aoc.
		New(2021, 1, example).
		Part(1, "7", func(r io.Reader) string {
			return strconv.Itoa(countIncrements(aoc.ReadLinesToIntList(r)))
		}).
		Part(2, "5", func(r io.Reader) string {
			l := aoc.ReadLinesToIntList(r)
			return strconv.Itoa(countIncrements(aoc.SlidingSum(l, 3)))
		})
}

const example = `199
200
208
210
200
207
240
269
260
263`
