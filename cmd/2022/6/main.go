package main

import (
	. "aoc2022/pkg/aoc"
	"strconv"
)

func part1() func(text string, rows []string) string {
	return func(text string, rows []string) string {
		const n = 14
		for i := range text[:len(text)-n+1] {
			var S = NewSet[string]()
			for j := i; j < i+n; j++ {
				S.Add(text[j : j+1])
			}
			if len(S) == n {
				c := strconv.Itoa(i + n)
				return c
			}
		}
		return ""
	}
}

func part2(text string, rows []string) string {
	return ""
}

func main() {

	Assert("mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19", part1())
	Assert("EXAMPLE_PART_2", "EXPECTED_PART_2", part2)

	New(2022, 6).Part1(part1()).Part2(part2)
}
