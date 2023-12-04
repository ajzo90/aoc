package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func run(s string, rows []string) string {
	log.Println(s)
	var sum int
	for _, row := range rows {
		_, x, _ := strings.Cut(row, ": ")
		win, my, _ := strings.Cut(x, " | ")
		winners := SplitToSet(win, " ")
		mine := SplitToSet(my, " ")

		overlap := Intersection(winners, mine).Len()
		if overlap == 0 {
			continue
		}
		res := int(math.Pow(2, float64(overlap-1)))

		log.Println(winners, mine, res)

		sum += res

	}

	return strconv.Itoa(sum)
}

func run2(s string, rows []string) string {
	var sum int
	var cards = map[int]int{}

	for i, row := range rows {
		_, x, _ := strings.Cut(row, ": ")
		win, my, _ := strings.Cut(x, " | ")
		winners := SplitToSet(win, " ")
		mine := SplitToSet(my, " ")

		overlap := Intersection(winners, mine).Len()
		var score = 1
		if overlap == 0 {
			score = 0
		}

		sum += cards[i] + 1
		if score == 1 {
			for j := 0; j < overlap; j++ {
				cards[i+j+1] += cards[i] + 1
			}
		}
	}

	return strconv.Itoa(sum)
}

func main() {

	var aoc = New()

	var input, rows = aoc.Input()

	ToTable(rows)

	fmt.Println(input)

	Assert(example, "13", run)

	ans := run(input, rows)
	aoc.Submit(1, ans)

	Assert(example, "30", run2)

	aoc.Submit(2, run2(input, rows))

}

const example = `Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
