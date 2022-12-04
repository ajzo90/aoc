package main

import (
	"aoc2022/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

func main() {

	aoc.
		New(2021, 2, example).
		Part(1, "150", func(r io.Reader) int {
			horizontal, depth := 0, 0
			aoc.ReadLines(r, func(line []byte) {
				command, arg, _ := strings.Cut(string(line), " ")
				v, _ := strconv.Atoi(arg)

				switch command {
				case "forward":
					horizontal += v
				case "down":
					depth += v
				case "up":
					depth -= v
				default:
					panic(command)
				}
			})
			return horizontal * depth
		}).
		Part(2, "900", func(r io.Reader) int {
			horizontal, depth, aim := 0, 0, 0
			aoc.ReadLines(r, func(line []byte) {
				command, arg, _ := strings.Cut(string(line), " ")
				v, _ := strconv.Atoi(arg)

				switch command {
				case "forward":
					horizontal += v
					depth += aim * v
				case "down":
					aim += v
				case "up":
					aim -= v
				default:
					panic(command)
				}
			})
			return horizontal * depth
		})
}

const example = `forward 5
down 5
forward 8
up 3
down 8
forward 2`
