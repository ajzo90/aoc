package main

import (
	"aoc2022/pkg/aoc"
	"io"
	"strconv"
	"strings"
)

const day = 4

func main() {

	const example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8`

	aoc.Assert(example, fn1, "2")

	aoc.WithData(day, fn1, "515")

	aoc.Assert(example, fn2, "4")

	aoc.WithData(day, fn2, "883")
}

var fn1 = func(f io.Reader) string {
	var res int
	for _, line := range aoc.ReadLinesToList(f) {
		var p1, p2, _ = strings.Cut(string(line), ",")
		a, b := toSet(p1), toSet(p2)

		inter := intersection(a, b)
		if len(a) == len(inter) || len(b) == len(inter) {
			res++
		}
	}
	return strconv.Itoa(res)
}

var fn2 = func(f io.Reader) string {
	var res int
	for _, line := range aoc.ReadLinesToList(f) {
		var p1, p2, _ = strings.Cut(string(line), ",")
		a, b := toSet(p1), toSet(p2)

		if len(intersection(a, b)) > 0 {
			res++
		}
	}
	return strconv.Itoa(res)
}

func toSet(s string) map[int]struct{} {
	a, b, _ := strings.Cut(s, "-")
	n1, _ := strconv.Atoi(a)
	n2, _ := strconv.Atoi(b)
	m := map[int]struct{}{}
	for i := n1; i <= n2; i++ {
		m[i] = struct{}{}
	}
	return m
}

func intersection(a, b map[int]struct{}) map[int]struct{} {
	m := map[int]struct{}{}
	for k := range a {
		if _, ok := b[k]; ok {
			m[k] = struct{}{}
		}
	}
	return m
}
