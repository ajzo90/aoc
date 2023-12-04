package main

import (
	. "aoc2022/pkg/aoc"
	"fmt"
	"log"
	"strings"
)

const desc = `
DESCRIPTION_HERE
`

const example = `
    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func code(text string, rows []string) string {

	// 1. read stacks
	state, instructions, _ := strings.Cut(text, "\n\n")
	fmt.Println(state)

	var stacks = map[int][]string{}

	var stateRows = strings.Split(state, "\n")

	var numStacks = (len(stateRows[len(stateRows)-1]) + 1) / 4

	for i := len(stateRows) - 2; i >= 0; i-- {
		row := stateRows[i] + strings.Repeat(" ", numStacks*4-1-len(stateRows[i]))

		for j := 1; j <= numStacks; j++ {
			val := row[1+4*(j-1):][:1]
			if val != " " {
				stacks[j] = append(stacks[j], val)
			}
		}
	}

	for _, line := range strings.Split(instructions, "\n") {
		nums := ExtractNumbers(line)
		num, from, to := nums[0], nums[1], nums[2]

		for j := 0; j < num; j++ {
			stack := stacks[from]
			stacks[to] = append(stacks[to], stack[len(stack)-1])
			stacks[from] = stack[:len(stack)-1]
		}
	}

	var msg string
	for i := 1; i <= numStacks; i++ {
		msg += stacks[i][len(stacks[i])-1]
	}

	return msg
}
func code2(text string, rows []string) string {

	// 1. read stacks
	state, instructions, _ := strings.Cut(text, "\n\n")
	fmt.Println(state)

	var stacks = map[int][]string{}

	var stateRows = strings.Split(state, "\n")

	var numStacks = (len(stateRows[len(stateRows)-1]) + 1) / 4

	for i := len(stateRows) - 2; i >= 0; i-- {
		row := stateRows[i] + strings.Repeat(" ", numStacks*4-1-len(stateRows[i]))

		for j := 1; j <= numStacks; j++ {
			val := row[1+4*(j-1):][:1]
			if val != " " {
				stacks[j] = append(stacks[j], val)
			}
		}
	}

	for _, line := range strings.Split(instructions, "\n") {
		nums := ExtractNumbers(line)
		num, from, to := nums[0], nums[1], nums[2]

		stack := stacks[from]
		stacks[to] = append(stacks[to], stack[len(stack)-num:]...)
		stacks[from] = stack[:len(stack)-num]
	}

	log.Println(stacks)
	var msg string
	for i := 1; i <= numStacks; i++ {
		msg += stacks[i][len(stacks[i])-1]
	}

	return msg
}

func main() {
	New(2022, 5, example).
		Part1("CMZ", code).
		Part2("MCD", code2)
}
