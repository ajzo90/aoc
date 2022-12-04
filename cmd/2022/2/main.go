package main

import (
	"aoc2022/pkg/aoc"
	"fmt"
	"os"
)

var (
	gameScoreTbl  = [...]int{3, 0, 6}
	shapeScoreTbl = [...]int{1, 2, 3}
)

func score(c1, c2 uint8) int {
	return gameScoreTbl[(3+c1-c2)%3] + shapeScoreTbl[c2]
}

func conv1(c1, c2 uint8) (uint8, uint8) {
	return c1 - 'A', c2 - 'X'
}

func conv2(c1, c2 uint8) (uint8, uint8) {
	c1 -= 'A'
	c2 = (2 + (c2 - 'X') + c1) % 3
	return c1, c2
}

func main() {
	var convFn = conv2
	var sum int
	aoc.ReadLines(os.Stdin, func(line []byte) {
		sum += score(convFn(line[0], line[2]))
	})
	fmt.Println(sum)
}
