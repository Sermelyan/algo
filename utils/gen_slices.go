package utils

import "math/rand"

func GenRandomSlice[T any](n int, rnd func() T) []T {
	s := make([]T, n)
	for i := 0; i < n; i++ {
		s[i] = rnd()
	}
	return s
}

func RandomNumber[N int | int64 | float64](rnd func() N) func() N {
	return func() N {
		d := rand.Intn(2)
		if d == 0 {
			return rnd() * -1
		}
		return rnd()
	}
}
