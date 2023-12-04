package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
	"log"
	"strconv"
)

func getValue(rows []string, i, j int) int {
	if i >= len(rows) || i < 0 {
		return -1
	}
	var row = rows[i]
	if j >= len(row) || j < 0 {
		return -1
	}
	return Int(row[j:][:1])
}

func run(s string, rows []string) string {
	var sum int
	for i, row := range rows {
		for j, cell := range row {
			var visibleEdge = i == 0 || i+1 == len(rows) || j == 0 || j+1 == len(row)
			var maxLeft, mixRight, maxTop, maxBottom = -1, -1, -1, -1

			ref := int(cell - '0')

			for jj := range row {
				v := getValue(rows, i, jj)
				if jj < j && v > maxLeft {
					maxLeft = v
				} else if jj > j && v > mixRight {
					mixRight = v
				}
			}

			for ii := range rows {
				v := getValue(rows, ii, j)
				if ii < i && v > maxTop {
					maxTop = v
				} else if ii > i && v > maxBottom {
					maxBottom = v
				}
			}

			visible := visibleEdge || ref > maxLeft || ref > maxTop || ref > mixRight || ref > maxBottom
			if visible {
				sum++
			}
		}
	}
	return strconv.Itoa(sum)
}

func run2(s string, rows []string) string {

	//var facit = map[int]map[int]int{
	//	1: {1: 1},
	//	2: {1: 4, 3: 2},
	//}

	var max int
	for i, row := range rows {
		for j := range row {

			if i == 0 || j == 0 || i+1 == len(rows) || j+1 == len(row) {
				continue
			}

			var l, r, t, b int
			var ref = getValue(rows, i, j)

			var leftV = -1
			for jj := j - 1; jj >= 0 && leftV < ref; jj-- {
				if v := getValue(rows, i, jj); true || v >= leftV {
					l++
					leftV = v
				}
			}

			var rightV = -1
			for jj := j + 1; jj < len(row) && rightV < ref; jj++ {
				if v := getValue(rows, i, jj); true || v >= rightV {
					r++
					rightV = v
				}
			}

			var topV = -1
			for ii := i - 1; ii >= 0 && topV < ref; ii-- {
				if v := getValue(rows, ii, j); true || v >= topV {
					t++
					topV = v
				}
			}

			var bottomV = -1
			for ii := i + 1; ii < len(rows) && bottomV < ref; ii++ {
				if v := getValue(rows, ii, j); true || v >= bottomV {
					b++
					bottomV = v
				}
			}

			var score = l * r * t * b

			//if v, ok := facit[i][j]; ok && v != score {
			//	panic("invalid")
			//}

			if score >= max {
				log.Println(i, j, max, score, l, r, t, b)
				max = score
			}
		}
	}
	return strconv.Itoa(max)
}

func main() {

	var aoc = New()

	var input, rows = aoc.Input()

	fmt.Println(input)

	Assert(example, "21", run)

	ans := run(input, rows)
	aoc.Submit(1, ans)

	Assert(example, "8", run2)
	log.Println("XX")

	aoc.Submit(2, run2(input, rows))

}

const example = `30373
25512
65332
33549
35390`
