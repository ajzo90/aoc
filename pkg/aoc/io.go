package aoc

import (
	"bufio"
	"io"
	"strconv"
)

func ReadLines(r io.Reader, f func(line []byte)) {

	var rd = bufio.NewReader(r)

	var linebuf []byte

	for {
		line, isPrefix, err := rd.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			panic("invalid")
		} else if isPrefix {
			linebuf = append(linebuf[:0], line...)
			// need to read next chunk
			continue
		}

		if len(linebuf) > 0 {
			linebuf = append(linebuf, line...)
			line = linebuf
			linebuf = linebuf[:0]
		}

		f(line)
	}
}

func ReadLinesToList(r io.Reader) [][]byte {
	var lines [][]byte
	ReadLines(r, func(line []byte) {
		lines = append(lines, append([]byte{}, line...))
	})
	return lines
}

func ReadLinesToIntList(r io.Reader) []int {
	var out = make([]int, 0, 128)
	ReadLines(r, func(line []byte) {
		v, _ := strconv.Atoi(string(line))
		out = append(out, v)
	})
	return out
}
