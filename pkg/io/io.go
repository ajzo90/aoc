package io

import (
	"bufio"
	"io"
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
