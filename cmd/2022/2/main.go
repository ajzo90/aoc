package main

import (
	"aoc2022/pkg/io"
	"fmt"
	"os"
)

var (
	gameScoreTbl  = []int{3, 0, 6}
	shapeScoreTbl = []int{1, 2, 3}
)

func score(c1, c2 uint8) int {
	return gameScoreTbl[(3+c1-c2)%3]
}

func ex1(c1, c2 uint8) int {
	c1, c2 = c1-'A', c2-'X'
	return score(c1, c2) + shapeScoreTbl[c2]
}

func ex2(c1, c2 uint8) int {
	c1 -= 'A'
	myPts := []int{0, 3, 6}[c2-'X']

	for _, c2 := range []uint8{0, 1, 2} {
		if score(c1, c2) == myPts {
			return myPts + shapeScoreTbl[c2]
		}
	}
	return 0
}

func main() {
	var fn = ex2
	var sum int
	io.ReadLines(os.Stdin, func(line []byte) {
		c1, c2 := line[0], line[2]
		sum += fn(c1, c2)
	})
	fmt.Println(sum)
}
