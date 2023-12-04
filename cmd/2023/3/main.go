package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
	"log"
)

const desc = `
PASTE_THE_EXERCISE_HERE
`

const example = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

type p struct {
	a, b int
}
type acc struct {
	res   int
	arr   []string
	gears map[p][]int
}

func (a *acc) Reduce(s string, i int, arr []string) {

	a.arr = arr
}

func (a *acc) isSym(row, col int, num int) bool {
	if row < 0 || row >= len(a.arr) {
		return false
	}
	colV := a.arr[row]
	if col < 0 || col >= len(colV) {
		return false
	}

	c := a.arr[row][col]

	isGear := c == '*'
	if isGear {
		a.gears[p{row, col}] = append(a.gears[p{row, col}], num)
	}

	return c != '.'
}

func (a *acc) Result() int {
	a.res = 0
	a.gears = make(map[p][]int)

	for i, row := range a.arr {
		//fmt.Println(i, a.res)
		var j int
		fmt.Print(row)
		for len(row) > 0 {
			var n = []int{}
			for _, c := range row {
				if c >= '0' && c <= '9' {
					n = append(n, int(c-'0'))
				} else {
					break
				}
			}
			if len(n) > 0 {

				var hasNeightbour bool
				var s = 0
				for _, v := range n {
					s = s*10 + v
				}

				for k, _ := range n {
					jj := j + k
					hasNeightbour = hasNeightbour || a.isSym(i-1, jj, s) || a.isSym(i+1, jj, s)
				}

				hasNeightbour = hasNeightbour || a.isSym(i, j-1, s) || a.isSym(i, j+len(n), s) || a.isSym(i-1, j-1, s) || a.isSym(i+1, j-1, s) || a.isSym(i-1, j+len(n), s) || a.isSym(i+1, j+len(n), s)
				fmt.Print(" ", s, hasNeightbour)

				if hasNeightbour {
					a.res += s
				}

				row = row[len(n):]
				j += len(n)

			} else {
				row = row[1:]
				j++
			}

		}
		fmt.Println()
	}
	for _, row := range a.arr[len(a.arr)-10:] {
		fmt.Println(row)
	}

	if a.res != 4361 {
		if a.res != 532331 {
			log.Fatal(a.res, 532331)
		}
	}

	var acc int
	for _, v := range a.gears {
		if len(v) == 2 {
			acc += v[0] * v[1]
		}
	}

	log.Println("ACC", acc)

	return acc
}

func main() {
	New(2023, 3, example).
		//ReducePart(1, 4361, &acc{}).
		ReducePart(2, 467835, &acc{})
}
