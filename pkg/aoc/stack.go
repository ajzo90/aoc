package aoc

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(v T) {
	s.values = append(s.values, v)
}

func (s *Stack[T]) Pop() T {
	v := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return v
}
