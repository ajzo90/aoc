package aoc

type numerical interface {
	int
}

func SlidingSum[T numerical, Coll ~[]T](l Coll, n int) Coll {
	var out = make([]T, len(l)-n+1)
	for i := range out {
		out[i] = Sum[T](l[i : i+n])
	}
	return out
}

func Sum[T numerical, Coll ~[]T](coll Coll) T {
	var sum T
	for _, v := range coll {
		sum += v
	}
	return sum
}
