package insert

import (
	"math/rand"
	"slices"
	"testing"
)

func genRandSlice[T any](n int, rnd func() T) []T {
	s := make([]T, n)
	for i := 0; i < n; i++ {
		s[i] = rnd()
	}
	return s
}

func rndNumber[N int | int64 | float64](rnd func() N) func() N {
	return func() N {
		d := rand.Intn(2)
		if d == 0 {
			return rnd() * -1
		}
		return rnd()
	}
}

func TestInsertInts(t *testing.T) {
	for i := 0; i < 500; i++ {
		n := rand.Intn(10000)
		s := genRandSlice(n, rndNumber(rand.Int))
		Insert(s)
		if !slices.IsSorted(s) {
			t.Error("expected slice to be sorted", s)
		}
	}
}

func TestInsertFloats(t *testing.T) {
	for i := 0; i < 500; i++ {
		n := rand.Intn(10000)
		s := genRandSlice(n, rndNumber(rand.NormFloat64))
		Insert(s)
		if !slices.IsSorted(s) {
			t.Error("expected slice to be sorted", s)
		}
	}
}

func Benchmark(b *testing.B) {
	const n = 10000
	sI := genRandSlice(n, rndNumber(rand.Int))
	sF := genRandSlice(n, rndNumber(rand.NormFloat64))
	b.ResetTimer()

	b.Run("insert ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Insert(sI)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sI[i], sI[j] = sI[j], sI[i]
			})
			b.StartTimer()
		}
	})
	b.Run("slices.sort ints", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slices.Sort(sI)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sI[i], sI[j] = sI[j], sI[i]
			})
			b.StartTimer()
		}
	})
	b.Run("insert floats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			Insert(sF)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sF[i], sF[j] = sF[j], sF[i]
			})
			b.StartTimer()
		}
	})
	b.Run("slices.sort floats", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			slices.Sort(sF)
			b.StopTimer()
			rand.Shuffle(n, func(i, j int) {
				sF[i], sF[j] = sF[j], sF[i]
			})
			b.StartTimer()
		}
	})
}
