package main

import (
	"aoc2022/pkg/aoc"
	"bytes"
	"io"
	"strconv"
	"strings"
)

func main() {

	var run = func(part2 bool) func(r io.Reader) int {
		return func(r io.Reader) int {
			var limit = map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}
			var sum int
			aoc.ReadLines(r, func(line []byte) {
				game, rest, _ := bytes.Cut(line, []byte(": "))
				gameId, _ := strconv.Atoi(strings.Split(string(game), " ")[1])
				var maxSeen = map[string]int{}
				sets := bytes.Split(rest, []byte("; "))
				var possible = true
				for _, v := range sets {
					bags := bytes.Split(v, []byte(", "))
					for _, bag := range bags {
						numS, color, _ := bytes.Cut(bag, []byte(" "))
						num, _ := strconv.Atoi(string(numS))
						if limit[string(color)] < num {
							possible = false
						}
						if maxSeen[string(color)] < num {
							maxSeen[string(color)] = num
						}
					}
				}
				if part2 {
					var power = 1
					for _, v := range maxSeen {
						power *= v
					}
					sum += power
				} else if possible {
					sum += gameId
				}

			})
			return sum
		}
	}
	aoc.
		New(2023, 2, example).
		Part(1, "8", run(false)).
		Part(2, "2286", run(true))
}

const example = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
