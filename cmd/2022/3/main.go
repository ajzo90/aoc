package main

import (
	"aoc2022/pkg/io"
	"fmt"
	"os"
)

func main() {
	var sum int
	var lines [][]byte

	var parts = [2]func([]byte){
		func(l []byte) {
			sum += compute(l[:len(l)/2], l[len(l)/2:])
		},
		func(line []byte) {
			lines = append(lines, append([]byte{}, line...))
			if len(lines) == 3 {
				sum += compute(lines...)
				lines = lines[:0]
			}
		},
	}

	io.ReadLines(os.Stdin, parts[0])

	fmt.Println(sum)
}

func compute(lines ...[]byte) int {
	var m [256]byte
	var n = len(lines)

	for i, line := range lines {
		for _, c := range line {
			m[c] = m[c] | (1 << i)
		}
	}

	for k, v := range m {
		if intersect := v == (1<<n - 1); intersect {
			return prio(uint8(k))
		}
	}
	return 0
}

func prio(c uint8) int {
	if c >= 'a' && c <= 'z' {
		return 1 + int(c-'a')
	} else if c >= 'A' && c <= 'Z' {
		return 27 + int(c-'A')
	}
	return 0
}
