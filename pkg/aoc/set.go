package aoc

import (
	"golang.org/x/exp/maps"
	"strings"
)

type Set[K comparable] map[K]struct{}

func (s Set[K]) Len() int {
	return len(s)
}

func SplitToSet(st string, del string) Set[string] {

	s := strings.Split(strings.TrimSpace(st), del)

	var out []string
	for i := range s {
		x := strings.TrimSpace(s[i])
		if x != "" {
			out = append(out, x)
		}

	}

	return NewSet(out...)
}

func NewSet[K comparable](values ...K) Set[K] {
	var s = Set[K]{}
	for _, v := range values {
		s[v] = struct{}{}
	}
	return s
}

func (s Set[K]) Add(k K) Set[K] {
	s[k] = struct{}{}
	return s
}

func Intersection[K comparable](a, b Set[K]) Set[K] {
	var res = Set[K]{}
	for k := range a {
		if _, ok := b[k]; ok {
			res[k] = struct{}{}
		}
	}
	return res
}

func (s Set[K]) Keys() []K {
	return maps.Keys(s)
}

func ConvertSet[K comparable, T comparable](from Set[K], conv func(f K) T) Set[T] {
	var t = Set[T]{}
	for k := range from {
		t.Add(conv(k))
	}
	return t
}

func Difference[K comparable](a, b Set[K]) Set[K] {
	var res = Set[K]{}
	for k := range a {
		if _, ok := b[k]; !ok {
			res[k] = struct{}{}
		}
	}
	return res
}

func Union[K comparable](a, b Set[K]) Set[K] {
	var res = Set[K]{}
	for k := range a {
		res[k] = struct{}{}
	}
	for k := range b {
		res[k] = struct{}{}
	}
	return res
}

func SymmetricDifference[K comparable](a, b Set[K]) Set[K] {
	var res = Set[K]{}
	for k := range a {
		if _, ok := b[k]; !ok {
			res[k] = struct{}{}
		}
	}
	for k := range b {
		if _, ok := a[k]; !ok {
			res[k] = struct{}{}
		}
	}
	return res
}
