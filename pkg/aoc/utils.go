package aoc

import (
	"fmt"
	"strconv"
)

func Must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func Print(idx int, line string, a ...any) {
	fmt.Printf("==> line %d ==> %s\n", idx, line)
	fmt.Println(a...)
	fmt.Println()
}

func Int(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}

func ExtractNumbers(s string) []int {
	var res []int
	var stack []int
	for _, c := range s {
		if c >= '0' && c <= '9' {
			stack = append(stack, int(c-'0'))
		} else if len(stack) > 0 {
			var v int
			for _, d := range stack {
				v = v*10 + d
			}
			res = append(res, v)
			stack = stack[:0]
		}
	}
	if len(stack) == 0 {
		return res
	}
	var v int
	for _, d := range stack {
		v = v*10 + d
	}
	res = append(res, v)
	return res
}

type Table struct {
	rows []string
}

func ToTable(rows []string) *Table {
	return &Table{rows: rows}
}

func (t *Table) Get(i, j int, defaultV int) int {
	c := t.GetChar(i, j, "")
	if c == "" {
		return defaultV
	}
	return Int(c)
}

func (t *Table) GetChar(i, j int, defaultV string) string {
	if i < 0 || i >= len(t.rows) {
		return defaultV
	} else if row := t.rows[i]; j < 0 || j >= len(row) {
		return defaultV
	} else {
		return string(row[j])
	}
}
