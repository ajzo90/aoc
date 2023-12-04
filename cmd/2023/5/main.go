package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
)

func run(s string, rows []string) string {
	return ""
}

func run2(s string, rows []string) string {
	return ""
}

func main() {

	var aoc = New()

	var input, rows = aoc.Input()

	fmt.Println(input)

	Assert(example, "13", run)

	aoc.Submit(1, run(input, rows))

	Assert(example, "30", run2)

	aoc.Submit(2, run2(input, rows))

}

const example = `PASTE_EXAMPLE`
