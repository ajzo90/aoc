package main

import (
	"aoc2022/pkg/aoc"
	"io"
)

func main() {

	aoc.
		New(2021, 3, example).
		Part(1, exResP1, func(r io.Reader) int {
			var res int
			return res
		}).
		Part(2, exResP2, func(r io.Reader) int {
			var res int
			return res
		})
}

const example = `XXX`
const exResP1 = `XXX`
const exResP2 = `XXX`
