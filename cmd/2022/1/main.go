package main

import (
	"aoc2022/pkg/aoc"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var sum int

	var maxSums []int

	const n = 3

	aoc.ReadLines(os.Stdin, func(line []byte) {
		if len(line) == 0 {
			maxSums = append(maxSums, sum)
			if len(maxSums) == n+1 {
				sort.Slice(maxSums, func(i, j int) bool {
					return maxSums[i] > maxSums[j]
				})
				maxSums = maxSums[:n]
			}
			sum = 0
		}
		v, _ := strconv.Atoi(string(line))
		sum += v
	})

	var totSum int
	for _, v := range maxSums {
		totSum += v
	}
	log.Println(totSum, maxSums)
}
