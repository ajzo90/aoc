package parser

import (
	"log"
	"testing"
)

func Test_parser_NextInt(t *testing.T) {
	var a, b, c, d int
	MustParse("15435-4,5-43", "x-x,x-x", "x", &a, &b, &c, &d)
	log.Println(a, b, c, d)
}
