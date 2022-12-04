package parser

import (
	"strconv"
	"strings"
)

type parser struct {
	s string
}

func New(s string) *parser {
	return &parser{s: s}
}

func (p *parser) Skip(b string) *parser {
	if strings.HasPrefix(p.s, b) {
		p.s = strings.TrimPrefix(p.s, b)
	} else {
		panic("invalid")
	}
	return p
}

func MustParse(s, pattern string, sep string, v ...*int) {
	if err := New(s).Parse(pattern, sep, v...); err != nil {
		panic(err)
	}
}

func (p *parser) Parse(pattern string, sep string, v ...*int) error {
	for {
		before, after, ok := strings.Cut(pattern, sep)
		p.Skip(before)
		if !ok {
			return nil
		}
		p.Int(v[0])
		v = v[1:]
		pattern = after
	}
}

func (p *parser) Int(v *int) *parser {
	s := p.s
	l := 0
	for i := 0; i < len(s) && (s[i] >= '0' && s[i] <= '9'); i++ {
		l++
	}

	num, _ := strconv.ParseInt(s[:l], 10, 64)
	p.s = s[l:]
	*v = int(num)
	return p
}
